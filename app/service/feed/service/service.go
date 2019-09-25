package service

import (
	"context"
	"net/url"
	"strings"

	"valerian/app/service/feed/conf"
	"valerian/app/service/feed/dao"
	"valerian/library/conf/env"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

// Service struct of service
type Service struct {
	c            *conf.Config
	d            IDao
	sc           stan.Conn
	missch       chan func()
	feedConsumer *FeedConsumer
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		missch: make(chan func(), 1024),
	}

	servers := strings.Join(c.Nats.Nodes, ",")
	if sc, err := stan.Connect("valerian",
		env.Hostname,
		stan.Pings(10, 5),
		stan.NatsURL(servers),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Errorf("Nats Connection lost, reason: %v", reason)
			panic(reason)
		}),
	); err != nil {
		log.Errorf("connect to servers failed %#v\n", err)
		panic(err)
	} else {
		s.sc = sc
	}

	s.initFeedConsumer()
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
	s.feedConsumer.Unsubscribe()
	s.sc.Close()
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
