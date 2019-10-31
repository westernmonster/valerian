package service

import (
	"context"

	account "valerian/app/service/account/api"
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
	DelTopicFeedByTopicID(c context.Context, node sqalx.Node, topicID int64) (err error)
	DelTopicFeedByCond(c context.Context, node sqalx.Node, topicID int64, targetType string, targetID int64) (err error)

	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
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

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
