package service

import (
	"context"

	"valerian/app/service/account-feed/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetAccountFeedPaged(c context.Context, node sqalx.Node, topicID int64, limit, offset int) (items []*model.AccountFeed, err error)
	AddAccountFeed(c context.Context, node sqalx.Node, item *model.AccountFeed) (err error)
	UpdateAccountFeed(c context.Context, node sqalx.Node, item *model.AccountFeed) (err error)
	DelAccountFeed(c context.Context, node sqalx.Node, id int64) (err error)
	DelAccountFeedByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error)

	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)

	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
