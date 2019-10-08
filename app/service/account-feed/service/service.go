package service

import (
	"context"
	"net/url"
	"strings"

	"valerian/app/service/account-feed/conf"
	"valerian/app/service/account-feed/dao"
	"valerian/app/service/account-feed/model"
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

	if err := s.mq.QueueSubscribe(model.BusArticleAdded, "account-feed", s.onArticleAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusArticleAdded, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusArticleUpdated, "account-feed", s.onArticleUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusArticleUpdated, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusArticleDeleted, "account-feed", s.onArticleDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusArticleDeleted, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusReviseAdded, "account-feed", s.onReviseAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusReviseAdded, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusReviseUpdated, "account-feed", s.onReviseUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusReviseUpdated, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusReviseDeleted, "account-feed", s.onReviseDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusReviseDeleted, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusDiscussionAdded, "account-feed", s.onDiscussionAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusDiscussionAdded, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusDiscussionUpdated, "account-feed", s.onDiscussionUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusDiscussionUpdated, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusDiscussionDeleted, "account-feed", s.onDiscussionDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusDiscussionDeleted, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusTopicAdded, "account-feed", s.onTopicAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusTopicAdded, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusTopicFollowed, "account-feed", s.onTopicFollowed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusTopicFollowed, "account-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(model.BusTopicDeleted, "account-feed", s.onTopicDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, model.BusTopicDeleted, "account-feed")
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
