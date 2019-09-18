package http

import (
	"fmt"
	"valerian/app/admin/config/conf"
	"valerian/app/admin/config/service"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

var (
	// vfySvr *verify.Verify
	svr *service.Service
)

// Init init config.
func Init(c *conf.Config, s *service.Service) {
	svr = s
	// vfySvr = verify.New(nil)
	engine := mars.DefaultServer(c.Mars)
	innerRouter(engine)
	if err := engine.Start(); err != nil {
		log.Errorf("engine.Start error(%v)", err)
		panic(err)
	}

	fmt.Println("http-success")
}

func innerRouter(e *mars.Engine) {
	e.Ping(ping)
	e.Register(register)
	b := e.Group("/x/admin/config")
	{
		app := b.Group("/app")
		{
			app.POST("/add", createApp)
		}

		config := b.Group("/config")
		{
			config.POST("/add", createConfig)
		}

		commonConfig := b.Group("/common_config")
		{
			commonConfig.POST("/add", createCommonConfig)
		}

		build := b.Group("/build")
		{
			build.POST("/add", createBuild)
		}

		tag := b.Group("/tag")
		{
			tag.POST("/add", createTag)
		}
	}
}

// ping check server ok.
func ping(c *mars.Context) {
	fmt.Println("ping")
	if err := svr.Ping(c); err != nil {
		log.For(c).Error(fmt.Sprintf("service ping error(%v)", err))
		c.JSON(nil, ecode.ServiceUnavailable)
		return
	}

	c.JSON(nil, nil)
}

func register(c *mars.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
