package http

import (
	"valerian/app/interface/discuss/conf"
	"valerian/app/interface/discuss/service"
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
	g := e.Group("/api/v1/discussion")
	{
		g.GET("/list/by_topic", authSvc.User, getDiscusstionsByTopic)
		g.GET("/list/by_account", authSvc.User, getDiscusstionsByAccount)
		g.GET("/list/categories", authSvc.User, discussCategories)
		g.GET("/list/files", authSvc.User, discussionFiles)

		g.POST("/get", authSvc.User, getDiscussion)
		g.POST("/add", authSvc.User, addDiscussion)
		g.POST("/edit", authSvc.User, updateDiscussion)
		g.POST("/del", authSvc.User, delDiscussion)
		g.POST("/categories", authSvc.User, editDiscussCategories)
		g.POST("/files", authSvc.User, editDiscussionFiles)
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
