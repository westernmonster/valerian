package service

import (
	"context"

	"valerian/app/service/account/conf"
	"valerian/app/service/account/dao"
	"valerian/app/service/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

type Service struct {
	c *conf.Config
	d interface {
		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)

		GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountStat, err error)

		SetAccountCache(c context.Context, m *model.Account) (err error)
		AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
		DelAccountCache(c context.Context, accountID int64) (err error)
		BatchAccountCache(c context.Context, aids []int64) (cached map[int64]*model.Account, missed []int64, err error)
		SetBatchAccountCache(c context.Context, bs []*model.Account) (err error)

		Ping(c context.Context) (err error)
		Close()
		DB() sqalx.Node
	}

	missch chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		missch: make(chan func(), 1024),
	}
	go s.cacheproc()
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

func (s *Service) addCache(f func()) {
	select {
	case s.missch <- f:
	default:
		log.Warn("cacheproc chan full")
	}
}

// cacheproc is a routine for executing closure.
func (s *Service) cacheproc() {
	for {
		f := <-s.missch
		f()
	}
}
