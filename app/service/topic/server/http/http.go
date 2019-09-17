package http

import (
	"valerian/app/service/topic/conf"
	"valerian/app/service/topic/service"
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
	g := e.Group("/x/internal/topic")
	{
		g.GET("/get", authSvc.User, getTopic)

		g.POST("/add", authSvc.User, createTopic)
		g.POST("/edit", authSvc.User, editTopic)
		g.POST("/del", authSvc.User, deleteTopic)
		g.POST("/owner", authSvc.User, changeOwner)
		g.POST("/follow", authSvc.User, followTopic)
		g.POST("/leave", authSvc.User, leave)
		g.POST("/invite", authSvc.User, inviteFans)

		g.POST("/discuss_categories", authSvc.User, editDiscussCategories)
		g.POST("/members", authSvc.User, editTopicMembers)
		g.POST("/auth_topics", authSvc.User, editAuthTopics)
		g.POST("/catalogs", authSvc.User, editTopicCatalogs)

		g.GET("/list/discuss_categories", authSvc.User, discussCategories)
		g.GET("/list/catalogs", authSvc.User, topicCatalogs)
		g.GET("/list/members", authSvc.User, topicMembers)
		g.GET("/list/member_fans", authSvc.User, memberFansList)
		g.GET("/list/auth_topics", authSvc.User, authTopics)
		g.GET("/list/catalog_taxonomies", authSvc.User, topicCatalogTaxonomies)
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