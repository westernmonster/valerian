package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"valerian/app/admin/topic/model"
	search "valerian/app/service/search/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) Follow(c context.Context, arg *model.ArgTopicFollow) (status int32, err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var resp *topic.StatusResp
	if resp, err = p.d.FollowTopic(c, &topic.ArgTopicFollow{TopicID: arg.TopicID, AllowViewCert: arg.AllowViewCert, Reason: arg.Reason, Aid: aid}); err != nil {
		return
	}

	status = resp.Status
	return
}

func (p *Service) FollowedTopics(c context.Context, query string, pn, ps int) (resp *model.JoinedTopicsResp, err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var idsResp *topic.IDsResp
	if idsResp, err = p.d.GetFollowedTopicsIDs(c, &topic.AidReq{AccountID: aid}); err != nil {
		return
	}

	var data *search.SearchResult
	if data, err = p.d.SearchTopic(c, &search.SearchParam{KW: query, Pn: int32(pn), Ps: int32(ps), IDs: idsResp.IDs}); err != nil {
		err = ecode.SearchTopicFailed
		return
	}

	resp = &model.JoinedTopicsResp{
		Items:  make([]*model.JoinedTopicItem, len(data.Result)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Result {
		t := new(model.ESTopic)
		err = json.Unmarshal(v, t)
		if err != nil {
			return
		}

		item := &model.JoinedTopicItem{
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

		resp.Items[i] = item

	}

	if resp.Paging.Prev, err = genURL("/api/v1/topic/list/followed", url.Values{
		"query": []string{query},
		"pn":    []string{strconv.Itoa(pn - 1)},
		"ps":    []string{strconv.Itoa(ps)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/topic/list/followed", url.Values{
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

func (p *Service) AuditFollow(c context.Context, arg *model.ArgAuditFollow) (err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	item := &topic.ArgAuditFollow{
		Aid:     aid,
		ID:      arg.ID,
		Reason:  arg.Reason,
		Approve: arg.Approve,
	}
	if err = p.d.AuditFollow(c, item); err != nil {
		return
	}

	return
}
