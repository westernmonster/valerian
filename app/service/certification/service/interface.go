package service

import (
	"context"
	"valerian/app/service/certification/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetIDCertificationsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.IDCertification, err error)
	GetIDCertifications(c context.Context, node sqalx.Node) (items []*model.IDCertification, err error)
	GetIDCertificationByID(c context.Context, node sqalx.Node, id int64) (item *model.IDCertification, err error)
	GetIDCertificationByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.IDCertification, err error)
	AddIDCertification(c context.Context, node sqalx.Node, item *model.IDCertification) (err error)
	UpdateIDCertification(c context.Context, node sqalx.Node, item *model.IDCertification) (err error)
	DelIDCertification(c context.Context, node sqalx.Node, id int64) (err error)

	GetWorkCertificationsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.WorkCertification, err error)
	GetWorkCertifications(c context.Context, node sqalx.Node) (items []*model.WorkCertification, err error)
	GetWorkCertificationByID(c context.Context, node sqalx.Node, id int64) (item *model.WorkCertification, err error)
	GetWorkCertificationByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.WorkCertification, err error)
	AddWorkCertification(c context.Context, node sqalx.Node, item *model.WorkCertification) (err error)
	UpdateWorkCertification(c context.Context, node sqalx.Node, item *model.WorkCertification) (err error)
	DelWorkCertification(c context.Context, node sqalx.Node, id int64) (err error)

	GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)
	UpdateAccount(c context.Context, node sqalx.Node, item *model.Account) (err error)

	AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
	SetAccountCache(c context.Context, m *model.Account) (err error)
	DelAccountCache(c context.Context, accountID int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
