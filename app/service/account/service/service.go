package service

import (
	"context"

	"valerian/app/service/account/conf"
	"valerian/app/service/account/dao"
	"valerian/library/conf/env"
	"valerian/library/log"
	"valerian/library/mq"
)

type Service struct {
	c      *conf.Config
	d      *dao.Dao
	mq     *mq.MessageQueue
	missch chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		mq:     mq.New(env.Hostname, c.Nats),
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
