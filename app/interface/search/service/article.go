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

func (p *Service) ArticleSuggest(c context.Context, kw string) (resp []string, err error) {
	resp = make([]string, 0)

	var articleSuggestions []string
	if articleSuggestions, err = p.d.ArticleSuggest(c, kw, 10); err != nil {
		return
	}
	resp = append(resp, articleSuggestions...)

	return
}

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

	var data *model.SearchResult
	if data, err = p.d.ArticleSearch(c, arg); err != nil {
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

	if resp.Paging.Prev, err = genURL("/api/v1/search/articles", url.Values{
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

	if resp.Paging.Next, err = genURL("/api/v1/search/articles", url.Values{
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
	p.addCache(func() {
		p.emitSearchStatAdded(context.Background(), arg.KW, "article", aid, data.Page.Total)
	})

	return
}
