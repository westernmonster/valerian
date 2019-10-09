package service

import (
	"context"

	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
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
