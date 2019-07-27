package http

import (
	"valerian/app/interface/locale/conf"
	"valerian/app/interface/locale/service"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

var (
	srv *service.Service
)

// Init init
func Init(c *conf.Config) {
	srv = service.New(c)

	engine := mars.DefaultServer(c.HTTPServer)

	route(engine)

	if err := engine.Start(); err != nil {
		log.Errorf("engineOut.Start() error(%v) | config(%v)", err, c)
		panic(err)
	}
}

func route(e *mars.Engine) {
	g := e.Group("/api/v1")
	{
		g.GET("/list/locales", locales)
	}
}
