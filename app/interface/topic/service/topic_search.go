package service

import (
	"context"
	"strings"
	"valerian/app/interface/topic/model"
)

func (p *Service) SearchTopics(c context.Context, query string, include string) (resp *model.TopicSearchResp, err error) {
	inc := includeParam(include)

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

		if inc["items[*].versions"] {
			if item.Versions, err = p.getTopicVersionsResp(c, p.d.DB(), item.TopicSetID); err != nil {
				return
			}
		}

		if inc["items[*].has_catalog_taxonomy"] {
			if count, e := p.d.GetTopicCatalogsCountByCondition(c, p.d.DB(), map[string]interface{}{
				"type":     model.TopicCatalogTaxonomy,
				"topic_id": v.ID,
			}); e != nil {
				return nil, e
			} else if count > 0 {
				item.HasCatalogTaxonomy = true
			}

		}

		resp.Items = append(resp.Items, item)
	}

	return
}

func includeParam(include string) (dic map[string]bool) {
	arr := strings.Split(include, ",")

	dic = make(map[string]bool)
	for _, v := range arr {
		dic[v] = true
	}

	return
}
