package service

import (
	"context"
	"encoding/json"
	"valerian/app/interface/search/model"
	"valerian/library/ecode"
	"valerian/library/xstr"
)

func (p *Service) AllSearch(c context.Context, kw string) (resp *model.AllSearchResult, err error) {
	resp = &model.AllSearchResult{
		Topics:      make([]*model.ESTopic, 0),
		Accounts:    make([]*model.ESAccount, 0),
		Articles:    make([]*model.ESArticle, 0),
		Discussions: make([]*model.ESDiscussion, 0),
	}

	var accData *model.SearchResult
	if accData, err = p.d.AccountSearch(c, &model.AccountSearchParams{&model.BasicSearchParams{Pn: 1, Ps: 2}}); err != nil {
		err = ecode.SearchAccountFailed
		return
	}

	accounts := make([]*model.ESAccount, 0)
	for _, v := range accData.Result {
		acc := new(model.ESAccount)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		accounts = append(accounts, acc)
	}

	resp.Accounts = accounts
	resp.AccountsCount = accData.Page.Total

	var topicData *model.SearchResult
	if topicData, err = p.d.TopicSearch(c, &model.TopicSearchParams{&model.BasicSearchParams{Pn: 1, Ps: 2}}); err != nil {
		err = ecode.SearchTopicFailed
		return
	}

	topics := make([]*model.ESTopic, 0)
	for _, v := range topicData.Result {
		acc := new(model.ESTopic)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		topics = append(topics, acc)
	}

	resp.Topics = topics
	resp.TopicsCount = topicData.Page.Total

	var articleData *model.SearchResult
	if articleData, err = p.d.ArticleSearch(c, &model.ArticleSearchParams{&model.BasicSearchParams{Pn: 1, Ps: 2}}); err != nil {
		err = ecode.SearchArticleFailed
		return
	}

	articles := make([]*model.ESArticle, 0)
	for _, v := range articleData.Result {
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

		articles = append(articles, acc)
	}

	resp.Articles = articles
	resp.ArticlesCount = articleData.Page.Total

	var discussionData *model.SearchResult
	if discussionData, err = p.d.DiscussionSearch(c, &model.DiscussionSearchParams{&model.BasicSearchParams{Pn: 1, Ps: 2}}); err != nil {
		err = ecode.SearchDiscussionFailed
		return
	}

	discussions := make([]*model.ESDiscussion, 0)
	for _, v := range discussionData.Result {
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

		discussions = append(discussions, acc)
	}

	resp.Discussions = discussions
	resp.DiscussionsCount = discussionData.Page.Total

	return
}
