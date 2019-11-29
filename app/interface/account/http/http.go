package http

import (
	"valerian/app/interface/account/conf"
	"valerian/app/interface/account/service"
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
	g := e.Group("/api/v1/me")
	{
		g.PUT("/password", authSvc.User, changePassword)
		g.GET("/profile", authSvc.User, getProfile)
		g.PUT("/edit", authSvc.User, updateProfile)
		g.POST("/activity_setting/edit", authSvc.User, updateActivitySetting)
		g.POST("/notify_setting/edit", authSvc.User, updateNotifySetting)
		g.POST("/language_setting/edit", authSvc.User, updateLanguage)
		g.GET("/setting/get", authSvc.User, getSetting)
	}

	x := e.Group("/api/v1/account")
	{
		x.POST("/follow", authSvc.User, follow)
		x.POST("/unfollow", authSvc.User, unfollow)

		x.GET("/info", authSvc.User, memberInfo)
		x.GET("/cert", authSvc.User, memberCert)

		x.GET("/list/recent", authSvc.User, recent)
		x.GET("/list/topics", authSvc.User, memberTopics)
		x.GET("/list/activities", authSvc.User, memberActivites)
		x.GET("/list/followed_topics", authSvc.User, memberFollowedTopics)
		x.GET("/list/managed_topics", authSvc.User, memberManagedTopics)
		x.GET("/list/discussions", authSvc.User, memberDiscuss)
		x.GET("/list/articles", authSvc.User, memberArticles)
		x.GET("/list/fans", authSvc.User, fans)
		x.GET("/list/followings", authSvc.User, followed)
	}

	y := e.Group("/api/v1/oauth")
	{
		y.PUT("/password/reset", forgetPassword)
		y.PUT("/password/reset/confirm", resetPassword)
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
