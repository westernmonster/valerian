package service

import (
	"context"
	"valerian/app/interface/topic/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	"valerian/app/service/discuss/api"
	discuss "valerian/app/service/discuss/api"
	relation "valerian/app/service/relation/api"
	topicFeed "valerian/app/service/topic-feed/api"
	stopic "valerian/app/service/topic/api"

	"valerian/library/database/sqalx"
)

type IDao interface {
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)

	GetTopicCatalogChildrenCount(c context.Context, node sqalx.Node, topicID, parentID int64) (count int, err error)
	GetTopicCatalogsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicCatalog, err error)
	GetTopicCatalogs(c context.Context, node sqalx.Node) (items []*model.TopicCatalog, err error)
	GetTopicCatalogByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicCatalog, err error)
	GetTopicCatalogByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicCatalog, err error)
	AddTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)
	UpdateTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)
	DelTopicCatalog(c context.Context, node sqalx.Node, id int64) (err error)
	GetTopicCatalogMaxChildrenSeq(c context.Context, node sqalx.Node, topicID, parentID int64) (seq int, err error)
	HasTaxonomy(c context.Context, node sqalx.Node, topicID int64) (has bool, err error)

	GetAuthTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AuthTopic, err error)
	GetAuthTopics(c context.Context, node sqalx.Node) (items []*model.AuthTopic, err error)
	GetAuthTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.AuthTopic, err error)
	GetAuthTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AuthTopic, err error)
	AddAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error)
	UpdateAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error)
	DelAuthTopic(c context.Context, node sqalx.Node, id int64) (err error)
	GetUserCanEditTopicIDs(c context.Context, node sqalx.Node, aid int64) (items []int64, err error)

	GetTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Topic, err error)
	GetTopics(c context.Context, node sqalx.Node) (items []*model.Topic, err error)
	GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error)
	GetTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Topic, err error)
	AddTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error)
	UpdateTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error)
	DelTopic(c context.Context, node sqalx.Node, id int64) (err error)

	GetFollowedTopicsPaged(c context.Context, node sqalx.Node, aid int64, query string, limit, offset int) (items []*model.Topic, err error)
	GetFollowedTopicsIDs(c context.Context, node sqalx.Node, aid int64) (items []int64, err error)
	GetTopicMembersCount(c context.Context, node sqalx.Node, topicID int64) (count int, err error)
	GetTopicMembersPaged(c context.Context, node sqalx.Node, topicID int64, page, pageSize int) (count int, items []*model.TopicMember, err error)
	GetTopicMembersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicMember, err error)
	GetTopicMembers(c context.Context, node sqalx.Node) (items []*model.TopicMember, err error)
	GetTopicMemberByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicMember, err error)
	GetTopicMemberByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicMember, err error)
	AddTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error)
	UpdateTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error)
	DelTopicMember(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopicInviteRequestsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicInviteRequest, err error)
	GetTopicInviteRequests(c context.Context, node sqalx.Node) (items []*model.TopicInviteRequest, err error)
	GetTopicInviteRequestByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicInviteRequest, err error)
	GetTopicInviteRequestByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicInviteRequest, err error)
	AddTopicInviteRequest(c context.Context, node sqalx.Node, item *model.TopicInviteRequest) (err error)
	UpdateTopicInviteRequest(c context.Context, node sqalx.Node, item *model.TopicInviteRequest) (err error)
	DelTopicInviteRequest(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopicFollowRequestsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicFollowRequest, err error)
	GetTopicFollowRequests(c context.Context, node sqalx.Node) (items []*model.TopicFollowRequest, err error)
	GetTopicFollowRequestByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicFollowRequest, err error)
	GetTopicFollowRequestByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicFollowRequest, err error)
	AddTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error)
	UpdateTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error)
	DelTopicFollowRequest(c context.Context, node sqalx.Node, id int64) (err error)

	GetAccountTopicSettingsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountTopicSetting, err error)
	GetAccountTopicSettings(c context.Context, node sqalx.Node) (items []*model.AccountTopicSetting, err error)
	GetAccountTopicSettingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountTopicSetting, err error)
	GetAccountTopicSettingByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountTopicSetting, err error)
	AddAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error)
	UpdateAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error)
	DelAccountTopicSetting(c context.Context, node sqalx.Node, id int64) (err error)

	GetDiscussCategoriesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.DiscussCategory, err error)
	GetDiscussCategories(c context.Context, node sqalx.Node) (items []*model.DiscussCategory, err error)
	GetDiscussCategoryByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussCategory, err error)
	GetDiscussCategoryByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.DiscussCategory, err error)
	AddDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error)
	UpdateDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error)
	DelDiscussCategory(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopicStatForUpdate(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicStat, err error)
	AddTopicStat(c context.Context, node sqalx.Node, item *model.TopicStat) (err error)
	UpdateTopicStat(c context.Context, node sqalx.Node, item *model.TopicStat) (err error)
	IncrTopicStat(c context.Context, node sqalx.Node, item *model.TopicStat) (err error)
	GetTopicStatByID(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicStat, err error)

	IncrAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error)
	GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountStat, err error)

	TopicSearch(c context.Context, arg *model.TopicSearchParams, ids []int64) (res *model.SearchResult, err error)
	AccountSearch(c context.Context, arg *model.AccountSearchParams, ids []int64) (res *model.SearchResult, err error)

	GetFans(c context.Context, accountID int64, limit, offset int) (resp *relation.FansResp, err error)
	GetFansIDs(c context.Context, aid int64) (resp *relation.IDsResp, err error)

	GetTopicMeta(c context.Context, aid, topicID int64) (info *stopic.TopicMetaInfo, err error)
	GetTopicFeedPaged(c context.Context, topicID int64, limit, offset int) (info *topicFeed.TopicFeedResp, err error)
	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetDiscussionCategories(c context.Context, topicID int64) (resp *api.CategoriesResp, err error)

	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)

	SetTopicCache(c context.Context, m *model.Topic) (err error)
	TopicCache(c context.Context, topicID int64) (m *model.Topic, err error)
	DelTopicCache(c context.Context, topicID int64) (err error)

	SetTopicCatalogCache(c context.Context, topicID int64, m []*model.TopicLevel1Catalog) (err error)
	TopicCatalogCache(c context.Context, topicID int64) (m []*model.TopicLevel1Catalog, err error)
	DelTopicCatalogCache(c context.Context, topicID int64) (err error)

	SetAuthTopicsCache(c context.Context, topicID int64, m []*model.AuthTopic) (err error)
	AuthTopicsCache(c context.Context, topicID int64) (m []*model.AuthTopic, err error)
	DelAuthTopicsCache(c context.Context, topicID int64) (err error)

	SetAccountTopicSettingCache(c context.Context, m *model.AccountTopicSetting) (err error)
	AccountTopicSettingCache(c context.Context, aid, topicID int64) (m *model.AccountTopicSetting, err error)
	DelAccountTopicSettingCache(c context.Context, aid, topicID int64) (err error)

	SetTopicMembersCache(c context.Context, topicID int64, count, page, pageSize int, data []*model.TopicMember) (err error)
	TopicMembersCache(c context.Context, topicID int64, page, pageSize int) (count int, data []*model.TopicMember, err error)
	DelTopicMembersCache(c context.Context, topicID int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
