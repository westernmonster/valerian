package service

import (
	"context"
	"net/url"
	"strings"

	"valerian/app/service/feed/def"
	"valerian/app/service/message/conf"
	"valerian/app/service/message/dao"
	"valerian/library/conf/env"
	"valerian/library/jpush"
	"valerian/library/log"
	"valerian/library/mq"
)

// Service struct of service
type Service struct {
	c      *conf.Config
	d      IDao
	mq     *mq.MessageQueue
	jp     *jpush.JpushClient
	missch chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		mq:     mq.New(env.Hostname, c.Nats),
		missch: make(chan func(), 1024),
		jp:     jpush.NewJpushClient(c.JPush),
	}

	if err := s.mq.QueueSubscribe(def.BusArticleLiked, "message", s.onArticleLiked); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleLiked, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseLiked, "message", s.onReviseLiked); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseLiked, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionLiked, "message", s.onDiscussionLiked); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionLiked, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusCommentLiked, "message", s.onCommentLiked); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusCommentLiked, "message")
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
