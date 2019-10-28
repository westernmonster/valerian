package service

import (
	"context"
	"net/url"

	"valerian/app/admin/search/conf"
	"valerian/app/admin/search/dao"
	"valerian/app/admin/search/model"
	search "valerian/app/service/search/api"
	"valerian/library/conf/env"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		SearchTopic(c context.Context, arg *model.BasicSearchParams) (info *search.SearchResult, err error)
		SearchAccount(c context.Context, arg *model.BasicSearchParams) (info *search.SearchResult, err error)
		SearchArticle(c context.Context, arg *model.BasicSearchParams) (info *search.SearchResult, err error)
		SearchDiscussion(c context.Context, arg *model.BasicSearchParams) (info *search.SearchResult, err error)

		GetDiscussionStatByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussionStat, err error)
		GetArticleStatByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleStat, err error)
		GetTopicStatByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicStat, err error)
		GetAccountStatByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountStat, err error)

		Ping(c context.Context) (err error)
		Close()
		DB() sqalx.Node
	}
	missch chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		missch: make(chan func(), 1024),
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

func genURL(path string, param url.Values) (uri string, err error) {
	u, err := url.Parse(env.SiteURL + path)
	if err != nil {
		return
	}
	u.RawQuery = param.Encode()

	return u.String(), nil
}
