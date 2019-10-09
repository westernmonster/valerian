package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

// isTopicMember 是否话题成员
func (p *Service) isTopicMember(c context.Context, node sqalx.Node, accountID, topicID int64) (isMember bool, err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": accountID, "topic_id": topicID}); err != nil {
		return
	} else if member != nil {
		isMember = true
	}

	return
}

// Leave 退出话题
func (p *Service) Leave(c context.Context, topicID int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member == nil {
		err = ecode.NotTopicMember
		return
	}

	if member.Role == model.MemberRoleOwner {
		err = ecode.OwnerNeedTransfer
		return
	}

	if err = p.d.DelTopicMember(c, tx, member.ID); err != nil {
		return
	}

	// 更新成员统计信息
	if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: topicID, MemberCount: -1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicMembersCache(context.TODO(), topicID)
		p.onTopicLeaved(c, aid, topicID, time.Now().Unix())
	})

	return
}

//  GetTopicMembersPaged 分页获取话题成员
func (p *Service) GetTopicMembersPaged(c context.Context, topicID int64, page, pageSize int) (resp *model.TopicMembersPagedResp, err error) {
	resp = new(model.TopicMembersPagedResp)
	resp.PageSize = pageSize
	resp.Data = make([]*model.TopicMemberResp, 0)

	var (
		addCache = true
		count    int
		items    []*model.TopicMember
	)

	if count, items, err = p.d.TopicMembersCache(c, topicID, page, pageSize); err != nil {
		addCache = false
	}

	if items == nil {
		if count, items, err = p.d.GetTopicMembersPaged(c, p.d.DB(), topicID, page, pageSize); err != nil {
			return
		}
	}

	if items != nil && addCache {
		p.addCache(func() {
			p.d.SetTopicMembersCache(context.TODO(), topicID, count, page, pageSize, items)
		})
	}

	for _, v := range items {
		account, e := p.d.GetAccountBaseInfo(c, v.AccountID)
		if e != nil {
			return
		}
		resp.Data = append(resp.Data, &model.TopicMemberResp{
			AccountID: v.AccountID,
			Role:      v.Role,
			Avatar:    account.Avatar,
			UserName:  account.UserName,
		})

	}

	resp.Count = count
	return
}

func (p *Service) getTopicMembers(c context.Context, node sqalx.Node, topicID int64, limit int) (total int, resp []*model.TopicMemberResp, err error) {
	resp = make([]*model.TopicMemberResp, 0)

	var (
		addCache = true
		items    []*model.TopicMember
	)

	if total, items, err = p.d.TopicMembersCache(c, topicID, 1, 10); err != nil {
		addCache = false
	}

	if items == nil {
		if total, items, err = p.d.GetTopicMembersPaged(c, p.d.DB(), topicID, 1, 10); err != nil {
			return
		}
	}

	if items != nil && addCache {
		p.addCache(func() {
			p.d.SetTopicMembersCache(context.TODO(), topicID, total, 1, 10, items)
		})
	}

	for _, v := range items {
		account, e := p.d.GetAccountBaseInfo(c, v.AccountID)
		if e != nil {
			return
		}
		resp = append(resp, &model.TopicMemberResp{
			AccountID: v.AccountID,
			Role:      v.Role,
			Avatar:    account.Avatar,
			UserName:  account.UserName,
		})
	}

	return
}

func (p *Service) createTopicMemberOwner(c context.Context, node sqalx.Node, aid, topicID int64) (err error) {
	var tx sqalx.Node
	if tx, err = node.Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	item := &model.TopicMember{
		ID:        gid.NewID(),
		AccountID: aid,
		Role:      model.MemberRoleOwner,
		TopicID:   topicID,
		Deleted:   false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	if err = p.d.AddTopicMember(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) checkTopicMemberAdmin(c context.Context, node sqalx.Node, topicID, aid int64) (err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member == nil {
		err = ecode.NotTopicMember
		return
	} else if member.Role == model.MemberRoleUser {
		err = ecode.NotTopicAdmin
		return
	}

	return nil
}

func (p *Service) BulkSaveMembers(c context.Context, req *model.ArgBatchSavedTopicMember) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	if err = p.checkTopic(c, tx, req.TopicID); err != nil {
		return
	}
	if err = p.checkTopicMemberAdmin(c, tx, req.TopicID, aid); err != nil {
		return
	}

	dic := make(map[int64]bool)
	for _, v := range req.Members {
		if v.Role == model.MemberRoleOwner {
			continue
		}

		if dic[v.AccountID] {
			err = ecode.TopicMemberDuplicate
			return
		}

		var member *model.TopicMember
		if member, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{"account_id": v.AccountID, "topic_id": req.TopicID}); err != nil {
			return
		}

		switch v.Opt {
		case "C":
			if member != nil {
				continue
			}
			item := &model.TopicMember{
				ID:        gid.NewID(),
				AccountID: v.AccountID,
				Role:      v.Role,
				TopicID:   req.TopicID,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
			if err = p.d.AddTopicMember(c, tx, item); err != nil {
				return
			}

			setting := &model.AccountTopicSetting{
				ID:               gid.NewID(),
				AccountID:        v.AccountID,
				TopicID:          req.TopicID,
				Important:        types.BitBool(false),
				Fav:              types.BitBool(false),
				MuteNotification: types.BitBool(false),
				CreatedAt:        time.Now().Unix(),
				UpdatedAt:        time.Now().Unix(),
			}

			if err = p.d.AddAccountTopicSetting(c, tx, setting); err != nil {
				return
			}

			if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: req.TopicID, MemberCount: 1}); err != nil {
				return
			}

			p.addCache(func() {
				p.onTopicFollowed(c, req.TopicID, v.AccountID, time.Now().Unix())
			})

			break
		case "U":
			if member == nil {
				continue
			}
			member.Role = v.Role
			if err = p.d.UpdateTopicMember(c, tx, member); err != nil {
				return
			}
			break
		case "D":
			if member == nil {
				continue
			}
			if err = p.d.DelTopicMember(c, tx, member.ID); err != nil {
				return
			}

			if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: req.TopicID, MemberCount: -1}); err != nil {
				return
			}

			p.addCache(func() {
				p.onTopicLeaved(c, req.TopicID, v.AccountID, time.Now().Unix())
			})
			break
		}

		dic[v.AccountID] = true
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicCache(context.TODO(), req.TopicID)
		p.d.DelTopicMembersCache(context.TODO(), req.TopicID)
	})

	return
}

// addMember 添加成员
func (p *Service) addMember(c context.Context, node sqalx.Node, topicID, aid int64, role string) (err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()
	{
		var member *model.TopicMember
		if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
			return
		} else if member != nil {
			return
		}

		item := &model.TopicMember{
			ID:        gid.NewID(),
			AccountID: aid,
			Role:      role,
			TopicID:   topicID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		if err = p.d.AddTopicMember(c, node, item); err != nil {
			return
		}

		setting := &model.AccountTopicSetting{
			ID:               gid.NewID(),
			AccountID:        aid,
			TopicID:          topicID,
			Important:        types.BitBool(false),
			Fav:              types.BitBool(false),
			MuteNotification: types.BitBool(false),
			CreatedAt:        time.Now().Unix(),
			UpdatedAt:        time.Now().Unix(),
		}

		if err = p.d.AddAccountTopicSetting(c, node, setting); err != nil {
			return
		}

		if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: topicID, MemberCount: -1}); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicMembersCache(context.TODO(), topicID)
		p.onTopicFollowed(c, topicID, aid, time.Now().Unix())
	})

	return
}

// ChangeOwner 更改主理人
func (p *Service) ChangeOwner(c context.Context, arg *model.ArgChangeOwner) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var curMember *model.TopicMember
	if curMember, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{
		"topic_id":   arg.TopicID,
		"account_id": aid,
	}); err != nil {
		return
	} else if curMember == nil {
		err = ecode.NotTopicMember
		return
	} else if curMember.Role != model.MemberRoleOwner {
		err = ecode.NotTopicOwner
		return
	}

	curMember.Role = model.MemberRoleAdmin
	if err = p.d.UpdateTopicMember(c, tx, curMember); err != nil {
		return
	}

	var toMember *model.TopicMember
	if toMember, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{
		"topic_id":   arg.TopicID,
		"account_id": arg.ToAccountID,
	}); err != nil {
		return
	} else if toMember == nil {
		err = ecode.NotTopicMember
		return
	}

	toMember.Role = model.MemberRoleOwner
	if err = p.d.UpdateTopicMember(c, tx, toMember); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicMembersCache(context.TODO(), arg.TopicID)
	})

	return
}
