package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	account "valerian/app/service/account/api"
	"valerian/app/service/search/model"
	"valerian/library/ecode"
)

func (p *Service) AccountSearch(c context.Context, arg *model.AccountSearchParams) (resp *model.AccountSearchResult, err error) {
	var data *model.SearchResult
	if data, err = p.d.AccountSearch(c, arg); err != nil {
		err = ecode.SearchAccountFailed
		return
	}

	resp = &model.AccountSearchResult{
		Paging: &model.Paging{},
		Debug:  data.Debug,
		Data:   make([]*model.ESAccount, 0),
	}

	for _, v := range data.Result {
		acc := new(model.ESAccount)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		var stat *account.AccountStatInfo
		if stat, err = p.d.GetAccountStat(c, acc.ID); err != nil {
			return
		}

		acc.FansCount = (stat.FansCount)
		acc.FollowingCount = (stat.FollowingCount)

		resp.Data = append(resp.Data, acc)
	}

	resp.Paging.Total = data.Page.Total

	if resp.Paging.Prev, err = genURL("/api/v1/search/accounts", url.Values{
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

	if resp.Paging.Next, err = genURL("/api/v1/search/accounts", url.Values{
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
