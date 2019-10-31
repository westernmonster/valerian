package service

import (
	"context"
	"time"
	account "valerian/app/service/account/api"
	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
)

func (p *Service) follow(c context.Context, node sqalx.Node, arg *api.ArgTopicFollow) (status int32, err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": arg.Aid, "topic_id": arg.TopicID}); err != nil {
		return
	} else if member != nil {
		return model.FollowStatusFollowed, nil
	}

	var req *model.TopicFollowRequest
	if req, err = p.d.GetTopicFollowRequestByCond(c, node, map[string]interface{}{
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
	if t, err = p.d.GetTopicByID(c, node, arg.TopicID); err != nil {
		return
	} else if t == nil {
		return 0, ecode.TopicNotExist
	}

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, arg.Aid); err != nil {
		return
	} else if account == nil {
		return 0, ecode.UserNotExist
	}

	item := &model.TopicFollowRequest{
		ID:        gid.NewID(),
		Status:    model.FollowRequestStatusCommited,
		TopicID:   arg.TopicID,
		Reason:    arg.Reason,
		AccountID: arg.Aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	switch t.JoinPermission {
	case model.JoinPermissionMember:
		status = model.FollowStatusFollowed
		if err = p.addMember(c, node, arg.TopicID, arg.Aid, model.MemberRoleUser); err != nil {
			return
		}
		break
	case model.JoinPermissionMemberApprove:
		status = model.FollowStatusApproving
		if err = p.d.AddTopicFollowRequest(c, node, item); err != nil {
			return
		}
		break
	case model.JoinPermissionCertApprove:
		if !account.IDCert || !account.WorkCert {
			return model.FollowStatusUnfollowed, ecode.NeedWorkCert
		}

		status = model.FollowStatusApproving
		if err = p.d.AddTopicFollowRequest(c, node, item); err != nil {
			return
		}
		break
	case model.JoinPermissionManualAdd:
		status = model.FollowStatusUnfollowed
		err = ecode.OnlyAllowAdminAdded
		return
	}

	return

}

func (p *Service) auditFollow(c context.Context, node sqalx.Node, arg *api.ArgAuditFollow) (req *model.TopicFollowRequest, err error) {
	if req, err = p.d.GetTopicFollowRequestByID(c, node, arg.ID); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicFollowRequestNotExist
		return
	}

	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": req.AccountID, "topic_id": req.TopicID}); err != nil {
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

		if err = p.addMember(c, node, req.TopicID, req.AccountID, model.MemberRoleUser); err != nil {
			return
		}
	} else {
		req.Status = model.FollowRequestStatusRejected
		req.UpdatedAt = time.Now().Unix()
	}

	if err = p.d.UpdateTopicFollowRequest(c, node, req); err != nil {
		return
	}

	return

}
