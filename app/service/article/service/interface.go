package service

import (
	"context"

	account "valerian/app/service/account/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetUserArticlesPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.Article, err error)
	AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
	GetArticles(c context.Context, node sqalx.Node) (items []*model.Article, err error)
	GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error)
	UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
	DelArticle(c context.Context, node sqalx.Node, id int64) (err error)

	GetLastArticleHistory(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleHistory, err error)

	GetRevisesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Revise, err error)
	GetRevises(c context.Context, node sqalx.Node) (items []*model.Revise, err error)
	GetReviseByID(c context.Context, node sqalx.Node, id int64) (item *model.Revise, err error)
	GetReviseByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Revise, err error)
	AddRevise(c context.Context, node sqalx.Node, item *model.Revise) (err error)
	UpdateRevise(c context.Context, node sqalx.Node, item *model.Revise) (err error)
	DelRevise(c context.Context, node sqalx.Node, id int64) (err error)

	GetImageUrlsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ImageURL, err error)
	GetImageUrls(c context.Context, node sqalx.Node) (items []*model.ImageURL, err error)
	GetImageURLByID(c context.Context, node sqalx.Node, id int64) (item *model.ImageURL, err error)
	GetImageURLByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ImageURL, err error)
	AddImageURL(c context.Context, node sqalx.Node, item *model.ImageURL) (err error)
	UpdateImageURL(c context.Context, node sqalx.Node, item *model.ImageURL) (err error)
	DelImageURL(c context.Context, node sqalx.Node, id int64) (err error)

	GetArticleStatByID(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleStat, err error)

	GetReviseStatByID(c context.Context, node sqalx.Node, reviseID int64) (item *model.ReviseStat, err error)

	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)

	SetArticleCache(c context.Context, m *model.Article) (err error)
	ArticleCache(c context.Context, articleID int64) (m *model.Article, err error)
	DelArticleCache(c context.Context, articleID int64) (err error)

	SetReviseCache(c context.Context, m *model.Revise) (err error)
	ReviseCache(c context.Context, reviseID int64) (m *model.Revise, err error)
	DelReviseCache(c context.Context, reviseID int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
