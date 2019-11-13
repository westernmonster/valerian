package service

import (
	"context"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/ecode"
)

func (p *Service) GetTopicMeta(c context.Context, aid, topicID int64) (meta *api.TopicMetaInfo, err error) {
	var t *model.Topic
	if t, err = p.getTopic(c, p.d.DB(), topicID); err != nil {
		return
	}

	if aid == 0 {
		return p.GetGuestTopicMeta(c, t)
	}

	var account *model.Account
	if account, err = p.getAccount(c, p.d.DB(), aid); err != nil {
		return
	}

	var isMember bool
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": t.ID}); err != nil {
		return
	} else if member != nil {
		isMember = true
	}

	meta = new(api.TopicMetaInfo)

	meta.FollowStatus = model.FollowStatusUnfollowed

	if isMember {
		meta.IsMember = isMember
		meta.MemberRole = member.Role
		meta.FollowStatus = model.FollowStatusFollowed
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

	if t.ViewPermission == model.ViewPermissionPublic {
		meta.CanView = true
	} else {
		if meta.CanView, err = p.CanView(c, aid, topicID); err != nil {
			return
		}
	}

	if meta.CanEdit, err = p.CanEdit(c, aid, topicID); err != nil {
		return
	}

	if meta.Fav, err = p.isFav(c, p.d.DB(), aid, topicID, model.TargetTypeTopic); err != nil {
		return
	}

	return
}

func (p *Service) GetGuestTopicMeta(c context.Context, t *model.Topic) (meta *api.TopicMetaInfo, err error) {
	meta = new(api.TopicMetaInfo)
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
