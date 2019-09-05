package http

import (
	"valerian/app/interface/account/conf"
	"valerian/app/interface/account/service"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/auth"
)

var (
	srv     *service.Service
	authSvc *auth.Auth
)

// Init init
func Init(c *conf.Config, s *service.Service) {
	srv = s
	authSvc = auth.New(conf.Conf.Auth)

	engine := mars.DefaultServer(c.Mars)
	route(engine)

	if err := engine.Start(); err != nil {
		log.Errorf("engine.Start() error(%v)", err)
		panic(err)
	}
}

func route(e *mars.Engine) {
	e.Ping(ping)
	e.Register(register)
	g := e.Group("/api/v1/me")
	{
		g.PUT("/password", authSvc.User, changePassword)
		g.GET("/", authSvc.User, getProfile)
		g.PUT("/edit", authSvc.User, updateProfile)
	}

	x := e.Group("/api/v1/account")
	{
		x.GET("/list/fans", authSvc.User, fans)
		x.GET("/list/follow", authSvc.User, followed)
	}

	y := e.Group("/api/v1/oauth")
	{
		y.PUT("/password/reset", forgetPassword)
		y.PUT("/password/reset/confirm", resetPassword)
	}
}

// ping check server ok.
func ping(c *mars.Context) {
	var err error
	if err = srv.Ping(c); err != nil {
		log.Errorf("service ping error(%v)", err)
		c.JSON(nil, ecode.ServiceUnavailable)
		return
	}

	c.JSON(nil, nil)
}

// register support discovery.
func register(c *mars.Context) {
	c.JSON(map[string]struct{}{}, nil)
}
