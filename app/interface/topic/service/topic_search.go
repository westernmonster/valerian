package service

import (
	"context"
	"valerian/app/interface/topic/model"
)

func (p *Service) SearchTopics(c context.Context, query string) (resp *model.TopicSearchResp, err error) {
	resp = &model.TopicSearchResp{
		Items: make([]*model.TopicSearchItem, 0),
	}

	resp.Paging = &model.TopicSearchPaging{
		IsEnd: true,
	}

	var items []*model.Topic
	if items, err = p.d.GetAllTopics(c, p.d.DB()); err != nil {
		return
	}

	for _, v := range items {
		item := &model.TopicSearchItem{
			ID:           v.ID,
			Cover:        v.Cover,
			Name:         v.Name,
			Introduction: v.Introduction,
			TopicSetID:   v.TopicSetID,
			VersionName:  v.VersionName,
		}

		if item.Versions, err = p.getTopicVersionsResp(c, p.d.DB(), item.TopicSetID); err != nil {
			return
		}

		resp.Items = append(resp.Items, item)
	}

	return
}
