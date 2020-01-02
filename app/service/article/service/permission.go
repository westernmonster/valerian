package service

import (
	"context"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// CanEdit 是否能编辑文章
func (p *Service) CanEdit(c context.Context, arg *api.IDReq) (canEdit bool, err error) {
	if canEdit, err = p.canEdit(c, p.d.DB(), arg.Aid, arg.ID); err != nil {
		return
	}

	return
}

// CanView 是否能查看文章
func (p *Service) CanView(c context.Context, arg *api.IDReq) (canView bool, err error) {
	if canView, err = p.canView(c, p.d.DB(), arg.Aid, arg.ID); err != nil {
		return
	}

	return
}

// GetUserCanEditArticleIDs 获取用户能编辑的文章ID列表
func (p *Service) GetUserCanEditArticleIDs(c context.Context, aid int64) (ids []int64, err error) {
	if ids, err = p.d.GetUserCanEditArticleIDs(c, p.d.DB(), aid); err != nil {
		return
	}

	return
}

// checkEditPermission 检查编辑权限
func (p *Service) checkEditPermission(c context.Context, node sqalx.Node, aid, articleID int64) (err error) {
	var canEdit bool
	if canEdit, err = p.canEdit(c, node, aid, articleID); err != nil {
		return
	}

	if !canEdit {
		err = ecode.NoArticleEditPermission
	}

	return
}

// checkViewPermission 检测查看权限
func (p *Service) checkViewPermission(c context.Context, node sqalx.Node, aid, articleID int64) (err error) {
	var canView bool
	if canView, err = p.canView(c, node, aid, articleID); err != nil {
		return
	}

	if !canView {
		err = ecode.NoArticleViewPermission
	}

	return
}

// IsSystemAdmin 是否系统管理员
func (p *Service) IsSystemAdmin(c context.Context, aid int64) (ret bool, err error) {
	return p.isSystemAdmin(c, p.d.DB(), aid)
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

// canView 是否能查看
func (p *Service) canView(c context.Context, node sqalx.Node, aid int64, articleID int64) (canView bool, err error) {
	var isSystemAdmin bool
	if isSystemAdmin, err = p.isSystemAdmin(c, node, aid); err != nil {
		return
	} else if isSystemAdmin {
		canView = true
		return
	}

	if canView, err = p.d.IsAllowedViewMember(c, node, aid, articleID); err != nil {
		return
	}

	var catalogs []*model.TopicCatalog
	if catalogs, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"ref_id": articleID,
		"type":   model.TopicCatalogArticle,
	}); err != nil {
		return
	}

	for _, v := range catalogs {
		var viewPermission string
		if viewPermission, err = p.d.GetTopicViewPermissionByID(c, node, v.TopicID); err != nil {
			return
		}
		if viewPermission == model.ViewPermissionPublic {
			canView = true
			return
		}
	}

	var article *model.Article
	if article, err = p.getArticle(c, node, articleID); err != nil {
		return
	}

	if article.CreatedBy == aid {
		canView = true
	}

	return
}

// canEdit 是否能编辑
func (p *Service) canEdit(c context.Context, node sqalx.Node, aid int64, articleID int64) (canEdit bool, err error) {
	var isSystemAdmin bool
	if isSystemAdmin, err = p.isSystemAdmin(c, node, aid); err != nil {
		return
	} else if isSystemAdmin {
		canEdit = true
		return
	}

	if canEdit, err = p.d.IsAllowedEditMember(c, node, aid, articleID); err != nil {
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, node, articleID); err != nil {
		return
	}

	if article.CreatedBy == aid {
		canEdit = true
	}

	return
}
