package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

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
		account, e := p.getAccountByID(c, p.d.DB(), v.AccountID)
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
		account, e := p.getAccountByID(c, node, v.AccountID)
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

func (p *Service) bulkCreateMembers(c context.Context, node sqalx.Node, aid, topicID int64, req *model.ArgCreateTopic) (err error) {
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

	for _, v := range req.Members {
		if v.AccountID == aid {
			continue
		}

		if v.Role == model.MemberRoleOwner {
			return ecode.OnlyAllowOneOwner
		}

		item := &model.TopicMember{
			ID:        gid.NewID(),
			AccountID: v.AccountID,
			Role:      v.Role,
			TopicID:   topicID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		if e := p.d.AddTopicMember(c, tx, item); e != nil {
			return e
		}
	}

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

func (p *Service) BulkSaveMembers(c context.Context, req *model.ArgBatchSavedTopicMember) (err error) {
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

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, tx, req.TopicID); err != nil {
		return
	} else if t == nil {
		return ecode.TopicNotExist
	}

	for _, v := range req.Members {
		if v.Role == model.MemberRoleOwner {
			continue
		}

		var member *model.TopicMember
		if member, err = p.d.GetTopicMemberByCondition(c, tx, req.TopicID, v.AccountID); err != nil {
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
			if err = p.d.DeleteTopicMember(c, tx, member.ID); err != nil {
				return
			}
			break
		}
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

func (p *Service) addMember(c context.Context, node sqalx.Node, topicID, aid int64, role string) (err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCondition(c, node, topicID, aid); err != nil {
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

	p.addCache(func() {
		p.d.DelTopicCache(context.TODO(), topicID)
		p.d.DelTopicMembersCache(context.TODO(), topicID)
	})
	return
}

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
	if curMember, err = p.d.GetTopicMemberByCondition(c, tx, arg.TopicID, aid); err != nil {
		return
	} else if curMember == nil {
		err = ecode.NotBelongToTopic
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
	if toMember, err = p.d.GetTopicMemberByCondition(c, tx, arg.TopicID, arg.ToAccountID); err != nil {
		return
	} else if toMember == nil {
		err = ecode.NotBelongToTopic
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
		p.d.DelTopicCache(context.TODO(), arg.TopicID)
		p.d.DelTopicMembersCache(context.TODO(), arg.TopicID)
	})

	return
}
