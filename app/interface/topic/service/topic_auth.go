package service

import (
	"context"

	"valerian/app/interface/topic/model"
)

func (p *Service) SaveAuthTopics(c context.Context, arg *model.ArgSaveAuthTopics) (err error) {
	return
}

func (p *Service) GetAuthTopics(c context.Context, topicID int64) (items []*model.AuthTopicResp, err error) {
	return
}

func (p *Service) GetUserCanEditTopics(c context.Context, query string, pn, ps int) (resp *model.CanEditTopicsResp, err error) {
	// aid, ok := metadata.Value(c, metadata.Aid).(int64)
	// if !ok {
	// 	err = ecode.AcquireAccountIDFailed
	// 	return
	// }

	// var ids []int64
	// if ids, err = p.d.GetUserCanEditTopicIDs(c, p.d.DB(), aid); err != nil {
	// 	return
	// }

	// var data *model.SearchResult
	// if data, err = p.d.TopicSearch(c, &model.TopicSearchParams{&model.BasicSearchParams{KW: query, Pn: pn, Ps: ps}}, ids); err != nil {
	// 	err = ecode.SearchTopicFailed
	// 	return
	// }

	// resp = &model.CanEditTopicsResp{
	// 	Items:  make([]*model.CanEditTopicItem, len(data.Result)),
	// 	Paging: &model.Paging{},
	// }

	// for i, v := range data.Result {
	// 	t := new(model.ESTopic)
	// 	err = json.Unmarshal(v, t)
	// 	if err != nil {
	// 		return
	// 	}
	// 	item := &model.CanEditTopicItem{
	// 		ID:             t.ID,
	// 		Name:           *t.Name,
	// 		Introduction:   *t.Introduction,
	// 		EditPermission: *t.EditPermission,
	// 		Avatar:         *t.Avatar,
	// 	}

	// 	var stat *model.TopicStat
	// 	if stat, err = p.GetTopicStat(c, t.ID); err != nil {
	// 		return
	// 	}

	// 	item.MemberCount = stat.MemberCount
	// 	item.ArticleCount = stat.ArticleCount
	// 	item.DiscussionCount = stat.DiscussionCount

	// 	if item.HasCatalogTaxonomy, err = p.d.HasTaxonomy(c, p.d.DB(), t.ID); err != nil {
	// 		return
	// 	}

	// 	resp.Items[i] = item
	// }

	// if resp.Paging.Prev, err = genURL("/api/v1/topic/list/has_edit_permission", url.Values{
	// 	"query": []string{query},
	// 	"pn":    []string{strconv.Itoa(pn - 1)},
	// 	"ps":    []string{strconv.Itoa(ps)},
	// }); err != nil {
	// 	return
	// }

	// if resp.Paging.Next, err = genURL("/api/v1/topic/list/has_edit_permission", url.Values{
	// 	"query": []string{query},
	// 	"pn":    []string{strconv.Itoa(pn + 1)},
	// 	"ps":    []string{strconv.Itoa(ps)},
	// }); err != nil {
	// 	return
	// }

	// if len(resp.Items) < ps {
	// 	resp.Paging.IsEnd = true
	// 	resp.Paging.Next = ""
	// }

	// if pn == 1 {
	// 	resp.Paging.Prev = ""
	// }

	return
}
