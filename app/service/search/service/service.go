package service

import (
	"context"
	"net/url"
	"strings"

	"valerian/app/service/feed/def"
	"valerian/app/service/search/conf"
	"valerian/app/service/search/dao"
	"valerian/library/conf/env"
	"valerian/library/log"
	"valerian/library/mq"
)

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

	if err := s.mq.QueueSubscribe(def.BusAccountAdded, "search", s.onAccountAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusAccountAdded, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusAccountUpdated, "search", s.onAccountUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusAccountUpdated, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusAccountDeleted, "search", s.onAccountDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusAccountDeleted, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleAdded, "search", s.onArticleAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleAdded, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleUpdated, "search", s.onArticleUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleUpdated, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleDeleted, "search", s.onArticleDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleDeleted, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicAdded, "search", s.onTopicAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicAdded, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicUpdated, "search", s.onTopicUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicUpdated, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicDeleted, "search", s.onTopicDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicDeleted, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionAdded, "search", s.onDiscussionAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionAdded, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionUpdated, "search", s.onDiscussionUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionUpdated, "search")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionDeleted, "search", s.onDiscussionDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionDeleted, "search")
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
	s.mq.Close()
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
