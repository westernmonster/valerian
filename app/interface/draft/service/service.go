package service

import (
	"context"
	"valerian/app/conf"
	"valerian/app/interface/draft/dao"
	"valerian/app/interface/draft/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetUserDrafts(c context.Context, node sqalx.Node, aid int64) (items []*model.Draft, err error)
		AddDraft(c context.Context, node sqalx.Node, item *model.Draft) (err error)
		UpdateDraft(c context.Context, node sqalx.Node, item *model.Draft) (err error)
		DelDraft(c context.Context, node sqalx.Node, id int64) (err error)

		GetAllDraftCategories(c context.Context, node sqalx.Node) (items []*model.DraftCategory, err error)
		AddDraftCategory(c context.Context, node sqalx.Node, item *model.DraftCategory) (err error)
		UpdateDraftCategory(c context.Context, node sqalx.Node, item *model.DraftCategory) (err error)
		DelDraftCategory(c context.Context, node sqalx.Node, id int64) (err error)

		GetAllColors(c context.Context, node sqalx.Node) (items []*model.Color, err error)
		AddColor(c context.Context, node sqalx.Node, item *model.Color) (err error)
		UpdateColor(c context.Context, node sqalx.Node, item *model.Color) (err error)
		DelColor(c context.Context, node sqalx.Node, id int64) (err error)

		SetColorsCache(c context.Context, items []*model.Color) (err error)
		ColorsCache(c context.Context) (res []*model.Color, err error)
		DelColorsCache(c context.Context) (err error)

		SetDraftCategoriesCache(c context.Context, items []*model.DraftCategory) (err error)
		DraftCategoriesCache(c context.Context) (res []*model.DraftCategory, err error)
		DelDraftCategoriesCache(c context.Context) (err error)

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
