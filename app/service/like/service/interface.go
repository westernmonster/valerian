package service

import (
	"context"

	"valerian/app/service/like/model"
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

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
