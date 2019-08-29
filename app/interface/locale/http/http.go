package http

import (
	"net/http"
	"valerian/app/interface/locale/conf"
	"valerian/app/interface/locale/service"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/antispam"
)

var (
	srv  *service.Service
	anti *antispam.Antispam
)

// Init init
func Init(c *conf.Config, s *service.Service) {
	srv = s

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
		g.GET("/list/locales", locales)
	}
}

// ping check server ok.
func ping(c *mars.Context) {
	var err error
	if err = srv.Ping(c); err != nil {
		c.JSON(nil, err)
		log.Errorf("service ping error(%v)", err)
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
	}
}
