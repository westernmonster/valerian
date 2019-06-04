package http

import (
	"valerian/app/conf"
	"valerian/app/interface/auth"
	"valerian/app/interface/topic/service"
	"valerian/library/net/http/mars"
)

var (
	srv *service.Service
)

// Init init
func Init(c *conf.Config, engine *mars.Engine) {
	srv = service.New(c)

	route(engine)
}

func route(e *mars.Engine) {
	g := e.Group("/api/v1/topic")
	{
		g.GET("/get", auth.User, getTopic)
		g.POST("/add", auth.User, createTopic)
		g.POST("/edit", auth.User, editTopic)
		g.POST("/del", auth.User, deleteTopic)

		g.POST("/members", auth.User, editTopicMembers)
		g.POST("/related", auth.User, editTopicRelations)

		g.GET("/list/catalogs", auth.User, topicCatalogs)
		g.GET("/list/members", auth.User, topicMembers)
		g.GET("/list/types", auth.User, topicTypeList)
		g.GET("/list/related", auth.User, relatedTopics)
	}
}
