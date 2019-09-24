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

	GetTopicMemberByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicMember, err error)

	SetTopicCache(c context.Context, m *model.Topic) (err error)
	TopicCache(c context.Context, topicID int64) (m *model.Topic, err error)
	DelTopicCache(c context.Context, topicID int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
