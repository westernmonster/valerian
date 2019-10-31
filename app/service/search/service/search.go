package service

import (
	"context"
	"valerian/app/service/search/model"
	"valerian/library/ecode"
)

func (p *Service) AccountSearch(c context.Context, arg *model.BasicSearchParams, ids []int64) (resp *model.SearchResult, err error) {
	if resp, err = p.d.AccountSearch(c, arg, ids); err != nil {
		err = ecode.SearchAccountFailed
		return
	}

	return
}

func (p *Service) TopicSearch(c context.Context, arg *model.BasicSearchParams, ids []int64) (resp *model.SearchResult, err error) {
	if resp, err = p.d.TopicSearch(c, arg, ids); err != nil {
		err = ecode.SearchTopicFailed
		return
	}

	return
}

func (p *Service) ArticleSearch(c context.Context, arg *model.BasicSearchParams, ids []int64) (resp *model.SearchResult, err error) {
	if resp, err = p.d.ArticleSearch(c, arg, ids); err != nil {
		err = ecode.SearchArticleFailed
		return
	}

	return
}

func (p *Service) DiscussionSearch(c context.Context, arg *model.BasicSearchParams, ids []int64) (resp *model.SearchResult, err error) {
	if resp, err = p.d.DiscussionSearch(c, arg, ids); err != nil {
		err = ecode.SearchDiscussionFailed
		return
	}

	return
}
