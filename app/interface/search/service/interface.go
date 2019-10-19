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

	GetDiscussionStatByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussionStat, err error)
	GetArticleStatByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleStat, err error)
	GetTopicStatByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicStat, err error)
	GetAccountStatByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountStat, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
