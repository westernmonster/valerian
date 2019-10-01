package http

import (
	"valerian/app/interface/article/conf"
	"valerian/app/interface/article/service"
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
	g := e.Group("/api/v1/article")
	{
		g.GET("/get", authSvc.User, getArticle)
		g.POST("/add", authSvc.User, addArticle)
		g.POST("/edit", authSvc.User, editArticle)
		g.POST("/del", authSvc.User, delArticle)
		g.POST("/fav", authSvc.User, favArticle)
		g.POST("/like", authSvc.User, likeArticle)

		g.POST("/files", authSvc.User, editArticleFiles)
		g.POST("/relations/add", authSvc.User, addArticleRelation)
		g.POST("/relations/del", authSvc.User, delArticleRelation)
		g.POST("/relations/primary", authSvc.User, setArticleRelationPrimary)

		g.GET("/list/files", authSvc.User, articleFiles)
		g.GET("/list/relations", authSvc.User, articleRelations)
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