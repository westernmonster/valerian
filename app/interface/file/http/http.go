package http

import (
	"valerian/app/conf"
	"valerian/app/interface/auth"
	"valerian/app/interface/file/model"
	"valerian/app/interface/file/service"
	"valerian/library/ecode"
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
	g := e.Group("/api/v1")
	{
		g.POST("/files/oss_token", auth.User, ossToken)
	}
}

func ossToken(c *mars.Context) {
	arg := new(model.ArgOSSToken)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetPolicyToken(arg.FileType, arg.FileName))
}
