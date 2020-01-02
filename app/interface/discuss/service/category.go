package service

import (
	"context"

	"valerian/app/interface/discuss/model"
	discussion "valerian/app/service/discuss/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) getDiscussCategories(c context.Context, aid, topicID int64) (items []*model.DiscussCategoryResp, err error) {
	var resp *discussion.CategoriesResp
	if resp, err = p.d.GetDiscussionCategories(c, aid, topicID); err != nil {
		return
	}

	items = make([]*model.DiscussCategoryResp, len(resp.Items))
	for i, v := range resp.Items {
		items[i] = &model.DiscussCategoryResp{
			ID:      v.ID,
			TopicID: v.TopicID,
			Name:    v.Name,
			Seq:     v.Seq,
		}
	}

	return
}

func (p *Service) GetDiscussCategories(c context.Context, topicID int64) (items []*model.DiscussCategoryResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	return p.getDiscussCategories(c, aid, topicID)
}

func (p *Service) SaveDiscussCategories(c context.Context, arg *model.ArgSaveDiscussCategories) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &discussion.ArgSaveDiscussCategories{
		Aid:     aid,
		TopicID: arg.TopicID,
		Items:   make([]*discussion.ArgDisucssCategory, 0),
	}

	for _, v := range arg.Items {
		item := &discussion.ArgDisucssCategory{
			Name: v.Name,
			Seq:  v.Seq,
		}

		if v.ID != nil {
			item.ID = &discussion.ArgDisucssCategory_IDValue{*v.ID}
		}

		req.Items = append(req.Items, item)
	}

	if err = p.d.SaveDiscussionCategories(c, req); err != nil {
		return
	}

	return
}
