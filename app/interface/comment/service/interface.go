package service

import (
	"context"

	"valerian/app/interface/comment/model"
	account "valerian/app/service/account/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetCommentByID(c context.Context, node sqalx.Node, id int64) (item *model.Comment, err error)
	AddComment(c context.Context, node sqalx.Node, item *model.Comment) (err error)
	UpdateComment(c context.Context, node sqalx.Node, item *model.Comment) (err error)
	DelComment(c context.Context, node sqalx.Node, id int64) (err error)

	GetCommentStatByID(c context.Context, node sqalx.Node, commentID int64) (item *model.CommentStat, err error)
	AddCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error)
	UpdateCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error)
	IncrCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error)

	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
