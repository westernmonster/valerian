package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetTopicCatalogChildrenCount(c context.Context, node sqalx.Node, topicID, parentID int64) (count int, err error)
	GetTopicCatalogsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicCatalog, err error)
	GetTopicCatalogs(c context.Context, node sqalx.Node) (items []*model.TopicCatalog, err error)
	GetTopicCatalogByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicCatalog, err error)
	GetTopicCatalogByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicCatalog, err error)
	AddTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)
	UpdateTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)
	DelTopicCatalog(c context.Context, node sqalx.Node, id int64) (err error)

	GetAuthTopicsResp(c context.Context, node sqalx.Node, topicID int64) (items []*model.AuthTopicResp, err error)
	GetAuthTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AuthTopic, err error)
	GetAuthTopics(c context.Context, node sqalx.Node) (items []*model.AuthTopic, err error)
	GetAuthTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.AuthTopic, err error)
	GetAuthTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AuthTopic, err error)
	AddAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error)
	UpdateAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error)
	DelAuthTopic(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Topic, err error)
	GetTopics(c context.Context, node sqalx.Node) (items []*model.Topic, err error)
	GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error)
	GetTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Topic, err error)
	AddTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error)
	UpdateTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error)
	DelTopic(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopicMembersCount(c context.Context, node sqalx.Node, topicID int64) (count int, err error)
	GetTopicMembersPaged(c context.Context, node sqalx.Node, topicID int64, page, pageSize int) (count int, items []*model.TopicMember, err error)
	GetTopicMembersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicMember, err error)
	GetTopicMembers(c context.Context, node sqalx.Node) (items []*model.TopicMember, err error)
	GetTopicMemberByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicMember, err error)
	GetTopicMemberByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicMember, err error)
	AddTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error)
	UpdateTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error)
	DelTopicMember(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopicFollowRequests(c context.Context, node sqalx.Node, topicID int64, status int) (items []*model.TopicFollowRequest, err error)
	GetTopicFollowRequest(c context.Context, node sqalx.Node, topicID, aid int64) (item *model.TopicFollowRequest, err error)
	AddTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error)
	UpdateTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error)
	DelTopicFollowRequest(c context.Context, node sqalx.Node, id int64) (err error)

	GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)
	GetAccountTopicSettings(c context.Context, node sqalx.Node, aid int64) (items []*model.AccountTopicSetting, err error)
	AddAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error)
	UpdateAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error)
	DelAccountTopicSetting(c context.Context, node sqalx.Node, id int64) (err error)
	GetAccountTopicSetting(c context.Context, node sqalx.Node, aid, topicID int64) (item *model.AccountTopicSetting, err error)

	GetDiscussCategoriesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.DiscussCategory, err error)
	GetDiscussCategories(c context.Context, node sqalx.Node) (items []*model.DiscussCategory, err error)
	GetDiscussCategoryByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussCategory, err error)
	GetDiscussCategoryByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.DiscussCategory, err error)
	AddDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error)
	UpdateDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error)
	DelDiscussCategory(c context.Context, node sqalx.Node, id int64) (err error)

	AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
	GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error)
	UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
	DelArticle(c context.Context, node sqalx.Node, id int64) (err error)

	GetArticleFiles(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleFile, err error)
	AddArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
	UpdateArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
	DelArticleFile(c context.Context, node sqalx.Node, id int64) (err error)
	GetArticleFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleFile, err error)

	AddArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error)
	GetArticleHistories(c context.Context, node sqalx.Node, articleVersionID int64) (items []*model.ArticleHistory, err error)
	GetArticleHistoryByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleHistory, err error)
	GetArticleHistoriesMaxSeq(c context.Context, node sqalx.Node, articleVersionID int64) (seq int, err error)

	AddAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error)
	UpdateAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error)
	DelAccountArticleAttr(c context.Context, node sqalx.Node, id int64) (err error)
	GetAccountArticleAttr(c context.Context, node sqalx.Node, aid, articleID int64) (item *model.AccountArticleAttr, err error)
	GetArticleFavCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error)
	GetArticleLikeCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error)

	GetLocaleByCondition(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Locale, err error)

	SetTopicCache(c context.Context, m *model.TopicResp) (err error)
	TopicCache(c context.Context, topicID int64) (m *model.TopicResp, err error)
	DelTopicCache(c context.Context, topicID int64) (err error)

	SetTopicCatalogCache(c context.Context, topicID int64, m []*model.TopicLevel1Catalog) (err error)
	TopicCatalogCache(c context.Context, topicID int64) (m []*model.TopicLevel1Catalog, err error)
	DelTopicCatalogCache(c context.Context, topicID int64) (err error)

	SetTopicMembersCache(c context.Context, topicID int64, count, page, pageSize int, data []*model.TopicMember) (err error)
	TopicMembersCache(c context.Context, topicID int64, page, pageSize int) (count int, data []*model.TopicMember, err error)
	DelTopicMembersCache(c context.Context, topicID int64) (err error)

	SetAuthTopicsCache(c context.Context, topicID int64, m []*model.AuthTopic) (err error)
	AuthTopicsCache(c context.Context, topicID int64) (m []*model.AuthTopic, err error)
	DelAuthTopicsCache(c context.Context, topicID int64) (err error)

	SetArticleCache(c context.Context, m *model.ArticleResp) (err error)
	ArticleCache(c context.Context, articleID int64) (m *model.ArticleResp, err error)
	DelArticleCache(c context.Context, articleID int64) (err error)

	SetArticleFileCache(c context.Context, articleID int64, m []*model.ArticleFileResp) (err error)
	ArticleFileCache(c context.Context, articleID int64) (m []*model.ArticleFileResp, err error)
	DelArticleFileCache(c context.Context, articleID int64) (err error)

	SetArticleHistoryCache(c context.Context, articleVersionID int64, m []*model.ArticleHistoryResp) (err error)
	ArticleHistoryCache(c context.Context, articleVersionID int64) (m []*model.ArticleHistoryResp, err error)
	DelArticleHistoryCache(c context.Context, articleVersionID int64) (err error)

	SetAccountCache(c context.Context, m *model.Account) (err error)
	AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
	DelAccountCache(c context.Context, accountID int64) (err error)

	SetAccountTopicSettingCache(c context.Context, m *model.AccountTopicSetting) (err error)
	AccountTopicSettingCache(c context.Context, aid, topicID int64) (m *model.AccountTopicSetting, err error)
	DelAccountTopicSettingCache(c context.Context, aid, topicID int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
