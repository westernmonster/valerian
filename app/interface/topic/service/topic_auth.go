package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"valerian/app/interface/topic/model"
	search "valerian/app/service/search/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) SaveAuthTopics(c context.Context, arg *model.ArgSaveAuthTopics) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	item := &topic.ArgSaveAuthTopics{
		TopicID:    arg.TopicID,
		Aid:        aid,
		AuthTopics: make([]*topic.ArgAuthTopic, 0),
	}

	if arg.AuthTopics != nil {
		for _, v := range arg.AuthTopics {
			item.AuthTopics = append(item.AuthTopics, &topic.ArgAuthTopic{
				TopicID:    v.TopicID,
				Permission: v.Permission,
			})
		}
	}

	if err = p.d.SaveAuthTopics(c, item); err != nil {
		return
	}
	return
}

func (p *Service) GetAuthTopics(c context.Context, topicID int64) (items []*model.AuthTopicResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	// 检测查看权限
	if err = p.checkViewPermission(c, aid, topicID); err != nil {
		return
	}

	return p.getAuthTopics(c, aid, topicID)
}

func (p *Service) getAuthTopics(c context.Context, aid, topicID int64) (items []*model.AuthTopicResp, err error) {
	var resp *topic.AuthTopicsResp
	if resp, err = p.d.GetAuthTopics(c, &topic.IDReq{ID: topicID, Aid: aid}); err != nil {
		return
	}

	items = make([]*model.AuthTopicResp, 0)

	if resp.Items == nil {
		return
	}

	for _, v := range resp.Items {
		items = append(items, &model.AuthTopicResp{
			ToTopicID:      v.ToTopicID,
			EditPermission: v.EditPermission,
			Permission:     v.Permission,
			MemberCount:    v.MemberCount,
			Avatar:         v.Avatar,
			Name:           v.Name,
		})
	}

	return
}

func (p *Service) GetUserCanEditTopics(c context.Context, query string, pn, ps int) (resp *model.CanEditTopicsResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var idsResp *topic.IDsResp
	if idsResp, err = p.d.GetUserCanEditTopicIDs(c, &topic.AidReq{AccountID: aid}); err != nil {
		return
	}

	if idsResp.IDs == nil || len(idsResp.IDs) == 0 {
		resp = &model.CanEditTopicsResp{
			Items:  make([]*model.CanEditTopicItem, 0),
			Paging: &model.Paging{IsEnd: true},
		}
		return
	}

	var data *search.SearchResult
	if data, err = p.d.SearchTopic(c, &search.SearchParam{KW: query, Sort: []string{"id"}, Order: []string{"desc"}, Pn: int32(pn), Ps: int32(ps), IDs: idsResp.IDs}); err != nil {
		err = ecode.SearchTopicFailed
		return
	}

	resp = &model.CanEditTopicsResp{
		Items:  make([]*model.CanEditTopicItem, len(data.Result)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Result {
		t := new(model.ESTopic)
		err = json.Unmarshal(v, t)
		if err != nil {
			return
		}
		item := &model.CanEditTopicItem{
			ID:             t.ID,
			Name:           *t.Name,
			Introduction:   *t.Introduction,
			EditPermission: *t.EditPermission,
			Avatar:         *t.Avatar,
		}

		var stat *topic.TopicStat
		if stat, err = p.d.GetTopicStat(c, &topic.TopicReq{ID: t.ID}); err != nil {
			return
		}

		item.MemberCount = stat.MemberCount
		item.ArticleCount = stat.ArticleCount
		item.DiscussionCount = stat.DiscussionCount

		if item.HasCatalogTaxonomy, err = p.d.HasTaxonomy(c, &topic.TopicReq{ID: t.ID}); err != nil {
			return
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/topic/list/has_edit_permission", url.Values{
		"query": []string{query},
		"pn":    []string{strconv.Itoa(pn - 1)},
		"ps":    []string{strconv.Itoa(ps)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/topic/list/has_edit_permission", url.Values{
		"query": []string{query},
		"pn":    []string{strconv.Itoa(pn + 1)},
		"ps":    []string{strconv.Itoa(ps)},
	}); err != nil {
		return
	}

	if len(resp.Items) < ps {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if pn == 1 {
		resp.Paging.Prev = ""
	}

	return
}
