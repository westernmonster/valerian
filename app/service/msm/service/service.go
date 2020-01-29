package service

import (
	"container/list"
	"context"
	"sync"
	"sync/atomic"

	confrpc "valerian/app/infra/config/rpc/client"
	"valerian/app/service/msm/conf"
	"valerian/app/service/msm/dao"
	"valerian/app/service/msm/model"
)

const (
	_maxVerNum = 100
)

// Service service
type Service struct {
	c *conf.Config

	// rpcconf config service Rpc
	confSvr *confrpc.Service

	d *dao.Dao
	// ecode
	lock    sync.RWMutex
	version *model.Version
	codes   atomic.Value

	langsLock    sync.RWMutex
	langsVersion *model.Version
	langsCodes   atomic.Value
}

// New new a service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:            c,
		d:            dao.New(c),
		version:      &model.Version{List: list.New(), Map: make(map[int64]*list.Element)},
		langsVersion: &model.Version{List: list.New(), Map: make(map[int64]*list.Element)},
	}
	if err := s.all(); err != nil {
		panic(err)
	}

	if err := s.allLang(); err != nil {
		panic(err)
	}
	go s.updateproc()
	go s.updateLangproc()
	return
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return s.d.Ping(c)
}

// Close close resource
func (s *Service) Close() {
	s.d.Close()
}
