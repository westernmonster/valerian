package service

import (
	"context"

	"valerian/app/interface/feedback/conf"
	"valerian/app/interface/feedback/dao"
	"valerian/app/interface/feedback/model"

	// "valerian/app/interface/locale/conf"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetFeedbacksByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Feedback, err error)
		GetFeedbacks(c context.Context, node sqalx.Node) (items []*model.Feedback, err error)
		GetFeedbackByID(c context.Context, node sqalx.Node, id int64) (item *model.Feedback, err error)
		GetFeedbackByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Feedback, err error)
		AddFeedback(c context.Context, node sqalx.Node, item *model.Feedback) (err error)
		UpdateFeedback(c context.Context, node sqalx.Node, item *model.Feedback) (err error)
		DelFeedback(c context.Context, node sqalx.Node, id int64) (err error)

		GetFeedbackTypesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.FeedbackType, err error)
		GetFeedbackTypes(c context.Context, node sqalx.Node) (items []*model.FeedbackType, err error)
		GetFeedbackTypeByID(c context.Context, node sqalx.Node, id int32) (item *model.FeedbackType, err error)
		GetFeedbackTypeByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.FeedbackType, err error)
		AddFeedbackType(c context.Context, node sqalx.Node, item *model.FeedbackType) (err error)
		UpdateFeedbackType(c context.Context, node sqalx.Node, item *model.FeedbackType) (err error)
		DelFeedbackType(c context.Context, node sqalx.Node, id int32) (err error)

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
