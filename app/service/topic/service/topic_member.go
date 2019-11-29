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

func (p *Service) GetManageTopicIDsPaged(c context.Context, arg *api.UserTopicsReq) (ids []int64, err error) {
	return p.d.GetManageTopicIDsPaged(c, p.d.DB(), arg.AccountID, arg.Limit, arg.Offset)
}

func (p *Service) GetFollowedTopicIDsPaged(c context.Context, arg *api.UserTopicsReq) (ids []int64, err error) {
	return p.d.GetFollowedTopicIDsPaged(c, p.d.DB(), arg.AccountID, arg.Limit, arg.Offset)
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
		account, e := p.getAccount(c, p.d.DB(), v.AccountID)
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
	// 检测是否系统管理员或者话题管理员
	if err = p.checkTopicManagePermission(c, req.Aid, req.TopicID); err != nil {
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

	var change *model.MemberChange
	if change, err = p.bulkSaveMembers(c, tx, req); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		for _, v := range change.NewMembers {
			p.onTopicFollowed(context.TODO(), req.TopicID, v, time.Now().Unix())
		}
		for _, v := range change.DelMembers {
			p.onTopicLeaved(context.TODO(), req.TopicID, v, time.Now().Unix())
		}
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

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicMembersCache(context.TODO(), arg.TopicID)
	})

	return
}
