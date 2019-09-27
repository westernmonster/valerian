package service

import (
	"context"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
)

func (p *Service) getDiscussCategories(c context.Context, node sqalx.Node, topicID int64) (items []*model.DiscussCategory, err error) {
	items = make([]*model.DiscussCategory, 0)
	if items, err = p.d.GetDiscussCategoriesByCond(c, p.d.DB(), map[string]interface{}{"topic_id": topicID}); err != nil {
		return
	}

	return
}

func (p *Service) GetDiscussCategories(c context.Context, topicID int64) (items []*model.DiscussCategory, err error) {
	return p.getDiscussCategories(c, p.d.DB(), topicID)
}
