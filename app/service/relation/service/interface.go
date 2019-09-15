package service

import (
	"context"

	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetAccountRelationByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountRelation, err error)
	GetAccountRelationByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountRelation, err error)
	AccountRelations(c context.Context, node sqalx.Node, aid int64) (items []*model.AccountRelation, err error)
	AccountRelationsIN(c context.Context, node sqalx.Node, aid int64, fids []int64) (items []*model.AccountRelation, err error)
	AddAccountRelation(c context.Context, node sqalx.Node, item *model.AccountRelation) (err error)
	UpdateAccountRelation(c context.Context, node sqalx.Node, item *model.AccountRelation) (err error)
	DelAccountRelation(c context.Context, node sqalx.Node, id int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
