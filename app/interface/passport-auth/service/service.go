package service

import (
	"context"

	"valerian/app/interface/passport-auth/conf"
	"valerian/app/interface/passport-auth/dao"
	"valerian/app/interface/passport-auth/model"
	"valerian/library/log"
)

var (
	_noLogin = &model.AuthReply{
		Login: false,
	}
)

// Service struct of service
type Service struct {
	c      *conf.Config
	d      *dao.Dao
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
	return
}

// Close dao.
func (s *Service) Close() {
	return
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
