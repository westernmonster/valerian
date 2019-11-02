package service

import (
	"context"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
)

// 检查编辑权限
// TODO：这里会成为性能瓶颈，业务目前是这么制定的，我也没辙，后续人员请持续优化
func (p *Service) checkEditPermission(c context.Context, node sqalx.Node, articleID, aid int64) (canEdit bool, err error) {
	var acc *model.Account
	if acc, err = p.getAccount(c, node, aid); err != nil {
		return
	}

	// 如果是系统管理员，则可编辑
	if acc.Role == model.UserRoleAdmin || acc.Role == model.UserRoleSuperAdmin {
		canEdit = true
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, node, articleID); err != nil {
		return
	}

	// 如果是该文章创建人，则可以编辑
	if article.CreatedBy == aid {
		canEdit = true
		return
	}

	// 所有关联话题列表

	var relatedTopics []*model.TopicCatalog
	if relatedTopics, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"type":   model.TopicCatalogArticle,
		"ref_id": articleID,
	}); err != nil {
		return
	}

	for _, v := range relatedTopics {
		// 如果不是主关联话题，查看该话题的编辑权限
		//     话题的编辑权限为管理员，则检查当前用户是否是管理员
		//     话题的编辑权限为成员，则检查当前用户是否是成员
		//
		if !v.IsPrimary && v.Permission == model.AuthPermissionEdit {
			var topicEditPermission string
			if topicEditPermission, err = p.d.GetTopicEditPermissionByID(c, node, v.TopicID); err != nil {
				return
			}

			// 如果
			switch topicEditPermission {
			case model.EditPermissionAdmin:
				var isAdmin bool
				if isAdmin, err = p.d.IsTopicAdmin(c, node, v.TopicID, aid); err != nil {
					return
				} else if isAdmin {
					canEdit = true
					return
				}
				break
			case model.EditPermissionMember:
				var isMember bool
				if isMember, err = p.d.IsTopicMember(c, node, v.TopicID, aid); err != nil {
					return
				} else if isMember {
					canEdit = true
					return
				}
				break
			}

			continue
		}

		// 如果是主关联话题
		//     主话题的编辑权限为管理员，则检查当前用户是否是管理员
		//     主话题的编辑权限为成员，则检查当前用户是否是成员
		//
		var topicEditPermission string
		if topicEditPermission, err = p.d.GetTopicEditPermissionByID(c, node, v.TopicID); err != nil {
			return
		}

		switch topicEditPermission {
		case model.EditPermissionAdmin:
			var isAdmin bool
			if isAdmin, err = p.d.IsTopicAdmin(c, node, v.TopicID, aid); err != nil {
				return
			} else if isAdmin {
				canEdit = true
				return
			}
			break
		case model.EditPermissionMember:
			var isMember bool
			if isMember, err = p.d.IsTopicMember(c, node, v.TopicID, aid); err != nil {
				return
			} else if isMember {
				canEdit = true
				return
			}
			break
		}

		// 检查主话题的授权话题列表
		//     授权为管理员可编辑，则检查是否该授权话题的管理员
		//     授权为所有人可以编辑，则检查是否该话题的成员
		var authTopics []*model.AuthTopic
		if authTopics, err = p.d.GetAuthTopicsByCond(c, node, map[string]interface{}{"topic_id": v.TopicID}); err != nil {
			return
		}
		for _, x := range authTopics {
			switch x.Permission {
			case model.AuthPermissionAdminEdit:
				var isAdmin bool
				if isAdmin, err = p.d.IsTopicAdmin(c, node, x.ToTopicID, aid); err != nil {
					return
				} else if isAdmin {
					canEdit = true
					return
				}
				break
			case model.AuthPermissionEdit:
				var isMember bool
				if isMember, err = p.d.IsTopicMember(c, node, x.ToTopicID, aid); err != nil {
					return
				} else if isMember {
					canEdit = true
					return
				}
				break
			}
		}
	}

	return
}
