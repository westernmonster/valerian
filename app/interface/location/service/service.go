package service

import (
	"context"

	"valerian/app/conf"
	"valerian/app/interface/location/dao"
	"valerian/app/interface/location/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetAllCountryCodes(c context.Context, node sqalx.Node) (items []*model.CountryCode, err error)

		SetCountryCodesCache(c context.Context, items []*model.CountryCode) (err error)
		CountryCodesCache(c context.Context) (res []*model.CountryCode, err error)
		DelCountryCodesCache(c context.Context, token string) (err error)

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
