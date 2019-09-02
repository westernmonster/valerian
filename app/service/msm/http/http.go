package http

import (
	"net/http"
	"valerian/app/service/msm/conf"
	"valerian/app/service/msm/service"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

var (
	// vfySvr *verify.Verify
	svr *service.Service
)

// Init init config.
func Init(c *conf.Config, s *service.Service) {
	svr = s
	// vfySvr = verify.New(nil)
	engine := mars.DefaultServer(c.Mars)
	oldRouter(engine)
	innerRouter(engine)
	if err := engine.Start(); err != nil {
		log.Errorf("engine.Start error(%v)", err)
		panic(err)
	}
}

func oldRouter(e *mars.Engine) {
	e.Ping(ping)
	group := e.Group("/x/v1/msm")
	{
		group.GET("/codes/2", codes)
	}
}

func innerRouter(e *mars.Engine) {
	group := e.Group("/x/internal/msm/v1")
	{
		group.GET("/codes/2", codes)
	}
}

// ping check server ok.
func ping(c *mars.Context) {
	if svr.Ping(c) != nil {
		log.Error("service ping error")
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
