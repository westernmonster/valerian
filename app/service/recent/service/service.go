package service

import (
	"context"
	"strings"

	"valerian/app/service/feed/def"
	"valerian/app/service/recent/conf"
	"valerian/app/service/recent/dao"
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
		missch: make(chan func(), 1024),
		mq:     mq.New(env.Hostname, c.Nats),
	}

	if err := s.mq.QueueSubscribe(def.BusArticleAdded, "recent", s.onArticleAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleAdded, "recent")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleViewed, "recent", s.onArticleViewed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleViewed, "recent")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleDeleted, "recent", s.onArticleDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleDeleted, "recent")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseAdded, "recent", s.onReviseAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseAdded, "recent")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseDeleted, "recent", s.onReviseDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseDeleted, "recent")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionAdded, "recent", s.onDiscussionAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionAdded, "recent")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionDeleted, "recent", s.onDiscussionDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionDeleted, "recent")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicDeleted, "recent", s.onTopicDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicDeleted, "recent")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicViewed, "recent", s.onTopicViewed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicViewed, "recent")
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
