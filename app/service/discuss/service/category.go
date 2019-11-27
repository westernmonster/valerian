package service

import (
	"context"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

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

func (p *Service) GetDiscussCategories(c context.Context, topicID int64) (items []*model.DiscussCategory, err error) {
	return p.getDiscussCategories(c, p.d.DB(), topicID)
}

func (p *Service) GetDiscussCategory(c context.Context, categoryID int64) (item *model.DiscussCategory, err error) {
	if item, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), categoryID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussCategoryNotExist
		return
	}
	return
}
