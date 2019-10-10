package service

import (
	"context"
	"valerian/app/service/recent/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetUserRecentViewsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.RecentView, err error)
	GetRecentViewsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.RecentView, err error)
	GetRecentViews(c context.Context, node sqalx.Node) (items []*model.RecentView, err error)
	GetRecentViewByID(c context.Context, node sqalx.Node, id int64) (item *model.RecentView, err error)
	GetRecentViewByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.RecentView, err error)
	AddRecentView(c context.Context, node sqalx.Node, item *model.RecentView) (err error)
	UpdateRecentView(c context.Context, node sqalx.Node, item *model.RecentView) (err error)
	DelRecentView(c context.Context, node sqalx.Node, id int64) (err error)

	GetUserRecentPubsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.RecentPub, err error)
	GetRecentPubsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.RecentPub, err error)
	GetRecentPubs(c context.Context, node sqalx.Node) (items []*model.RecentPub, err error)
	GetRecentPubByID(c context.Context, node sqalx.Node, id int64) (item *model.RecentPub, err error)
	GetRecentPubByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.RecentPub, err error)
	AddRecentPub(c context.Context, node sqalx.Node, item *model.RecentPub) (err error)
	UpdateRecentPub(c context.Context, node sqalx.Node, item *model.RecentPub) (err error)
	DelRecentPub(c context.Context, node sqalx.Node, id int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}