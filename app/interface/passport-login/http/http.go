package http

import (
	"valerian/app/conf"
	"valerian/app/interface/passport-login/service"
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
	g := e.Group("/oauth")
	{
		g.POST("/login/mobile", mobileLogin)
		g.POST("/login/digit", digitLogin)
		g.POST("/login/email", emailLogin)
	}
}
