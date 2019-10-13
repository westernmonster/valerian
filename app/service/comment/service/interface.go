package service

import (
	"context"

	account "valerian/app/service/account/api"
	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetCommentByID(c context.Context, node sqalx.Node, id int64) (item *model.Comment, err error)
	GetCommentStatByID(c context.Context, node sqalx.Node, commentID int64) (item *model.CommentStat, err error)
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
