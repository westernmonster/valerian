package service

import (
	"context"
	"strings"

	"valerian/app/service/article/conf"
	"valerian/app/service/article/dao"
	"valerian/library/conf/env"
	"valerian/library/log"
	"valerian/library/mq"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/imm"
	"github.com/pkg/errors"
)

// Service struct of service
type Service struct {
	c         *conf.Config
	d         *dao.Dao
	mq        *mq.MessageQueue
	immClient *imm.Client
	missch    chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		mq:     mq.New(env.Hostname, c.Nats),
		missch: make(chan func(), 1024),
	}

	if client, err := imm.NewClientWithAccessKey(c.Aliyun.RegionID, c.Aliyun.AccessKeyID, c.Aliyun.AccessKeySecret); err != nil {
		panic(errors.WithMessage(err, "Failed to init Aliyun IMM Client"))
	} else {
		s.immClient = client
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

func includeParam(include string) (dic map[string]bool) {
	arr := strings.Split(include, ",")
	dic = make(map[string]bool)
	for _, v := range arr {
		dic[v] = true
	}

	return
}
