package service

import (
	"context"

	"valerian/app/interface/comment/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetCommentsPaged(c context.Context, node sqalx.Node, resourceID int64, targetType string, limit, offset int) (items []*model.Comment, err error)
	GetCommentByID(c context.Context, node sqalx.Node, id int64) (item *model.Comment, err error)
	AddComment(c context.Context, node sqalx.Node, item *model.Comment) (err error)
	UpdateComment(c context.Context, node sqalx.Node, item *model.Comment) (err error)
	DelComment(c context.Context, node sqalx.Node, id int64) (err error)
	GetAllChildrenComments(c context.Context, node sqalx.Node, resourceID int64) (items []*model.Comment, err error)
	GetChildrenComments(c context.Context, node sqalx.Node, resourceID int64, limit int) (items []*model.Comment, err error)

	GetCommentStatByID(c context.Context, node sqalx.Node, commentID int64) (item *model.CommentStat, err error)
	AddCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error)
	UpdateCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error)
	IncrCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error)

	GetArticleStatByID(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleStat, err error)
	GetReviseStatByID(c context.Context, node sqalx.Node, reviseID int64) (item *model.ReviseStat, err error)
	GetDiscussionStatByID(c context.Context, node sqalx.Node, discussionID int64) (item *model.DiscussionStat, err error)

	IncrReviseStat(c context.Context, node sqalx.Node, item *model.ReviseStat) (err error)
	IncrArticleStat(c context.Context, node sqalx.Node, item *model.ArticleStat) (err error)
	IncrDiscussionStat(c context.Context, node sqalx.Node, item *model.DiscussionStat) (err error)

	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	IsLike(c context.Context, aid, targetID int64, targetType string) (isLike bool, err error)
	IsDislike(c context.Context, aid, targetID int64, targetType string) (isLike bool, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
