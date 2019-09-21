package service

import (
	"context"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
	GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error)
	UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
	DelArticle(c context.Context, node sqalx.Node, id int64) (err error)

	GetArticleFilesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ArticleFile, err error)
	GetArticleFiles(c context.Context, node sqalx.Node) (items []*model.ArticleFile, err error)
	GetArticleFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleFile, err error)
	GetArticleFileByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ArticleFile, err error)
	AddArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
	UpdateArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
	DelArticleFile(c context.Context, node sqalx.Node, id int64) (err error)

	GetArticleHistoriesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ArticleHistory, err error)
	GetArticleHistories(c context.Context, node sqalx.Node) (items []*model.ArticleHistory, err error)
	GetArticleHistoryByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleHistory, err error)
	GetArticleHistoryByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ArticleHistory, err error)
	AddArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error)
	UpdateArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error)
	DelArticleHistory(c context.Context, node sqalx.Node, id int64) (err error)

	AddAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error)
	UpdateAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error)
	DelAccountArticleAttr(c context.Context, node sqalx.Node, id int64) (err error)
	GetAccountArticleAttr(c context.Context, node sqalx.Node, aid, articleID int64) (item *model.AccountArticleAttr, err error)
	GetArticleFavCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error)
	GetArticleLikeCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error)

	SetArticleCache(c context.Context, m *model.ArticleResp) (err error)
	ArticleCache(c context.Context, articleID int64) (m *model.ArticleResp, err error)
	DelArticleCache(c context.Context, articleID int64) (err error)

	SetArticleFileCache(c context.Context, articleID int64, m []*model.ArticleFileResp) (err error)
	ArticleFileCache(c context.Context, articleID int64) (m []*model.ArticleFileResp, err error)
	DelArticleFileCache(c context.Context, articleID int64) (err error)

	SetArticleHistoryCache(c context.Context, articleVersionID int64, m []*model.ArticleHistoryResp) (err error)
	ArticleHistoryCache(c context.Context, articleVersionID int64) (m []*model.ArticleHistoryResp, err error)
	DelArticleHistoryCache(c context.Context, articleVersionID int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
