package service

import (
	"context"
	"sync"

	"valerian/app/admin/config/conf"
	"valerian/app/admin/config/dao"
	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"

	confrpc "valerian/app/infra/config/rpc/client"
)

// Service service
type Service struct {
	c *conf.Config

	// rpcconf config service Rpc
	confSvr *confrpc.Service

	d interface {
		GetAppsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.App, err error)
		GetApps(c context.Context, node sqalx.Node) (items []*model.App, err error)
		GetAppByID(c context.Context, node sqalx.Node, id int64) (item *model.App, err error)
		GetAppByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.App, err error)
		AddApp(c context.Context, node sqalx.Node, item *model.App) (err error)
		UpdateApp(c context.Context, node sqalx.Node, item *model.App) (err error)
		DelApp(c context.Context, node sqalx.Node, id int64) (err error)

		GetConfigsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Config, err error)
		GetConfigs(c context.Context, node sqalx.Node) (items []*model.Config, err error)
		GetConfigByID(c context.Context, node sqalx.Node, id int64) (item *model.Config, err error)
		GetConfigByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Config, err error)
		AddConfig(c context.Context, node sqalx.Node, item *model.Config) (err error)
		UpdateConfig(c context.Context, node sqalx.Node, item *model.Config) (err error)
		DelConfig(c context.Context, node sqalx.Node, id int64) (err error)

		Ping(c context.Context) (err error)
		Close()
		ApmDB() sqalx.Node
		ConfigDB() sqalx.Node
	}

	cLock sync.RWMutex
	// cache map[string]*model.CacheData
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:       c,
		confSvr: confrpc.New(c.ConfSvr),
		d:       dao.New(c),
	}
	return
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return s.d.Ping(c)
}

// Close dao.
func (s *Service) Close() {
	s.d.Close()
}
