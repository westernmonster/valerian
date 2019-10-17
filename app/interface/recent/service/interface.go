package service

import (
	"context"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	recent "valerian/app/service/recent/api"
	topic "valerian/app/service/topic/api"

	"valerian/library/database/sqalx"
)

type IDao interface {
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)

	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetRecentViewsPaged(c context.Context, aid int64, targetType string, limit, offset int) (resp *recent.RecentViewsResp, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
