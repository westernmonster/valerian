package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"valerian/library/net/metadata"

	"valerian/app/interface/search/model"
	"valerian/library/ecode"
	"valerian/library/xstr"
)

func (p *Service) DiscussionSearch(c context.Context, arg *model.DiscussionSearchParams) (resp *model.DiscussionSearchResult, err error) {
	arg.BasicSearchParams = p.addContextTextSource(arg.BasicSearchParams)
	var data *model.SearchResult
	if data, err = p.d.DiscussionSearch(c, arg); err != nil {
		err = ecode.SearchAccountFailed
	}

	resp = &model.DiscussionSearchResult{
		Paging: &model.Paging{},
		Debug:  data.Debug,
		Data:   make([]*model.ESDiscussion, 0),
	}

	for _, v := range data.Result {
		acc := new(model.ESDiscussion)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		if acc.ContentText != nil {
			excerpt := xstr.Excerpt(*acc.ContentText)
			acc.Excerpt = &excerpt
			acc.ContentText = nil
		}

		var stat *model.DiscussionStat
		if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), acc.ID); err != nil {
			return
		}

		acc.LikeCount = stat.LikeCount
		acc.DislikeCount = stat.DislikeCount
		acc.CommentCount = stat.CommentCount

		resp.Data = append(resp.Data, acc)
	}

	resp.Paging.Total = data.Page.Total

	if resp.Paging.Prev, err = genURL("/api/v1/search/discussions", url.Values{
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

	if resp.Paging.Next, err = genURL("/api/v1/search/discussions", url.Values{
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
	go p.emitSearchStatAdded(context.Background(), arg.KW, "discussion", aid, data.Page.Total)

	return
}
