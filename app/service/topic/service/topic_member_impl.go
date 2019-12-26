package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

// leave 离开话题
func (p *Service) leave(c context.Context, node sqalx.Node, aid, topicID int64) (err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member == nil {
		err = ecode.NotTopicMember
		return
	}

	if member.Role == model.MemberRoleOwner {
		err = ecode.OwnerNeedTransfer
		return
	}

	if err = p.d.DelTopicMember(c, node, member.ID); err != nil {
		return
	}

	// 更新成员统计信息
	if err = p.d.IncrTopicStat(c, node, &model.TopicStat{TopicID: topicID, MemberCount: -1}); err != nil {
		return
	}

	return
}

// getTopicMembers 获取话题成员
func (p *Service) getTopicMembers(c context.Context, node sqalx.Node, topicID int64, limit int32) (total int32, resp []*api.TopicMemberInfo, err error) {
	resp = make([]*api.TopicMemberInfo, 0)

	var (
		addCache = true
		items    []*model.TopicMember
	)

	if total, items, err = p.d.TopicMembersCache(c, topicID, int32(1), int32(10)); err != nil {
		addCache = false
	}

	if items == nil {
		if total, items, err = p.d.GetTopicMembersPaged(c, p.d.DB(), topicID, int32(1), int32(10)); err != nil {
			return
		}
	}

	if items != nil && addCache {
		p.addCache(func() {
			p.d.SetTopicMembersCache(context.TODO(), topicID, total, int32(1), int32(10), items)
		})
	}

	for _, v := range items {
		acc, e := p.getAccount(c, node, v.AccountID)
		if e != nil {
			err = e
			return
		}
		resp = append(resp, &api.TopicMemberInfo{
			AccountID: v.AccountID,
			Role:      v.Role,
			Avatar:    acc.Avatar,
			UserName:  acc.UserName,
		})
	}

	return
}

// createOwner 更改主理人
func (p *Service) createOwner(c context.Context, node sqalx.Node, aid, topicID int64) (err error) {
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

// addMember 添加成员
func (p *Service) addMember(c context.Context, node sqalx.Node, topicID, aid int64, role string) (err error) {
	if _, err = p.getAccount(c, node, aid); err != nil {
		return
	}

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

	if err = p.d.IncrTopicStat(c, node, &model.TopicStat{TopicID: topicID, MemberCount: 1}); err != nil {
		return
	}

	return
}

// bulkSaveMembers 批量保存话题成员
func (p *Service) bulkSaveMembers(c context.Context, node sqalx.Node, req *api.ArgBatchSavedTopicMember) (change *model.MemberChange, err error) {
	change = &model.MemberChange{
		NewMembers: make([]int64, 0),
		DelMembers: make([]int64, 0),
	}

	if err = p.checkTopic(c, node, req.TopicID); err != nil {
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
		if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": v.AccountID, "topic_id": req.TopicID}); err != nil {
			return
		}

		switch v.Opt {
		case "C":
			if member != nil {
				continue
			}

			if _, err = p.getAccount(c, node, v.AccountID); err != nil {
				return
			}

			item := &model.TopicMember{
				ID:        gid.NewID(),
				AccountID: v.AccountID,
				Role:      v.Role,
				TopicID:   req.TopicID,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
			if err = p.d.AddTopicMember(c, node, item); err != nil {
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

			if err = p.d.AddAccountTopicSetting(c, node, setting); err != nil {
				return
			}

			if err = p.d.IncrTopicStat(c, node, &model.TopicStat{TopicID: req.TopicID, MemberCount: 1}); err != nil {
				return
			}

			change.NewMembers = append(change.NewMembers, v.AccountID)

			break
		case "U":
			if member == nil {
				continue
			}
			member.Role = v.Role
			if err = p.d.UpdateTopicMember(c, node, member); err != nil {
				return
			}
			break
		case "D":
			if member == nil {
				continue
			}
			if err = p.d.DelTopicMember(c, node, member.ID); err != nil {
				return
			}

			if err = p.d.IncrTopicStat(c, node, &model.TopicStat{TopicID: req.TopicID, MemberCount: -1}); err != nil {
				return
			}
			change.DelMembers = append(change.DelMembers, member.ID)

			break
		}

		dic[v.AccountID] = true
	}

	return
}

// changeOwner 变更主理人
func (p *Service) changeOwner(c context.Context, node sqalx.Node, arg *api.ArgChangeOwner) (err error) {
	var curMember *model.TopicMember
	if curMember, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{
		"topic_id":   arg.TopicID,
		"account_id": arg.Aid,
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
	if err = p.d.UpdateTopicMember(c, node, curMember); err != nil {
		return
	}

	var toMember *model.TopicMember
	if toMember, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{
		"topic_id":   arg.TopicID,
		"account_id": arg.ToAccountID,
	}); err != nil {
		return
	} else if toMember == nil {
		err = ecode.NotTopicMember
		return
	}

	toMember.Role = model.MemberRoleOwner
	if err = p.d.UpdateTopicMember(c, node, toMember); err != nil {
		return
	}

	return
}
