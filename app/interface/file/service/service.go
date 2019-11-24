package service

import (
	"context"

	"valerian/app/interface/file/conf"
	"valerian/library/log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/pkg/errors"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
	}
	missch    chan func()
	stsClient *sts.Client
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		missch: make(chan func(), 1024),
	}
	go s.cacheproc()

	if client, err := sts.NewClientWithAccessKey(c.Aliyun.RegionID, c.Aliyun.AccessKeyID, c.Aliyun.AccessKeySecret); err != nil {
		panic(errors.WithMessage(err, "Failed to init Aliyun STS Client"))
	} else {
		s.stsClient = client
	}
	return
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return nil
}

// Close dao.
func (s *Service) Close() {
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
