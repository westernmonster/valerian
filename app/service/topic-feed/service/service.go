package service

import (
	"context"
	"net/url"
	"strings"

	"valerian/app/service/feed/def"
	"valerian/app/service/topic-feed/conf"
	"valerian/app/service/topic-feed/dao"
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

	if err := s.mq.QueueSubscribe(def.BusCatalogArticleAdded, "topic-feed", s.onArticleAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusCatalogArticleAdded, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusCatalogArticleDeleted, "topic-feed", s.onArticleDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusCatalogArticleDeleted, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusArticleUpdated, "topic-feed", s.onArticleUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleUpdated, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicTaxonomyCatalogAdded, "topic-feed", s.onTopicTaxonomyCatalogAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicTaxonomyCatalogAdded, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicTaxonomyCatalogMoved, "topic-feed", s.onTopicTaxonomyCatalogMoved); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicTaxonomyCatalogMoved, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicTaxonomyCatalogRenamed, "topic-feed", s.onTopicTaxonomyCatalogRenamed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicTaxonomyCatalogRenamed, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicTaxonomyCatalogDeleted, "topic-feed", s.onTopicTaxonomyCatalogDeleted); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicTaxonomyCatalogDeleted, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicUpdated, "topic-feed", s.onTopicUpdated); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicUpdated, "topic-feed")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicFollowed, "topic-feed", s.onTopicFollowed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicFollowed, "topic-feed")
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

func PromError(name string, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.Errorf(format, args...)
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
