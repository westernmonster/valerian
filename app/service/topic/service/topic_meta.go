package service

import (
	"context"

	account "valerian/app/service/account/api"
	"valerian/app/service/topic/model"
	"valerian/library/ecode"
)

func (p *Service) GetTopicMeta(c context.Context, aid, topicID int64) (meta *model.TopicMeta, err error) {
	var t *model.Topic
	if t, err = p.getTopic(c, p.d.DB(), topicID); err != nil {
		return
	}

	if aid == 0 {
		return p.GetGuestTopicMeta(c, t)
	}

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, aid); err != nil {
		return
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

	var setting *model.AccountTopicSetting
	if setting, err = p.getAccountTopicSetting(c, p.d.DB(), aid, t.ID); err != nil {
		return
	} else if setting == nil {
		meta.Fav = false
	} else {
		meta.Fav = bool(setting.Fav)
	}

	if meta.CanView, meta.CanEdit, err = p.CheckEditPermission(c, aid, member, t); err != nil {
		return
	}

	return
}

func (p *Service) GetGuestTopicMeta(c context.Context, t *model.Topic) (meta *model.TopicMeta, err error) {
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

func (p *Service) GetTopicPermission(c context.Context, aid, topicID int64) (isMember bool, role string, editPermission string, err error) {
	var t *model.Topic
	if t, err = p.getTopic(c, p.d.DB(), topicID); err != nil {
		return
	}

	editPermission = t.EditPermission

	if aid == 0 {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member != nil {
		isMember = true
		role = member.Role
	}

	return
}

func (p *Service) CheckEditPermission(c context.Context, aid int64, member *model.TopicMember, t *model.Topic) (canView, canEdit bool, err error) {
	// 1. 是否当前话题成员，且当前话题的权限是否允许其编辑
	switch t.EditPermission {
	case model.EditPermissionMember:
		if member != nil {
			// 是话题成员
			// 话题权限为成员编辑
			canEdit = true
			return
		}
		break
	case model.EditPermissionAdmin:
		if member != nil && (member.Role == model.MemberRoleAdmin || member.Role == model.MemberRoleOwner) {
			// 是话题成员
			// 话题权限为管理员编辑
			canEdit = true
			return
		}
		break
	}

	// 2. 授权话题的成员列表中 查看是否有编辑权限
	var authTopics []*model.AuthTopic
	if authTopics, err = p.d.GetAuthTopicsByCond(c, p.d.DB(), map[string]interface{}{"topic_id": t.ID}); err != nil {
		return
	}

	for _, v := range authTopics {
		var m *model.TopicMember
		if m, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": v.ToTopicID}); err != nil {
			return
		} else if m != nil {
			switch v.Permission {
			case model.AuthPermissionView:
				// 授权为成员可查看
				canView = true
				return
			case model.AuthPermissionEdit:
				// 授权为成员可编辑
				canEdit = true
				canView = true
				return
			case model.AuthPermissionAdminEdit:
				// 授权为管理员可编辑
				if m.Role == model.MemberRoleAdmin || m.Role == model.MemberRoleOwner {
					canView = true
					canEdit = true
					return
				}
			}
		}
	}

	return

}
