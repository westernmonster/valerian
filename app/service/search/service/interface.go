package service

import (
	"context"
	"valerian/app/service/search/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	CreateAccountIndices(c context.Context) (err error)
	CreateArticleIndices(c context.Context) (err error)
	CreateDiscussionIndices(c context.Context) (err error)
	CreateTopicIndices(c context.Context) (err error)

	PutAccount2ES(c context.Context, item *model.ESAccount) (err error)
	PutArticle2ES(c context.Context, item *model.ESArticle) (err error)
	PutDiscussion2ES(c context.Context, item *model.ESDiscussion) (err error)
	PutTopic2ES(c context.Context, item *model.ESTopic) (err error)

	GetAccounts(c context.Context, node sqalx.Node) (items []*model.Account, err error)
	GetArticles(c context.Context, node sqalx.Node) (items []*model.Article, err error)
	GetTopics(c context.Context, node sqalx.Node) (items []*model.Topic, err error)
	GetDiscussions(c context.Context, node sqalx.Node) (items []*model.Discussion, err error)
	GetDiscussCategoryByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussCategory, err error)
	GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)
	GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
