package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"valerian/app/admin/search/model"
	search "valerian/app/service/search/api"
	"valerian/library/ecode"
	"valerian/library/xstr"
)

func (p *Service) addContextTextSource(arg *model.BasicSearchParams) (resp *model.BasicSearchParams) {
	if arg.Source == nil {
		return arg
	}

	hasContentText := false
	for _, v := range arg.Source {
		if v == "content_text" {
			hasContentText = true
		}
	}

	if !hasContentText {
		arg.Source = append(arg.Source, "content_text")
	}
	return arg
}

func (p *Service) ArticleSearch(c context.Context, arg *model.ArticleSearchParams) (resp *model.ArticleSearchResult, err error) {
	arg.BasicSearchParams = p.addContextTextSource(arg.BasicSearchParams)

	var data *search.SearchResult
	if data, err = p.d.SearchArticle(c, arg.BasicSearchParams); err != nil {
		err = ecode.SearchAccountFailed
		return
	}

	resp = &model.ArticleSearchResult{
		Paging: &model.Paging{},
		Debug:  data.Debug,
		Data:   make([]*model.ESArticle, 0),
	}

	for _, v := range data.Result {
		acc := new(model.ESArticle)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		if acc.ContentText != nil {
			excerpt := xstr.Excerpt(*acc.ContentText)
			acc.Excerpt = &excerpt
			acc.ContentText = nil
		}

		var stat *model.ArticleStat
		if stat, err = p.d.GetArticleStatByID(c, p.d.DB(), acc.ID); err != nil {
			return
		}

		acc.LikeCount = stat.LikeCount
		acc.DislikeCount = stat.DislikeCount
		acc.ReviseCount = stat.ReviseCount
		acc.CommentCount = stat.CommentCount

		resp.Data = append(resp.Data, acc)
	}

	resp.Paging.Total = data.Page.Total

	if resp.Paging.Prev, err = genURL("/api/v1/admin/search/articles", url.Values{
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

	if resp.Paging.Next, err = genURL("/api/v1/admin/search/articles", url.Values{
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

func (p *Service) DiscussionSearch(c context.Context, arg *model.DiscussionSearchParams) (resp *model.DiscussionSearchResult, err error) {
	arg.BasicSearchParams = p.addContextTextSource(arg.BasicSearchParams)
	var data *search.SearchResult
	if data, err = p.d.SearchDiscussion(c, arg.BasicSearchParams); err != nil {
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

	if resp.Paging.Prev, err = genURL("/api/v1/admin/search/discussions", url.Values{
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

	if resp.Paging.Next, err = genURL("/api/v1/admin/search/discussions", url.Values{
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

func (p *Service) TopicSearch(c context.Context, arg *model.TopicSearchParams) (resp *model.TopicSearchResult, err error) {
	var data *search.SearchResult
	if data, err = p.d.SearchTopic(c, arg.BasicSearchParams); err != nil {
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

	if resp.Paging.Prev, err = genURL("/api/v1/admin/search/topics", url.Values{
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

	if resp.Paging.Next, err = genURL("/api/v1/admin/search/topics", url.Values{
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

func (p *Service) AccountSearch(c context.Context, arg *model.AccountSearchParams) (resp *model.AccountSearchResult, err error) {
	var data *search.SearchResult
	if data, err = p.d.SearchAccount(c, arg.BasicSearchParams); err != nil {
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

		var stat *model.AccountStat
		if stat, err = p.d.GetAccountStatByID(c, p.d.DB(), acc.ID); err != nil {
			return
		}

		acc.FansCount = stat.Fans
		acc.FollowingCount = stat.Following

		resp.Data = append(resp.Data, acc)
	}

	resp.Paging.Total = data.Page.Total

	if resp.Paging.Prev, err = genURL("/api/v1/admin/search/accounts", url.Values{
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

	if resp.Paging.Next, err = genURL("/api/v1/admin/search/accounts", url.Values{
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
