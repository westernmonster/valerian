package service

import (
	"context"

	"valerian/app/interface/topic/model"
	account "valerian/app/service/account/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetTopicMeta(c context.Context, t *model.TopicResp) (meta *model.TopicMeta, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok || aid == 0 {
		return p.GetGuestTopicMeta(c, t)
	}

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, aid); err != nil {
		return
	} else if account == nil {
		return nil, ecode.UserNotExist
	}

	var isMember bool
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": t.ID}); err != nil {
		return
	} else if member != nil {
		isMember = true
	}

	meta = new(model.TopicMeta)

	meta.FollowStatus = model.FollowStatusUnfollowed

	if isMember {
		meta.CanView = true
		meta.IsMember = isMember
		meta.MemberRole = member.Role
		meta.FollowStatus = model.FollowStatusFollowed
	} else if t.ViewPermission == model.ViewPermissionPublic {
		meta.CanView = true
	}

	if !isMember {
		var req *model.TopicFollowRequest
		if req, err = p.d.GetTopicFollowRequest(c, p.d.DB(), t.ID, aid); err != nil {
			return
		} else if req != nil {
			if req.Status == model.FollowRequestStatusCommited {
				meta.FollowStatus = model.FollowStatusApproving
			}
		}
	}

	switch t.JoinPermission {
	case model.JoinPermissionMember:
		meta.CanFollow = true
		break
	case model.JoinPermissionCertApprove:
		if account.IDCert && account.WorkCert {
			meta.CanFollow = true
		}
		break
	case model.JoinPermissionMemberApprove:
		meta.CanFollow = true
		break
	case model.JoinPermissionManualAdd:
		meta.CanFollow = false
		break
	}

	switch t.EditPermission {
	case model.EditPermissionMember:
		if isMember {
			meta.CanEdit = true
		}
		break
	case model.EditPermissionAdmin:
		if isMember && (member.Role == model.MemberRoleAdmin || member.Role == model.MemberRoleOwner) {
			meta.CanEdit = true
		}
		break
	}

	if isMember && (member.Role == model.MemberRoleAdmin || member.Role == model.MemberRoleOwner) {
		meta.CanEdit = true
	}

	var setting *model.AccountTopicSetting
	if setting, err = p.getAccountTopicSetting(c, p.d.DB(), aid, t.ID); err != nil {
		return
	} else if setting == nil {
		meta.Fav = false
	} else {
		meta.Fav = bool(setting.Fav)
	}

	return
}

func (p *Service) GetGuestTopicMeta(c context.Context, t *model.TopicResp) (meta *model.TopicMeta, err error) {
	meta = new(model.TopicMeta)
	if t.ViewPermission == model.ViewPermissionPublic {
		meta.CanView = true
	}

	meta.CanView = false
	meta.CanEdit = false
	meta.FollowStatus = model.FollowStatusUnfollowed
	meta.IsMember = false
	meta.MemberRole = ""
	meta.CanFollow = false

	return
}
