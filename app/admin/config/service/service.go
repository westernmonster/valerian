package service

import (
	"context"
	"sync"

	"valerian/app/admin/config/conf"
	"valerian/app/admin/config/dao"

	confrpc "valerian/app/infra/config/rpc/client"
)

// Service service
type Service struct {
	c *conf.Config

	// rpcconf config service Rpc
	confSvr *confrpc.Service

	d *dao.Dao

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
