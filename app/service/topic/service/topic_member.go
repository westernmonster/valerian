package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Leave 退出话题
func (p *Service) Leave(c context.Context, aid, topicID int64) (err error) {
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

	if err = p.leave(c, tx, aid, topicID); err != nil {
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
func (p *Service) GetTopicMembersPaged(c context.Context, arg *api.ArgTopicMembers) (resp *api.TopicMembersPagedResp, err error) {
	resp = new(api.TopicMembersPagedResp)
	resp.PageSize = arg.PageSize
	resp.Data = make([]*api.TopicMemberInfo, 0)

	var (
		addCache = true
		count    int32
		items    []*model.TopicMember
	)

	if count, items, err = p.d.TopicMembersCache(c, arg.TopicID, arg.Page, arg.PageSize); err != nil {
		addCache = false
	}

	if items == nil {
		if count, items, err = p.d.GetTopicMembersPaged(c, p.d.DB(), arg.TopicID, arg.Page, arg.PageSize); err != nil {
			return
		}
	}

	if items != nil && addCache {
		p.addCache(func() {
			p.d.SetTopicMembersCache(context.TODO(), arg.TopicID, count, arg.Page, arg.PageSize, items)
		})
	}

	for _, v := range items {
		account, e := p.d.GetAccountBaseInfo(c, v.AccountID)
		if e != nil {
			return
		}
		resp.Data = append(resp.Data, &api.TopicMemberInfo{
			AccountID: v.AccountID,
			Role:      v.Role,
			Avatar:    account.Avatar,
			UserName:  account.UserName,
		})

	}

	resp.Count = count
	return
}

func (p *Service) BulkSaveMembers(c context.Context, req *api.ArgBatchSavedTopicMember) (err error) {
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

	if err = p.bulkSaveMembers(c, tx, req); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelTopicCache(context.TODO(), req.TopicID)
		p.d.DelTopicMembersCache(context.TODO(), req.TopicID)
	})

	return
}

// ChangeOwner 更改主理人
func (p *Service) ChangeOwner(c context.Context, arg *api.ArgChangeOwner) (err error) {
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

	if err = p.changeOwner(c, tx, arg); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelTopicMembersCache(context.TODO(), arg.TopicID)
	})

	return
}

func (p *Service) IsTopicMember(c context.Context, aid int64, topicID int64) (ret bool, err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member != nil {
		ret = true
	}
	return
}
