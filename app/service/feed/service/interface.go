package service

import (
	"context"

	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/model"
	relation "valerian/app/service/relation/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetFeedPaged(c context.Context, node sqalx.Node, accountID int64, limit, offset int) (items []*model.Feed, err error)
	AddFeed(c context.Context, node sqalx.Node, item *model.Feed) (err error)
	UpdateFeed(c context.Context, node sqalx.Node, item *model.Feed) (err error)
	DelFeedByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error)

	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)
	GetFansIDs(c context.Context, aid int64) (resp *relation.IDsResp, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	GetBelongsTopicIDs(c context.Context, aid int64) (resp *topic.IDsResp, err error)
	GetTopicMemberIDs(c context.Context, topicID int64) (resp *topic.IDsResp, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
