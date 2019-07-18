package http

import (
	"valerian/app/infra/discovery/conf"
	"valerian/app/infra/discovery/service"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

var (
	svr *service.Service
)

// Init init http
func Init(c *conf.Config, s *service.Service) {
	svr = s
	engineInner := mars.DefaultServer(c.Mars.Inner)
	innerRouter(engineInner)
	if err := engineInner.Start(); err != nil {
		log.Errorf("mars.DefaultServer error(%v)", err)
		panic(err)
	}
}

// innerRouter init local router api path.
func innerRouter(e *mars.Engine) {
	group := e.Group("/discovery")
	{
		group.POST("/register", register)
		group.POST("/renew", renew)
		group.POST("/cancel", cancel)
		group.GET("/fetch/all", fetchAll)
		group.GET("/fetch", fetch)
		group.GET("/fetchs", fetchs)
		group.GET("/poll", poll)
		group.GET("/polls", polls)
		//manager
		group.POST("/set", set)
		group.GET("/polling", polling)
		group.GET("/nodes", nodes)
	}
}
