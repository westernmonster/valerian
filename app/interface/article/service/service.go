package service

import (
	"context"
	"valerian/app/conf"
	"valerian/app/interface/article/dao"
	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

type Service struct {
	c *conf.Config
	d interface {
		AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
		GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error)
		UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
		DelArticle(c context.Context, node sqalx.Node, id int64) (err error)

		GetArticleFiles(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleFile, err error)
		AddArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
		UpdateArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
		DelArticleFile(c context.Context, node sqalx.Node, id int64) (err error)

		AddArticleSet(c context.Context, node sqalx.Node, item *model.ArticleSet) (err error)
		DelArticleSet(c context.Context, node sqalx.Node, id int64) (err error)
		GetArticleVersionByName(c context.Context, node sqalx.Node, articleSetID int64, versionName string) (item *model.ArticleVersionResp, err error)

		GetTopicCatalogByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicCatalog, err error)
		GetTopicCatalogMaxChildrenSeq(c context.Context, node sqalx.Node, topicID, parentID int64) (seq int, err error)
		AddTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)

		GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error)

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
