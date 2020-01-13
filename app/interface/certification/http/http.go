package http

import (
	"valerian/app/interface/certification/conf"
	"valerian/app/interface/certification/service"
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

	a := e.Group("/admin/certification")
	{
		a.POST("/workcert/audit", authSvc.User, auditWorkCert)
		a.GET("/workcert/list", authSvc.User, listWorkCert)
		a.GET("/workcert/history/list", authSvc.User, workCertHistory)
	}

	g := e.Group("/certification")
	{
		g.POST("/id", authSvc.User, idCertificationRequest)
		g.GET("/idcert", authSvc.User, idCert)
		g.GET("/id/status", authSvc.User, idCertificationStatus)
		g.POST("/work", authSvc.User, reqWorkCert)
		g.GET("/workcert", authSvc.User, workCert)
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
