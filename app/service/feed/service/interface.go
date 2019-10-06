package service

import (
	"context"

	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetTopicFeedPaged(c context.Context, node sqalx.Node, topicID int64, limit, offset int) (items []*model.TopicFeed, err error)
	AddTopicFeed(c context.Context, node sqalx.Node, item *model.TopicFeed) (err error)
	UpdateTopicFeed(c context.Context, node sqalx.Node, item *model.TopicFeed) (err error)
	DelTopicFeed(c context.Context, node sqalx.Node, id int64) (err error)
	DelTopicFeedByCond(c context.Context, node sqalx.Node, topicID int64, targetType string, targetID int64) (err error)

	GetAccountFeedPaged(c context.Context, node sqalx.Node, accountID int64, limit, offset int) (items []*model.AccountFeed, err error)
	AddAccountFeed(c context.Context, node sqalx.Node, item *model.AccountFeed) (err error)
	UpdateAccountFeed(c context.Context, node sqalx.Node, item *model.AccountFeed) (err error)
	DelAccountFeedByCond(c context.Context, node sqalx.Node, accountID int64, targetType string, targetID int64) (err error)

	GetFeedPaged(c context.Context, node sqalx.Node, accountID int64, limit, offset int) (items []*model.Feed, err error)
	AddFeed(c context.Context, node sqalx.Node, item *model.Feed) (err error)
	UpdateFeed(c context.Context, node sqalx.Node, item *model.Feed) (err error)
	DelFeedByCond(c context.Context, node sqalx.Node, accountID int64, targetType string, targetID int64) (err error)

	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
