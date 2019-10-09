package service

import (
	"context"
	"net/url"
	"strings"

	"valerian/app/service/feed/conf"
	"valerian/app/service/feed/dao"
	"valerian/app/service/feed/def"
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

	if err := s.mq.QueueSubscribe(def.BusArticleAdded, "feed", s.onArticleAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleAdded, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleUpdated, "feed", s.onArticleUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleUpdated, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleDeleted, "feed", s.onArticleDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleDeleted, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseAdded, "feed", s.onReviseAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseAdded, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseUpdated, "feed", s.onReviseUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseUpdated, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseDeleted, "feed", s.onReviseDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseDeleted, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionAdded, "feed", s.onDiscussionAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionAdded, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionUpdated, "feed", s.onDiscussionUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionUpdated, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionDeleted, "feed", s.onDiscussionDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionDeleted, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicAdded, "feed", s.onTopicAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicAdded, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicFollowed, "feed", s.onTopicFollowed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicFollowed, "feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicDeleted, "feed", s.onTopicDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicDeleted, "feed")
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
