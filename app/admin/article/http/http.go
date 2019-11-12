package http

import (
	"valerian/app/admin/article/conf"
	"valerian/app/admin/article/service"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/permit"
)

var (
	srv       *service.Service
	permitSvc *permit.Permit
)

// Init init
func Init(c *conf.Config, s *service.Service) {
	srv = s
	permitSvc = permit.New(&permit.Config{
		Session: c.Session,
	})

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
	g := e.Group("/api/v1/admin/article", permitSvc.Verify())
	{
		g.GET("/get", getArticle)
		g.POST("/add", addArticle)
		g.POST("/edit", editArticle)
		g.POST("/del", delArticle)

		g.GET("/history/get", articleHistory)

		g.POST("/files", editArticleFiles)
		g.POST("/relations/add", addArticleRelation)
		g.POST("/relations/edit", editArticleRelation)
		g.POST("/relations/del", delArticleRelation)
		g.POST("/relations/primary", setArticleRelationPrimary)

		g.POST("/revise/add", addRevise)
		g.POST("/revise/edit", updateRevise)
		g.POST("/revise/del", delRevise)
		g.GET("/revise/get", getRevise)

		g.GET("/list/files", articleFiles)
		g.GET("/list/histories", articleHistories)
		g.GET("/list/relations", articleRelations)
		g.GET("/list/revises", getRevises)
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
