package service

import (
	"context"

	"valerian/app/interface/like/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetDislikesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Dislike, err error)
	GetDislikes(c context.Context, node sqalx.Node) (items []*model.Dislike, err error)
	GetDislikeByID(c context.Context, node sqalx.Node, id int64) (item *model.Dislike, err error)
	GetDislikeByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Dislike, err error)
	AddDislike(c context.Context, node sqalx.Node, item *model.Dislike) (err error)
	UpdateDislike(c context.Context, node sqalx.Node, item *model.Dislike) (err error)
	DelDislike(c context.Context, node sqalx.Node, id int64) (err error)

	GetLikes(c context.Context, node sqalx.Node) (items []*model.Like, err error)
	GetLikesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Like, err error)
	GetLikeByID(c context.Context, node sqalx.Node, id int64) (item *model.Like, err error)
	GetLikeByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Like, err error)
	AddLike(c context.Context, node sqalx.Node, item *model.Like) (err error)
	UpdateLike(c context.Context, node sqalx.Node, item *model.Like) (err error)
	DelLike(c context.Context, node sqalx.Node, id int64) (err error)

	Like(c context.Context, aid, targetID int64, targetType string) (err error)
	CancelLike(c context.Context, aid, targetID int64, targetType string) (err error)
	Dislike(c context.Context, aid, targetID int64, targetType string) (err error)
	CancelDislike(c context.Context, aid, targetID int64, targetType string) (err error)

	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	GetAccountProfile(c context.Context, aid int64) (info *account.ProfileReply, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error)
	GetComment(c context.Context, id int64) (info *comment.CommentInfo, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
