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
	"valerian/library/stat/prom"
)

// Service struct of service
type Service struct {
	c      *conf.Config
	d      *dao.Dao
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

	if err := s.mq.QueueSubscribe(def.BusArticleCommented, "message", s.onArticleCommented); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusArticleCommented, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseCommented, "message", s.onReviseCommented); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseCommented, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusDiscussionCommented, "message", s.onDiscussionCommented); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusDiscussionCommented, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusCommentReplied, "message", s.onCommentReplied); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusCommentReplied, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusReviseAdded, "message", s.onReviseAdded); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusReviseAdded, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusMemberFollowed, "message", s.onMemberFollowed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusMemberFollowed, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicFollowed, "message", s.onTopicFollowed); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicFollowed, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicInviteSent, "message", s.onTopicInvite); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicInviteSent, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicFollowRequested, "message", s.onTopicFollowRequested); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicFollowRequested, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicFollowApproved, "message", s.onTopicFollowApproved); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicFollowApproved, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusTopicFollowRejected, "message", s.onTopicFollowRejected); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusTopicFollowRejected, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusFeedBackAccuseSuit, "message", s.onFeedBackAccuseSuit); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusFeedBackAccuseSuit, "message")
		panic(err)
	}

	if err := s.mq.QueueSubscribe(def.BusFeedBackAccuseNotSuit, "message", s.onFeedBackAccuseNotSuit); err != nil {
		log.Errorf("mq.QueueSubscribe(), error(%+v),subject(%s), queue(%s)", err, def.BusFeedBackAccuseNotSuit, "message")
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
