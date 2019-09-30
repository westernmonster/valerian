package service

import (
	"context"

	"valerian/app/service/fav/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetFavsPaged(c context.Context, node sqalx.Node, aid int64, targetType string, limit, offset int) (items []*model.Fav, err error)
	GetFavsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Fav, err error)
	GetFavs(c context.Context, node sqalx.Node) (items []*model.Fav, err error)
	GetFavByID(c context.Context, node sqalx.Node, id int64) (item *model.Fav, err error)
	GetFavByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Fav, err error)
	AddFav(c context.Context, node sqalx.Node, item *model.Fav) (err error)
	UpdateFav(c context.Context, node sqalx.Node, item *model.Fav) (err error)
	DelFav(c context.Context, node sqalx.Node, id int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
