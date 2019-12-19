package service

import (
	"context"
	"net/url"
	"strings"

	"valerian/app/service/account-feed/conf"
	"valerian/app/service/account-feed/dao"
	"valerian/app/service/feed/def"
	"valerian/library/conf/env"
	"valerian/library/log"
	"valerian/library/mq"
	"valerian/library/stat/prom"
)

// Service struct of service
type Service struct {
	c      *conf.Config
	d      *dao.Dao
	mq     *mq.MessageQueue
	prom   *prom.Prom
	missch chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		mq:     mq.New(env.Hostname, c.Nats),
		prom:   prom.New().WithTimer("account-feed", []string{"method"}),
		missch: make(chan func(), 1024),
	}

	if err := s.mq.QueueSubscribe(def.BusArticleAdded, "account-feed", s.onArticleAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleAdded, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleUpdated, "account-feed", s.onArticleUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleUpdated, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseAdded, "account-feed", s.onReviseAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseAdded, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseUpdated, "account-feed", s.onReviseUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseUpdated, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseDeleted, "account-feed", s.onReviseDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseDeleted, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionAdded, "account-feed", s.onDiscussionAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionAdded, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionUpdated, "account-feed", s.onDiscussionUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionUpdated, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionDeleted, "account-feed", s.onDiscussionDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionDeleted, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicAdded, "account-feed", s.onTopicAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicAdded, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicFollowed, "account-feed", s.onTopicFollowed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicFollowed, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicDeleted, "account-feed", s.onTopicDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicDeleted, "account-feed")
		panic(err)
	}

	go s.cacheproc()
	return
}

func PromError(name string, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.Errorf(format, args...)
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
