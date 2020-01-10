package http

import (
	"valerian/app/interface/common/conf"
	"valerian/app/interface/common/service"
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
	g := e.Group("/api/v1/like")
	{
		g.POST("/like", authSvc.User, like)
		g.POST("/cancel_like", authSvc.User, cancelLike)

		g.POST("/dislike", authSvc.User, dislike)
		g.POST("/cancel_dislike", authSvc.User, cancelDislike)
	}

	editor := e.Group("/api/v1/editor")
	{
		editor.POST("/link_info", linkInfo)
	}

	init := e.Group("/api/v1/init")
	{
		init.GET("/list/major", getMajorData)
		init.GET("/list/related", getRelatedData)
		init.GET("/list/members", getMembersData)
	}

	static := e.Group("/static")
	{
		static.GET("/user-agreement", agreement)
		static.GET("/privacy", privacy)
	}

	e.GET("/api/v1/list/country_codes", countryCodes)
	e.GET("/api/v1/list/locales", locales)

	file := e.Group("/api/v1/file")
	{
		file.POST("/oss_token", authSvc.User, ossToken)
		file.POST("/office_convert", authSvc.User, officeConvert)
		file.POST("/url_upload", authSvc.User, urlUpload)
		file.GET("/sts_cred", authSvc.User, stsCred)
	}

	fav := e.Group("/api/v1/fav")
	{
		fav.GET("/list/all", authSvc.User, favList)
		fav.POST("/fav", authSvc.User, tofav)
		fav.POST("/unfav", authSvc.User, unfav)
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
