package http

import (
	"valerian/app/conf"
	"valerian/app/interface/account/service"
	"valerian/app/interface/auth"
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
		g.PUT("/me/password", auth.User, changePassword)
		g.GET("/me", auth.User, getProfile)
		g.PUT("/me/edit", auth.User, updateProfile)

		g.PUT("/oauth/password/reset", forgetPassword)
		g.PUT("/oauth/password/reset/confirm", resetPassword)
	}
}
