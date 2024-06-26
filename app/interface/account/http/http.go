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

	a := e.Group("/admin/account")
	{
		a.POST("/profile", authSvc.User, adminUpdateProfile)
		a.POST("/lock", authSvc.User, adminLockAccount)
		a.POST("/unlock", authSvc.User, adminUnlockAccount)
		a.POST("/add", authSvc.User, adminAddAccount)
		a.POST("/deactive", authSvc.User, adminDeactiveAccount)
		a.GET("/list", authSvc.User, adminAllAccounts)
	}

	g := e.Group("/me")
	{
		g.PUT("/password", authSvc.User, changePassword)
		g.GET("/profile", authSvc.User, getProfile)
		g.PUT("/edit", authSvc.User, updateProfile)
		g.POST("/activity_setting/edit", authSvc.User, updateActivitySetting)
		g.POST("/notify_setting/edit", authSvc.User, updateNotifySetting)
		g.POST("/language_setting/edit", authSvc.User, updateLanguage)
		g.GET("/setting/get", authSvc.User, getSetting)
		g.POST("/deactive", authSvc.User, deactiveAccount)
	}

	x := e.Group("/account")
	{
		x.POST("/follow", authSvc.User, follow)
		x.POST("/unfollow", authSvc.User, unfollow)

		x.GET("/info", authSvc.User, memberInfo)
		x.GET("/cert", authSvc.User, memberCert)

		x.GET("/list/recent", authSvc.User, recent)
		x.GET("/list/my_topics", authSvc.User, accountTopics)
		x.GET("/list/my_articles", authSvc.User, accountArticles)
		x.GET("/list/my_revises", authSvc.User, accountRevises)
		x.GET("/list/my_discussions", authSvc.User, accountDiscussions)

		x.GET("/list/topics", authSvc.User, memberTopics)
		x.GET("/list/activities", authSvc.User, memberActivites)
		x.GET("/list/followed_topics", authSvc.User, memberFollowedTopics)
		x.GET("/list/managed_topics", authSvc.User, memberManagedTopics)
		x.GET("/list/discussions", authSvc.User, memberDiscuss)
		x.GET("/list/articles", authSvc.User, memberArticles)
		x.GET("/list/fans", authSvc.User, fans)
		x.GET("/list/followings", authSvc.User, followed)
	}

	y := e.Group("/oauth")
	{
		y.PUT("/password/reset", forgetPassword)
		y.PUT("/password/reset/confirm", resetPassword)
	}

	i := e.Group("/admin/certification")
	{
		i.POST("/workcert/audit", authSvc.User, auditWorkCert)
		i.GET("/workcert/list", authSvc.User, listWorkCert)
		i.GET("/workcert/history/list", authSvc.User, workCertHistory)
	}

	m := e.Group("/certification")
	{
		m.POST("/id", authSvc.User, idCertificationRequest)
		m.GET("/idcert", authSvc.User, idCert)
		m.GET("/id/status", authSvc.User, idCertificationStatus)
		m.POST("/work", authSvc.User, reqWorkCert)
		m.GET("/workcert", authSvc.User, workCert)
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
