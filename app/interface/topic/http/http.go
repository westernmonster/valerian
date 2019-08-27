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

		g.POST("/search/accounts", auth.User, searchAccounts)
		g.POST("/search/topics", auth.User, searchTopics)

		g.POST("/add", auth.User, createTopic)
		g.POST("/edit", auth.User, editTopic)
		g.POST("/del", auth.User, deleteTopic)
		g.POST("/owner", auth.User, changeOwner)
		g.POST("/follow", auth.User, followTopic)
		g.POST("/leave", auth.User, leave)
		g.POST("/invite", auth.User, inviteFans)

		g.POST("/discuss_categories", auth.User, editDiscussCategories)
		g.POST("/members", auth.User, editTopicMembers)
		g.POST("/auth_topics", auth.User, editAuthTopics)
		g.POST("/catalogs", auth.User, editTopicCatalogs)

		g.GET("/list/activities", auth.User, getActivites)
		g.GET("/list/discuss_categories", auth.User, discussCategories)
		g.GET("/list/catalogs", auth.User, topicCatalogs)
		g.GET("/list/members", auth.User, topicMembers)
		g.GET("/list/member_fans", auth.User, memberFansList)
		g.GET("/list/auth_topics", auth.User, authTopics)
		g.GET("/list/catalog_taxonomies", auth.User, topicCatalogTaxonomies)
	}

	x := e.Group("/api/v1/article")
	{
		x.GET("/get", auth.User, getArticle)
		x.POST("/add", auth.User, addArticle)
		x.POST("/edit", auth.User, editArticle)
		x.POST("/del", auth.User, delArticle)
		x.POST("/fav", auth.User, favArticle)
		x.POST("/like", auth.User, likeArticle)

		x.POST("/files", auth.User, editArticleFiles)
		x.POST("/relations/add", auth.User, addArticleRelation)
		x.POST("/relations/del", auth.User, delArticleRelation)
		x.POST("/relations/primary", auth.User, setArticleRelationPrimary)

		x.GET("/list/files", auth.User, articleFiles)
		x.GET("/list/relations", auth.User, articleRelations)
	}
}
