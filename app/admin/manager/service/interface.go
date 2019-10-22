package service

import (
	"context"
	"valerian/app/admin/manager/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetUsersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.User, err error)
	GetUsers(c context.Context, node sqalx.Node) (items []*model.User, err error)
	GetUserByID(c context.Context, node sqalx.Node, id int64) (item *model.User, err error)
	GetUserByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.User, err error)
	AddUser(c context.Context, node sqalx.Node, item *model.User) (err error)
	UpdateUser(c context.Context, node sqalx.Node, item *model.User) (err error)
	DelUser(c context.Context, node sqalx.Node, id int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
