package service

import (
	"context"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) IsSystemAdmin(c context.Context, aid int64) (ret bool, err error) {
	return p.isSystemAdmin(c, p.d.DB(), aid)
}

func (p *Service) isSystemAdmin(c context.Context, node sqalx.Node, aid int64) (ret bool, err error) {
	var acc *model.Account
	if acc, err = p.getAccount(c, node, aid); err != nil {
		return
	}

	if acc.Role == "admin" || acc.Role == "superadmin" {
		ret = true
		return
	}

	return
}

func (p *Service) CanView(c context.Context, aid int64, topicID int64) (canView bool, err error) {
	return p.canView(c, p.d.DB(), aid, topicID)
}

func (p *Service) canView(c context.Context, node sqalx.Node, aid int64, topicID int64) (canView bool, err error) {
	var isSystemAdmin bool
	if isSystemAdmin, err = p.isSystemAdmin(c, node, aid); err != nil {
		return
	} else if isSystemAdmin {
		canView = true
		return
	}

	var t *model.Topic
	if t, err = p.getTopic(c, node, topicID); err != nil {
		return
	} else if t.ViewPermission == model.ViewPermissionPublic {
		canView = true
		return
	}

	if canView, err = p.d.IsAllowedViewMember(c, node, aid, topicID); err != nil {
		return
	}

	return
}

func (p *Service) CanEdit(c context.Context, aid int64, topicID int64) (canEdit bool, err error) {
	return p.canEdit(c, p.d.DB(), aid, topicID)
}

func (p *Service) canEdit(c context.Context, node sqalx.Node, aid int64, topicID int64) (canEdit bool, err error) {
	var isSystemAdmin bool
	if isSystemAdmin, err = p.isSystemAdmin(c, node, aid); err != nil {
		return
	} else if isSystemAdmin {
		canEdit = true
		return
	}

	if canEdit, err = p.d.IsAllowedEditMember(c, node, aid, topicID); err != nil {
		return
	}

	return
}

func (p *Service) IsTopicMember(c context.Context, aid int64, topicID int64) (ret bool, err error) {
	return p.isTopicMember(c, p.d.DB(), aid, topicID)
}

func (p *Service) isTopicMember(c context.Context, node sqalx.Node, aid int64, topicID int64) (ret bool, err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member != nil {
		ret = true
	}
	return
}

func (p *Service) IsTopicAdmin(c context.Context, aid int64, topicID int64) (ret bool, err error) {
	return p.isTopicAdmin(c, p.d.DB(), aid, topicID)
}

func (p *Service) isTopicAdmin(c context.Context, node sqalx.Node, aid int64, topicID int64) (ret bool, err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member != nil {
		if member.Role != model.MemberRoleUser {
			ret = true
			return
		}
	}
	return
}

func (p *Service) IsTopicOwner(c context.Context, aid int64, topicID int64) (ret bool, err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member != nil {
		if member.Role == model.MemberRoleOwner {
			ret = true
		}
	}
	return
}

func (p *Service) checkTopicManagePermission(c context.Context, aid, topicID int64) (err error) {
	var hasManagePermission bool
	if hasManagePermission, err = p.hasTopicManagePermission(c, p.d.DB(), aid, topicID); err != nil {
		return
	} else if !hasManagePermission {
		err = ecode.NoTopicManagePermission
	}

	return
}

func (p *Service) hasTopicManagePermission(c context.Context, node sqalx.Node, aid, topicID int64) (ret bool, err error) {
	var isSysAdmin bool
	if isSysAdmin, err = p.isSystemAdmin(c, node, aid); err != nil {
		return
	} else if isSysAdmin {
		ret = true
		return
	}

	var isTopicAdmin bool
	if isTopicAdmin, err = p.isTopicAdmin(c, node, aid, topicID); err != nil {
		return
	} else if isTopicAdmin {
		ret = true
		return
	}

	return
}

func (p *Service) checkIsTopicAdmin(c context.Context, aid, topicID int64) (err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member != nil {
		if member.Role == model.MemberRoleUser {
			err = ecode.NoTopicManagePermission
			return
		}
	}

	return
}

func (p *Service) checkIsMember(c context.Context, aid, topicID int64) (err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member == nil {
		err = ecode.NotTopicMember
		return
	}

	return
}

func (p *Service) checkViewPermission(c context.Context, aid, topicID int64) (err error) {
	var canView bool
	if canView, err = p.canView(c, p.d.DB(), aid, topicID); err != nil {
		return
	}

	if !canView {
		err = ecode.NoTopicViewPermission
		return
	}

	return
}

func (p *Service) checkEditPermission(c context.Context, aid, topicID int64) (err error) {
	var canEdit bool
	if canEdit, err = p.canEdit(c, p.d.DB(), aid, topicID); err != nil {
		return
	}

	if !canEdit {
		err = ecode.NoTopicEditPermission
		return
	}

	return
}
