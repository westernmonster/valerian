package service

import (
	"context"
	"valerian/app/interface/article/model"
	account "valerian/app/service/account/api"
	topic "valerian/app/service/topic/api"

	"valerian/library/database/sqalx"
)

type IDao interface {
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)

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

	GetArticleHistoriesPaged(c context.Context, node sqalx.Node, articleID int64, limit, offset int) (items []*model.ArticleHistory, err error)
	GetArticleHistoriesMaxSeq(c context.Context, node sqalx.Node, articleID int64) (seq int, err error)
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

	GetTopicCatalogChildrenCount(c context.Context, node sqalx.Node, topicID, parentID int64) (count int, err error)
	GetTopicCatalogsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicCatalog, err error)
	GetTopicCatalogs(c context.Context, node sqalx.Node) (items []*model.TopicCatalog, err error)
	GetTopicCatalogByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicCatalog, err error)
	GetTopicCatalogByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicCatalog, err error)
	AddTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)
	UpdateTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)
	DelTopicCatalog(c context.Context, node sqalx.Node, id int64) (err error)
	GetTopicCatalogMaxChildrenSeq(c context.Context, node sqalx.Node, topicID, parentID int64) (seq int, err error)

	GetRevisesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Revise, err error)
	GetRevises(c context.Context, node sqalx.Node) (items []*model.Revise, err error)
	GetReviseByID(c context.Context, node sqalx.Node, id int64) (item *model.Revise, err error)
	GetReviseByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Revise, err error)
	AddRevise(c context.Context, node sqalx.Node, item *model.Revise) (err error)
	UpdateRevise(c context.Context, node sqalx.Node, item *model.Revise) (err error)
	DelRevise(c context.Context, node sqalx.Node, id int64) (err error)

	GetReviseFilesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ReviseFile, err error)
	GetReviseFiles(c context.Context, node sqalx.Node) (items []*model.ReviseFile, err error)
	GetReviseFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ReviseFile, err error)
	GetReviseFileByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ReviseFile, err error)
	AddReviseFile(c context.Context, node sqalx.Node, item *model.ReviseFile) (err error)
	UpdateReviseFile(c context.Context, node sqalx.Node, item *model.ReviseFile) (err error)
	DelReviseFile(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)

	SetTopicCatalogCache(c context.Context, topicID int64, m []*model.TopicLevel1Catalog) (err error)
	TopicCatalogCache(c context.Context, topicID int64) (m []*model.TopicLevel1Catalog, err error)
	DelTopicCatalogCache(c context.Context, topicID int64) (err error)

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
