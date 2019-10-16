package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"valerian/app/interface/search/model"
	"valerian/library/ecode"
)

func (p *Service) TopicSearch(c context.Context, arg *model.TopicSearchParams) (resp *model.TopicSearchResult, err error) {
	var data *model.SearchResult
	if data, err = p.d.TopicSearch(c, arg); err != nil {
		err = ecode.SearchAccountFailed
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

	return
}
