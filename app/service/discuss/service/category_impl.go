package service

import (
	"context"

	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// loadDiscussCategoriesMap 获取分类字典
func (p *Service) loadDiscussCategoriesMap(c context.Context, node sqalx.Node, topicID int64) (dic map[int64]bool, err error) {
	dic = make(map[int64]bool)
	var dbItems []*model.DiscussCategory
	if dbItems, err = p.d.GetDiscussCategoriesByCond(c, node, map[string]interface{}{"topic_id": topicID}); err != nil {
		return
	}

	for _, v := range dbItems {
		dic[v.ID] = false
	}

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

// checkCategory 验证讨论
func (p *Service) checkCategory(c context.Context, node sqalx.Node, categoryID int64) (err error) {
	// -1 表示「问答」
	if categoryID == -1 {
		return
	}

	var t *model.DiscussCategory
	if t, err = p.d.GetDiscussCategoryByID(c, node, categoryID); err != nil {
		return
	} else if t == nil {
		return ecode.DiscussCategoryNotExist
	}

	return
}
