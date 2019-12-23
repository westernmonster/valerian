package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) Invite(c context.Context, arg *api.ArgTopicInvite) (err error) {
	if arg.Aid == arg.AccountID {
		err = ecode.InviteSelfNotAllowed
		return
	}

	if err = p.checkIsMember(c, arg.Aid, arg.TopicID); err != nil {
		return
	}

	var req *model.TopicInviteRequest
	if req, err = p.d.GetTopicInviteRequestByCond(c, p.d.DB(), map[string]interface{}{
		"topic_id":   arg.TopicID,
		"account_id": arg.AccountID,
		"status":     model.InviteStatusSent,
	}); err != nil {
		return
	} else if req != nil {
		return
	}

	if req == nil {
		item := &model.TopicInviteRequest{
			ID:            gid.NewID(),
			TopicID:       arg.TopicID,
			AccountID:     arg.AccountID,
			FromAccountID: arg.Aid,
			Status:        model.InviteStatusSent,
			CreatedAt:     time.Now().Unix(),
			UpdatedAt:     time.Now().Unix(),
		}

		if err = p.d.AddTopicInviteRequest(c, p.d.DB(), item); err != nil {
			return
		}

		p.addCache(func() {
			p.onTopicInviteSent(c, item.ID, item.TopicID, arg.Aid, item.CreatedAt)
		})
	}

	return
}

func (p *Service) ProcessInvite(c context.Context, arg *api.ArgProcessInvite) (err error) {
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

	var req *model.TopicInviteRequest
	if req, err = p.d.GetTopicInviteRequestByID(c, tx, arg.ID); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicInviteRequestNotExist
		return
	}

	if req.AccountID != arg.Aid {
		return
	}

	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{"account_id": req.AccountID, "topic_id": req.TopicID}); err != nil {
		return
	} else if member != nil {
		return
	}

	switch req.Status {
	case model.InviteStatusJoined:
	case model.InviteStatusRejected:
		return
	}

	if arg.Result {
		req.Status = model.InviteStatusJoined
		req.UpdatedAt = time.Now().Unix()
		if err = p.addMember(c, tx, req.TopicID, req.AccountID, model.MemberRoleUser); err != nil {
			return
		}
	} else {
		req.Status = model.InviteStatusRejected
		req.UpdatedAt = time.Now().Unix()
	}

	if err = p.d.UpdateTopicInviteRequest(c, tx, req); err != nil {
		return
	}

	if err = p.d.SetTopicUpdatedAt(c, tx, req.TopicID, time.Now().Unix()); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicMembersCache(context.TODO(), req.TopicID)
		if req.Status == model.InviteStatusJoined {
			p.onTopicFollowed(context.TODO(), req.TopicID, req.AccountID, time.Now().Unix())
		}
		p.d.DelTopicCache(context.TODO(), req.TopicID)
	})

	return

}
