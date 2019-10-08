package service

import (
	"context"
	"net/url"
	"strings"

	"valerian/app/service/topic-feed/conf"
	"valerian/app/service/topic-feed/dao"
	"valerian/app/service/topic-feed/model"
	"valerian/library/conf/env"
	"valerian/library/log"
	"valerian/library/mq"
)

// Service struct of service
type Service struct {
	c      *conf.Config
	d      IDao
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

	if err := s.mq.QueueSubscribe(model.BusArticleAdded, "topic-feed", s.onArticleAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusArticleAdded, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusArticleDeleted, "topic-feed", s.onArticleDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusArticleDeleted, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusReviseAdded, "topic-feed", s.onReviseAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusReviseAdded, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusReviseDeleted, "topic-feed", s.onReviseDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusReviseDeleted, "topic-feed")
		panic(err)
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

func genURL(path string, param url.Values) (uri string, err error) {
	u, err := url.Parse(env.SiteURL + path)
	if err != nil {
		return
	}
	u.RawQuery = param.Encode()

	return u.String(), nil
}
