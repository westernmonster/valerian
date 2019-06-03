package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetTopicMeta(c context.Context, t *model.TopicResp) (meta *model.TopicMeta, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok || aid == 0 {
		return p.GetGuestTopicMeta(c, t)
	}

	var account *model.Account
	if account, err = p.getAccountByID(c, p.d.DB(), aid); err != nil {
		return
	} else if account == nil {
		return nil, ecode.UserNotExist
	}

	var isMember bool
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCondition(c, p.d.DB(), t.ID, aid); err != nil {
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
	case model.JoinPermissionIDCert:
		if account.IDCert {
			meta.CanFollow = true
		}
		break
	case model.JoinPermissionWorkCert:
		if account.IDCert && account.WorkCert {
			meta.CanFollow = true
		}
		break
	case model.JoinPermissionMemberApprove:
		meta.CanFollow = true
		break
	case model.JoinPermissionIDCertApprove:
		if account.IDCert {
			meta.CanFollow = true
		}
		break
	case model.JoinPermissionWorkCertApprove:
		if account.IDCert && account.WorkCert {
			meta.CanFollow = true
		}
		break
	case model.JoinPermissionAdminAdd:
		meta.CanFollow = false
		break
	case model.JoinPermissionPurchase:
		break
	case model.JoinPermissionVIP:
		if account.IsVIP {
			meta.CanFollow = true
		}
		break
	}

	switch t.EditPermission {
	case model.EditPermissionIDCert:
		if account.IDCert {
			meta.CanEdit = true
		}
		break
	case model.EditPermissionWorkCert:
		if account.IDCert && account.WorkCert {
			meta.CanEdit = true
		}
		break
	case model.EditPermissionIDCertJoined:
		if bool(account.IDCert) && isMember {
			meta.CanEdit = true
		}
		break
	case model.EditPermissionWorkCertJoined:
		if bool(account.IDCert) && bool(account.WorkCert) && isMember {
			meta.CanEdit = true
		}
		break
	case model.EditPermissionApprovedIDCertJoined:
		if bool(account.IDCert) && isMember {
			meta.CanEdit = true
		}
		break
	case model.EditPermissionApprovedWorkCertJoined:
		if bool(account.IDCert) && bool(account.WorkCert) && isMember {
			meta.CanEdit = true
		}
		break
		// case model.EditPermissionAdmin:
		// 	if isMember && (member.Role == model.MemberRoleAdmin || member.Role == model.MemberRoleOwner) {
		// 		meta.CanEdit = true
		// 	}
		// 	break
	}

	if isMember && (member.Role == model.MemberRoleAdmin || member.Role == model.MemberRoleOwner) {
		meta.CanEdit = true
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
