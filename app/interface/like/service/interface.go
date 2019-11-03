package service

import (
	"context"

	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	Like(c context.Context, aid, targetID int64, targetType string) (err error)
	CancelLike(c context.Context, aid, targetID int64, targetType string) (err error)
	Dislike(c context.Context, aid, targetID int64, targetType string) (err error)
	CancelDislike(c context.Context, aid, targetID int64, targetType string) (err error)

	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	GetMemberInfo(c context.Context, aid int64) (info *account.MemberInfoReply, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error)
	GetComment(c context.Context, id int64) (info *comment.CommentInfo, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
