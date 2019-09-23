package service

import (
	"context"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	AddTopicFeed(c context.Context, node sqalx.Node, item *model.TopicFeed) (err error)
	UpdateTopicFeed(c context.Context, node sqalx.Node, item *model.TopicFeed) (err error)
	DelTopicFeed(c context.Context, node sqalx.Node, id int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
