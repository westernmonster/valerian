package service

import (
	"context"

	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetAccountRelationByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountRelation, err error)
	GetAccountRelationByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountRelation, err error)
	AddAccountRelation(c context.Context, node sqalx.Node, item *model.AccountRelation) (err error)
	UpdateAccountRelation(c context.Context, node sqalx.Node, item *model.AccountRelation) (err error)
	DelAccountRelation(c context.Context, node sqalx.Node, id int64) (err error)

	GetFansPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.AccountRelation, err error)
	GetFollowingsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.AccountRelation, err error)

	SetFollowingsCache(c context.Context, aid int64, page, pageSize int, data []*model.AccountRelation) (err error)
	FollowingsCache(c context.Context, aid int64, page, pageSize int) (data []*model.AccountRelation, err error)
	DelFollowingsCache(c context.Context, aid int64) (err error)

	SetFansCache(c context.Context, aid int64, page, pageSize int, data []*model.AccountRelation) (err error)
	FansCache(c context.Context, aid int64, page, pageSize int) (data []*model.AccountRelation, err error)
	DelFansCache(c context.Context, aid int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
