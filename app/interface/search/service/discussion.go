package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"valerian/app/interface/search/model"
	"valerian/library/ecode"
)

func (p *Service) DiscussionSearch(c context.Context, arg *model.DiscussionSearchParams) (resp *model.DiscussionSearchResult, err error) {
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

		resp.Data = append(resp.Data, acc)
	}

	resp.Paging.Total = data.Page.Total

	if resp.Paging.Prev, err = genURL("/api/v1/search/discussions", url.Values{
		"kw":        []string{arg.Bsp.KW},
		"kw_fields": arg.Bsp.KwFields,
		"order":     arg.Bsp.Order,
		"sort":      arg.Bsp.Sort,
		"debug":     []string{strconv.FormatBool(arg.Bsp.Debug)},
		"source":    arg.Bsp.Source,
		"pn":        []string{strconv.Itoa(arg.Bsp.Pn)},
		"ps":        []string{strconv.Itoa(arg.Bsp.Ps)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/search/discussions", url.Values{
		"kw":        []string{arg.Bsp.KW},
		"kw_fields": arg.Bsp.KwFields,
		"order":     arg.Bsp.Order,
		"sort":      arg.Bsp.Sort,
		"debug":     []string{strconv.FormatBool(arg.Bsp.Debug)},
		"source":    arg.Bsp.Source,
		"pn":        []string{strconv.Itoa(arg.Bsp.Pn + 1)},
		"ps":        []string{strconv.Itoa(arg.Bsp.Ps)},
	}); err != nil {
		return
	}

	if len(resp.Data) < arg.Bsp.Ps {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if arg.Bsp.Pn == 1 {
		resp.Paging.Prev = ""
	}

	return
}
