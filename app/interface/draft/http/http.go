package http

import (
	"valerian/app/conf"
	"valerian/app/interface/auth"
	"valerian/app/interface/draft/service"
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
	g := e.Group("/api/v1/draft", auth.User)
	{
		g.GET("/category/add", addDraftCategory)
		g.GET("/category/edit", updateDraftCategory)
		g.GET("/category/del", delDraftCategory)
		g.GET("/list/categories", draftCategories)

		g.GET("/color/add", addColor)
		g.GET("/color/edit", updateColor)
		g.GET("/color/del", delColor)
		g.GET("/list/colors", colors)
	}
}
