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
)

func (p *Service) Follow(c context.Context, aid, topicID int64) (status int, err error) {
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
	var isMember bool
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCondition(c, p.d.DB(), topicID, aid); err != nil {
		return
	} else if member != nil {
		isMember = true
	}

	if isMember {
		return model.FollowStatusFollowed, nil
	}

	var req *model.TopicFollowRequest
	if req, err = p.d.GetTopicFollowRequest(c, p.d.DB(), topicID, aid); err != nil {
		return
	}

	if req != nil {
		switch req.Status {
		case model.FollowRequestStatusCommited:
			return model.FollowStatusApproving, nil
		case model.FollowRequestStatusApproved:
			return
		case model.FollowRequestStatusRejected:
			break
		}
	}

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, tx, topicID); err != nil {
		return
	} else if t == nil {
		return 0, ecode.TopicNotExist
	}

	var account *model.Account
	if account, err = p.getAccountByID(c, tx, aid); err != nil {
		return
	} else if account == nil {
		return 0, ecode.UserNotExist
	}

	switch t.JoinPermission {
	case model.JoinPermissionMember:
		return model.FollowStatusFollowed, p.addMember(c, tx, topicID, aid, model.MemberRoleUser)
	case model.JoinPermissionIDCert:
		if account.IDCert {
			return model.FollowStatusFollowed, p.addMember(c, tx, topicID, aid, model.MemberRoleUser)
		}
		return model.FollowStatusUnfollowed, ecode.NeedIDCert
	case model.JoinPermissionWorkCert:
		if account.IDCert && account.WorkCert {
			return model.FollowStatusFollowed, p.addMember(c, tx, topicID, aid, model.MemberRoleUser)
		}

		return model.FollowStatusUnfollowed, ecode.NeedWorkCert
	case model.JoinPermissionIDCertApprove:
		if !account.IDCert {
			return model.FollowStatusUnfollowed, ecode.NeedIDCert
		}
		break
	case model.JoinPermissionWorkCertApprove:
		if !account.IDCert || !account.WorkCert {
			return model.FollowStatusUnfollowed, ecode.NeedWorkCert
		}
		break
	case model.JoinPermissionAdminAdd:
		return model.FollowStatusUnfollowed, ecode.OnlyAllowAdminAdded
	case model.JoinPermissionPurchase:
		return model.FollowStatusUnfollowed, ecode.NeedPurchase
	case model.JoinPermissionVIP:
		if account.IsVIP {
			return model.FollowStatusFollowed, p.addMember(c, tx, topicID, aid, model.MemberRoleUser)
		}
		return model.FollowStatusUnfollowed, ecode.NeedPurchase
	}

	item := &model.TopicFollowRequest{
		ID:        gid.NewID(),
		Status:    model.FollowRequestStatusCommited,
		TopicID:   topicID,
		AccountID: aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddTopicFollowRequest(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return model.FollowStatusApproving, nil

}
