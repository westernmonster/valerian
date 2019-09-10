package service

import (
	"context"
	"encoding/json"
	"strings"
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
)

// func (p *Service) SearchTopics(c context.Context, query string, include string) (resp *model.TopicSearchResp, err error) {
// 	inc := includeParam(include)

// 	resp = &model.TopicSearchResp{
// 		Items: make([]*model.TopicSearchItem, 0),
// 	}

// 	resp.Paging = &model.TopicSearchPaging{
// 		IsEnd: true,
// 	}

// 	var items []*model.Topic
// 	if items, err = p.d.GetAllTopics(c, p.d.DB()); err != nil {
// 		return
// 	}

// 	for _, v := range items {
// 		item := &model.TopicSearchItem{
// 			ID:           v.ID,
// 			Avatar:        v.Avatar,
// 			Name:         v.Name,
// 			Introduction: v.Introduction,
// 		}

// 		if inc["items[*].versions"] {
// 			if item.Versions, err = p.getTopicVersionsResp(c, p.d.DB(), item.ID); err != nil {
// 				return
// 			}
// 		}

// 		if inc["items[*].has_catalog_taxonomy"] {
// 			if count, e := p.d.GetTopicCatalogsCountByCondition(c, p.d.DB(), map[string]interface{}{
// 				"type":     model.TopicCatalogTaxonomy,
// 				"topic_id": v.ID,
// 			}); e != nil {
// 				return nil, e
// 			} else if count > 0 {
// 				item.HasCatalogTaxonomy = true
// 			}

// 		}

// 		resp.Items = append(resp.Items, item)
// 	}

// 	return
// }

func includeParam(include string) (dic map[string]bool) {
	arr := strings.Split(include, ",")
	dic = make(map[string]bool)
	for _, v := range arr {
		dic[v] = true
	}

	return
}

func (p *Service) TopicSearch(c context.Context, arg *model.TopicSearchParams) (res *model.TopicSearchResult, err error) {
	var data *model.SearchResult
	if data, err = p.d.TopicSearch(c, arg); err != nil {
		err = ecode.SearchAccountFailed
	}

	res = &model.TopicSearchResult{
		Page:  data.Page,
		Debug: data.Debug,
		Data:  make([]*model.ESTopic, 0),
	}

	for _, v := range data.Result {
		acc := new(model.ESTopic)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		res.Data = append(res.Data, acc)
	}
	return
}
