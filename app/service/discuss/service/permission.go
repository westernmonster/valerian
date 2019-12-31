package service

import (
	"context"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// CanEdit 是否能编辑讨论
func (p *Service) CanEdit(c context.Context, arg *api.IDReq) (canEdit bool, err error) {
	if canEdit, err = p.canEdit(c, p.d.DB(), arg.Aid, arg.ID); err != nil {
		return
	}

	return
}

// CanView 是否能查看讨论
func (p *Service) CanView(c context.Context, arg *api.IDReq) (canView bool, err error) {
	if canView, err = p.canView(c, p.d.DB(), arg.Aid, arg.ID); err != nil {
		return
	}

	return
}

// checkTopicExist 获取话题信息
func (p *Service) checkTopicExist(c context.Context, node sqalx.Node, topicID int64) (err error) {
	var exist bool
	if exist, err = p.d.IsTopicExist(c, node, topicID); err != nil {
		return
	} else if !exist {
		err = ecode.TopicNotExist
		return
	}
	return
}

// checkIsTopicMember 检测是否话题成员
func (p *Service) checkIsTopicMember(c context.Context, node sqalx.Node, aid, topicID int64) (err error) {
	var isMember bool
	if isMember, _, err = p.d.IsTopicMember(c, node, aid, topicID); err != nil {
		return
	} else if !isMember {
		err = ecode.NotTopicMember
		return
	}

	return
}

// checkIsTopicMember 检测是否话题管理员
func (p *Service) checkIsTopicManager(c context.Context, node sqalx.Node, aid, topicID int64) (err error) {
	var isMember bool
	var role string
	if isMember, role, err = p.d.IsTopicMember(c, node, aid, topicID); err != nil {
		return
	} else if !isMember {
		err = ecode.NotTopicMember
		return
	}

	if role == model.MemberRoleUser {
		err = ecode.NotTopicAdmin
		return
	}

	return
}

// isTopicManager 是否话题管理员
func (p *Service) isTopicManager(c context.Context, node sqalx.Node, aid int64, topicID int64) (ret bool, err error) {
	var isMember bool
	var role string
	if isMember, role, err = p.d.IsTopicMember(c, node, aid, topicID); err != nil {
		return
	} else if !isMember {
		return
	}

	if role == model.MemberRoleAdmin || role == model.MemberRoleOwner {
		ret = true
		return
	}

	return
}

// isSystemAdmin 是否系统管理员
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

// checkEditPermission 检查编辑权限
func (p *Service) checkEditPermission(c context.Context, node sqalx.Node, aid, discussionID int64) (err error) {
	var canEdit bool
	if canEdit, err = p.canEdit(c, node, aid, discussionID); err != nil {
		return
	}

	if !canEdit {
		err = ecode.NoDiscussionEditPermission
	}

	return
}

// checkViewPermission 检测查看权限
func (p *Service) checkViewPermission(c context.Context, node sqalx.Node, aid, discussionID int64) (err error) {
	var canView bool
	if canView, err = p.canView(c, node, aid, discussionID); err != nil {
		return
	}

	if !canView {
		err = ecode.NoDiscussionViewPermission
	}

	return
}

// canView 是否能查看
func (p *Service) canView(c context.Context, node sqalx.Node, aid int64, discussionID int64) (canView bool, err error) {
	var isSystemAdmin bool
	if isSystemAdmin, err = p.isSystemAdmin(c, node, aid); err != nil {
		return
	} else if isSystemAdmin {
		canView = true
		return
	}

	if canView, err = p.d.IsAllowedViewMember(c, node, aid, discussionID); err != nil {
		return
	}

	var item *model.Discussion
	if item, err = p.getDiscussion(c, node, discussionID); err != nil {
		return
	}
	if item.CreatedBy == aid {
		canView = true
	}

	var viewPermission string
	if viewPermission, err = p.d.GetTopicViewPermissionByID(c, node, item.TopicID); err != nil {
		return
	}
	if viewPermission == model.ViewPermissionPublic {
		canView = true
		return
	}

	return
}

// canEdit 是否能编辑
func (p *Service) canEdit(c context.Context, node sqalx.Node, aid int64, discussionID int64) (canEdit bool, err error) {
	// 系统管理员
	var isSystemAdmin bool
	if isSystemAdmin, err = p.isSystemAdmin(c, node, aid); err != nil {
		return
	} else if isSystemAdmin {
		canEdit = true
		return
	}

	// 话题管理员
	var item *model.Discussion
	if item, err = p.getDiscussion(c, node, discussionID); err != nil {
		return
	}
	if item.CreatedBy == aid {
		canEdit = true
	}

	var isTopicManager bool
	if isTopicManager, err = p.isTopicManager(c, node, aid, item.TopicID); err != nil {
		return
	} else if isTopicManager {
		canEdit = true
	}

	return
}
