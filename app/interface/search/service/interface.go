package service

import (
	"context"
	"valerian/app/interface/search/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	AccountSearch(c context.Context, arg *model.AccountSearchParams) (res *model.SearchResult, err error)
	TopicSearch(c context.Context, arg *model.TopicSearchParams) (res *model.SearchResult, err error)
	ArticleSearch(c context.Context, arg *model.ArticleSearchParams) (res *model.SearchResult, err error)
	DiscussionSearch(c context.Context, arg *model.DiscussionSearchParams) (res *model.SearchResult, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
