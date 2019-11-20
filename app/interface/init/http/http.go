package http

import (
	"valerian/app/interface/init/conf"
	"valerian/app/interface/init/service"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"

	"github.com/gobuffalo/packr"
)

var (
	srv *service.Service
	// authSvc *auth.Auth
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
	e.Ping(ping)
	e.Register(register)
	g := e.Group("/api/v1/init")
	{
		g.GET("/list/major", getMajorData)
		g.GET("/list/related", getRelatedData)
		g.GET("/list/members", getMembersData)
	}

	x := e.Group("/static")
	{
		x.GET("/user-agreement", agreement)
	}

}

func agreement(c *mars.Context) {

	data := packr.NewBox("./static").Bytes("user-agreement.html")

	c.Writer.WriteHeader(200)
	c.Writer.Write(data)
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
