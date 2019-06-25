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
		g.GET("/search", auth.User, searchTopics)

		g.POST("/add", auth.User, createTopic)
		g.POST("/edit", auth.User, editTopic)
		g.POST("/del", auth.User, deleteTopic)
		g.POST("/owner", auth.User, changeOwner)
		g.POST("/follow", auth.User, followTopic)

		g.POST("/members", auth.User, editTopicMembers)
		g.POST("/related", auth.User, editTopicRelations)
		g.POST("/catalogs", auth.User, editTopicCatalogs)
		g.POST("/versions", auth.User, addTopicVersion)
		g.POST("/versions/merge", auth.User, mergeTopicVersion)

		g.GET("/list/catalogs", auth.User, topicCatalogs)
		g.GET("/list/members", auth.User, topicMembers)
		g.GET("/list/types", auth.User, topicTypeList)
		g.GET("/list/related", auth.User, relatedTopics)
		g.GET("/list/versions", auth.User, topicVersions)
		g.GET("/list/catalog_taxonomies", auth.User, topicCatalogTaxonomies)
	}

	x := e.Group("/api/v1/article")
	{
		x.GET("/get", auth.User, getArticle)
		x.POST("/add", auth.User, addArticle)
		x.POST("/edit", auth.User, editArticle)
		x.POST("/del", auth.User, delArticle)

		x.POST("/files", auth.User, editArticleFiles)
		x.POST("/relations/add", auth.User, addArticleRelation)
		x.POST("/relations/del", auth.User, delArticleRelation)
		x.POST("/relations/primary", auth.User, setArticleRelationPrimary)
		x.POST("/versions", auth.User, addArticleVersion)
		x.POST("/versions/merge", auth.User, mergeArticleVersion)

		x.GET("/list/files", auth.User, articleFiles)
		x.GET("/list/relations", auth.User, articleRelations)
		x.GET("/list/versions", auth.User, articleVersions)
	}
}
