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

// Follow 加入话题
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

	// 如果已经是成员，返回已关注
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{"account_id": arg.Aid, "topic_id": arg.TopicID}); err != nil {
		return
	} else if member != nil {
		return model.FollowStatusFollowed, nil
	}

	var req *model.TopicFollowRequest
	if req, err = p.d.GetTopicFollowRequestByCond(c, tx, map[string]interface{}{
		"topic_id":   arg.TopicID,
		"account_id": arg.Aid,
	}); err != nil {
		return
	}

	if req != nil {
		switch req.Status {
		case model.FollowRequestStatusCommited:
			return model.FollowStatusApproving, nil
		case model.FollowRequestStatusApproved:
			return model.FollowStatusFollowed, nil
		case model.FollowRequestStatusRejected:
			break
		}
	}

	var t *model.Topic
	if t, err = p.getTopic(c, tx, arg.TopicID); err != nil {
		return
	}

	var acc *model.Account
	if acc, err = p.getAccount(c, tx, arg.Aid); err != nil {
		return
	}

	item := &model.TopicFollowRequest{
		ID:            gid.NewID(),
		Status:        model.FollowRequestStatusCommited,
		TopicID:       arg.TopicID,
		Reason:        arg.Reason,
		AccountID:     arg.Aid,
		AllowViewCert: types.BitBool(arg.AllowViewCert),
		CreatedAt:     time.Now().Unix(),
		UpdatedAt:     time.Now().Unix(),
	}

	sentRequest := false
	addedMember := false
	switch t.JoinPermission {
	case model.JoinPermissionMember:
		status = model.FollowStatusFollowed
		if err = p.addMember(c, tx, arg.TopicID, arg.Aid, model.MemberRoleUser); err != nil {
			return
		}
		addedMember = true
		break
	case model.JoinPermissionMemberApprove:
		status = model.FollowStatusApproving
		if err = p.d.AddTopicFollowRequest(c, tx, item); err != nil {
			return
		}
		sentRequest = true
		break
	case model.JoinPermissionCertApprove:
		if !acc.IDCert || !acc.WorkCert {
			return model.FollowStatusUnfollowed, ecode.NeedWorkCert
		}

		status = model.FollowStatusApproving
		if err = p.d.AddTopicFollowRequest(c, tx, item); err != nil {
			return
		}
		sentRequest = true
		break
	case model.JoinPermissionManualAdd:
		status = model.FollowStatusUnfollowed
		err = ecode.OnlyAllowAdminAdded
		return
	}
	if err = p.d.SetTopicUpdatedAt(c, tx, arg.TopicID, time.Now().Unix()); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		if sentRequest {
			p.onTopicFollowRequested(context.Background(), item.ID, arg.TopicID, arg.Aid, time.Now().Unix())
		}

		if addedMember {
			p.onTopicFollowed(context.Background(), arg.TopicID, arg.Aid, time.Now().Unix())
		}
		p.d.DelTopicMembersCache(context.Background(), req.TopicID)
		p.d.DelTopicCache(context.Background(), req.TopicID)
	})

	return

}

// GetFollowedTopicsIDs 获取加入的话题ID列表
func (p *Service) GetFollowedTopicsIDs(c context.Context, aid int64) (ids []int64, err error) {
	if ids, err = p.d.GetFollowedTopicsIDs(c, p.d.DB(), aid); err != nil {
		return
	}

	return
}

// AuditFollow 审批加入话题请求
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
	if req, err = p.d.GetTopicFollowRequestByID(c, tx, arg.ID); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicFollowRequestNotExist
		return
	}

	// 检测是否系统管理员或者话题管理员
	if err = p.checkTopicManagePermission(c, tx, arg.Aid, req.TopicID); err != nil {
		return
	}

	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{"account_id": req.AccountID, "topic_id": req.TopicID}); err != nil {
		return
	} else if member != nil {
		return
	}

	switch req.Status {
	case model.FollowRequestStatusApproved:
	case model.FollowRequestStatusRejected:
		return
	}

	if arg.Approve {
		req.Status = model.FollowRequestStatusApproved
		req.UpdatedAt = time.Now().Unix()

		if err = p.addMember(c, tx, req.TopicID, req.AccountID, model.MemberRoleUser); err != nil {
			return
		}
	} else {
		req.Status = model.FollowRequestStatusRejected
		req.RejectReason = arg.Reason
		req.UpdatedAt = time.Now().Unix()
	}

	if err = p.d.UpdateTopicFollowRequest(c, tx, req); err != nil {
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
		switch req.Status {
		case model.FollowRequestStatusApproved:
			p.onTopicFollowed(context.Background(), req.TopicID, req.AccountID, time.Now().Unix())
			p.onTopicFollowApproved(context.Background(), req.ID, req.TopicID, arg.Aid, time.Now().Unix())
			break
		case model.FollowRequestStatusRejected:
			p.onTopicFollowRejected(context.Background(), req.ID, req.TopicID, arg.Aid, time.Now().Unix())
			break
		}

		p.d.DelTopicCache(context.Background(), req.TopicID)
		p.d.DelTopicMembersCache(context.Background(), req.TopicID)
	})

	return

}
