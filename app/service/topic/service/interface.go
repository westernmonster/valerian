package service

import (
	"context"
	account "valerian/app/service/account/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)

	GetTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Topic, err error)
	GetTopics(c context.Context, node sqalx.Node) (items []*model.Topic, err error)
	GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error)
	GetTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Topic, err error)
	AddTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error)
	UpdateTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error)
	DelTopic(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopicInviteRequestsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicInviteRequest, err error)
	GetTopicInviteRequests(c context.Context, node sqalx.Node) (items []*model.TopicInviteRequest, err error)
	GetTopicInviteRequestByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicInviteRequest, err error)
	GetTopicInviteRequestByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicInviteRequest, err error)
	AddTopicInviteRequest(c context.Context, node sqalx.Node, item *model.TopicInviteRequest) (err error)
	UpdateTopicInviteRequest(c context.Context, node sqalx.Node, item *model.TopicInviteRequest) (err error)
	DelTopicInviteRequest(c context.Context, node sqalx.Node, id int64) (err error)

	GetAccountTopicSettingsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountTopicSetting, err error)
	GetAccountTopicSettings(c context.Context, node sqalx.Node) (items []*model.AccountTopicSetting, err error)
	GetAccountTopicSettingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountTopicSetting, err error)
	GetAccountTopicSettingByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountTopicSetting, err error)
	AddAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error)
	UpdateAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error)
	DelAccountTopicSetting(c context.Context, node sqalx.Node, id int64) (err error)

	GetTopicMemberStat(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicMemberStat, err error)

	GetTopicMemberByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicMember, err error)

	IsFav(c context.Context, aid, targetID int64, targetType string) (isFav bool, err error)

	SetAccountTopicSettingCache(c context.Context, m *model.AccountTopicSetting) (err error)
	AccountTopicSettingCache(c context.Context, aid, topicID int64) (m *model.AccountTopicSetting, err error)
	DelAccountTopicSettingCache(c context.Context, aid, topicID int64) (err error)

	GetAuthTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AuthTopic, err error)
	GetAuthTopics(c context.Context, node sqalx.Node) (items []*model.AuthTopic, err error)
	GetAuthTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.AuthTopic, err error)
	GetAuthTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AuthTopic, err error)
	AddAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error)
	UpdateAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error)
	DelAuthTopic(c context.Context, node sqalx.Node, id int64) (err error)

	SetTopicCache(c context.Context, m *model.Topic) (err error)
	TopicCache(c context.Context, topicID int64) (m *model.Topic, err error)
	DelTopicCache(c context.Context, topicID int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
