package http

import (
	"valerian/app/conf"
	"valerian/app/interface/feedback/service"
	"valerian/library/net/http/mars"
)

var (
	srv *service.Service
)

// Init init
func Init(c *conf.Config, engine *mars.Engine) {
	srv = service.New(c)

	route(engine)
}

func route(e *mars.Engine) {
	g := e.Group("/api/v1")
	{
		g.GET("/list/feedback_types", feedbackTypes)
		g.POST("/feedbacks", addFeedback)
	}
}
