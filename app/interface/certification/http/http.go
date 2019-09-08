package http

import (
	"valerian/app/interface/certification/conf"
	"valerian/app/interface/certification/service"
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
	g := e.Group("/api/v1")
	{
		g.POST("/me/certification/id", authSvc.User, idCertificationRequest)
		g.GET("/me/certification/id/status", authSvc.User, idCertificationStatus)
	}
}
