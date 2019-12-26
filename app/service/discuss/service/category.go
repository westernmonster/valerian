package service

import (
	"context"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// SaveDiscussCategories 批量保存讨论分类
func (p *Service) SaveDiscussCategories(c context.Context, arg *api.ArgSaveDiscussCategories) (err error) {
	return
}

// getDiscussCategories 获取话题所有讨论分类
func (p *Service) getDiscussCategories(c context.Context, node sqalx.Node, topicID int64) (items []*model.DiscussCategory, err error) {
	var addCache = true
	if items, err = p.d.DiscussionCategoriesCache(c, topicID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetDiscussCategoriesByCond(c, p.d.DB(), map[string]interface{}{"topic_id": topicID}); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetDiscussionCategoriesCache(context.TODO(), topicID, items)
		})
	}

	return
}

// GetDiscussCategories 获取话题所有讨论分类
func (p *Service) GetDiscussCategories(c context.Context, topicID int64) (items []*model.DiscussCategory, err error) {
	return p.getDiscussCategories(c, p.d.DB(), topicID)
}

// GetDiscussCategories 获取指定讨论分类
func (p *Service) GetDiscussCategory(c context.Context, categoryID int64) (item *model.DiscussCategory, err error) {
	if item, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), categoryID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussCategoryNotExist
		return
	}
	return
}
