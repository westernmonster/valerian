package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"valerian/library/net/metadata"

	"valerian/app/interface/search/model"
	"valerian/library/ecode"
)

func (p *Service) TopicSearch(c context.Context, arg *model.TopicSearchParams) (resp *model.TopicSearchResult, err error) {
	var data *model.SearchResult
	if data, err = p.d.TopicSearch(c, arg); err != nil {
		err = ecode.SearchTopicFailed
		return
	}

	resp = &model.TopicSearchResult{
		Paging: &model.Paging{},
		Debug:  data.Debug,
		Data:   make([]*model.ESTopic, 0),
	}

	for _, v := range data.Result {
		acc := new(model.ESTopic)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		var stat *model.TopicStat
		if stat, err = p.d.GetTopicStatByID(c, p.d.DB(), acc.ID); err != nil {
			return
		}

		acc.MemberCount = stat.MemberCount
		acc.ArticleCount = stat.ArticleCount
		acc.DiscussionCount = stat.DiscussionCount

		resp.Data = append(resp.Data, acc)
	}

	resp.Paging.Total = data.Page.Total

	if resp.Paging.Prev, err = genURL("/api/v1/search/topics", url.Values{
		"kw":        []string{arg.KW},
		"kw_fields": arg.KwFields,
		"order":     arg.Order,
		"sort":      arg.Sort,
		"debug":     []string{strconv.FormatBool(arg.Debug)},
		"source":    arg.Source,
		"pn":        []string{strconv.Itoa(arg.Pn)},
		"ps":        []string{strconv.Itoa(arg.Ps)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/search/topics", url.Values{
		"kw":        []string{arg.KW},
		"kw_fields": arg.KwFields,
		"order":     arg.Order,
		"sort":      arg.Sort,
		"debug":     []string{strconv.FormatBool(arg.Debug)},
		"source":    arg.Source,
		"pn":        []string{strconv.Itoa(arg.Pn + 1)},
		"ps":        []string{strconv.Itoa(arg.Ps)},
	}); err != nil {
		return
	}

	if len(resp.Data) < arg.Ps {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if arg.Pn == 1 {
		resp.Paging.Prev = ""
	}

	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		//err = ecode.AcquireAccountIDFailed
		//return
	}
	p.emitSearchStatAdded(context.Background(), arg.KW, "topic", aid, data.Page.Total)

	return
}
