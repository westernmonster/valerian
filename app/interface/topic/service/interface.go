package service

import (
	"context"
	"valerian/app/interface/topic/model"
	account "valerian/app/service/account/api"
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

	GetTopicInviteRequestsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicInviteRequest, err error)
	GetTopicInviteRequests(c context.Context, node sqalx.Node) (items []*model.TopicInviteRequest, err error)
	GetTopicInviteRequestByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicInviteRequest, err error)
	GetTopicInviteRequestByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicInviteRequest, err error)
	AddTopicInviteRequest(c context.Context, node sqalx.Node, item *model.TopicInviteRequest) (err error)
	UpdateTopicInviteRequest(c context.Context, node sqalx.Node, item *model.TopicInviteRequest) (err error)
	DelTopicInviteRequest(c context.Context, node sqalx.Node, id int64) (err error)

	AddTopicMemberStat(c context.Context, node sqalx.Node, item *model.TopicMemberStat) (err error)
	GetTopicMemberStatForUpdate(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicMemberStat, err error)
	UpdateTopicMemberStat(c context.Context, node sqalx.Node, item *model.TopicMemberStat) (err error)
	GetTopicMemberStat(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicMemberStat, err error)

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