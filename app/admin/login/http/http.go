package http

import (
	"valerian/app/admin/login/conf"
	"valerian/app/admin/login/service"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

var (
	srv *service.Service
	cfg *conf.Config
)

// Init init
func Init(c *conf.Config, s *service.Service) {
	srv = s
	cfg = c

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
	g := e.Group("/api/v1/admin/login")
	{
		g.POST("/mobile", mobileLogin)
		g.POST("/digit", digitLogin)
		g.POST("/email", emailLogin)
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