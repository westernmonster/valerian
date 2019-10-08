package service

import (
	"context"

	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/topic-feed/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetTopicFeedPaged(c context.Context, node sqalx.Node, topicID int64, limit, offset int) (items []*model.TopicFeed, err error)
	AddTopicFeed(c context.Context, node sqalx.Node, item *model.TopicFeed) (err error)
	UpdateTopicFeed(c context.Context, node sqalx.Node, item *model.TopicFeed) (err error)
	DelTopicFeed(c context.Context, node sqalx.Node, id int64) (err error)
	DelTopicFeedByCond(c context.Context, node sqalx.Node, topicID int64, targetType string, targetID int64) (err error)

	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
