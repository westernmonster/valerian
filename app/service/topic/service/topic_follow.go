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

func (p *Service) Follow(c context.Context, arg *api.ArgTopicFollow) (status int32, err error) {
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

	if status, err = p.follow(c, tx, arg); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return

}

func (p *Service) GetFollowedTopicsIDs(c context.Context, aid int64) (ids []int64, err error) {
	if ids, err = p.d.GetFollowedTopicsIDs(c, p.d.DB(), aid); err != nil {
		return
	}

	return
}

func (p *Service) AuditFollow(c context.Context, arg *api.ArgAuditFollow) (err error) {
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

	var req *model.TopicFollowRequest
	if req, err = p.auditFollow(c, tx, arg); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		switch req.Status {
		case model.FollowRequestStatusApproved:
			p.onTopicFollowApproved(c, req.ID, req.TopicID, arg.Aid, time.Now().Unix())
			break
		case model.FollowRequestStatusRejected:
			p.onTopicFollowRejected(c, req.ID, req.TopicID, arg.Aid, time.Now().Unix())
			break
		}
	})

	return

}
