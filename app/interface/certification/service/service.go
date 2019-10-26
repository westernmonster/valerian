package service

import (
	"context"

	"valerian/app/interface/certification/conf"
	"valerian/app/interface/certification/dao"
	certification "valerian/app/service/certification/api"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		RequestIDCert(c context.Context, aid int64) (info *certification.RequestIDCertResp, err error)
		RefreshIDCertStatus(c context.Context, aid int64) (info *certification.IDCertStatus, err error)
		GetIDCert(c context.Context, aid int64) (info *certification.IDCertInfo, err error)
		GetIDCertStatus(c context.Context, aid int64) (info *certification.IDCertStatus, err error)
		RequestWorkCert(c context.Context, req *certification.WorkCertReq) (err error)
		AuditWorkCert(c context.Context, req *certification.AuditWorkCertReq) (err error)
		GetWorkCert(c context.Context, aid int64) (info *certification.WorkCertInfo, err error)
		GetWorkCertStatus(c context.Context, aid int64) (info *certification.WorkCertStatus, err error)

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
