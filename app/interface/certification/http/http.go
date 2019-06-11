package http

import (
	"valerian/app/conf"
	"valerian/app/interface/auth"
	"valerian/app/interface/certification/service"
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
		g.POST("/me/certification/id", auth.User, idCertificationRequest)
		g.GET("/me/certification/id/status", auth.User, idCertificationStatus)
	}
}
