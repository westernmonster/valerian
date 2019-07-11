package http

import (
	"valerian/app/conf"
	"valerian/app/interface/editor/model"
	"valerian/app/interface/editor/service"
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
		g.POST("/editor/link_info", linkInfo)
	}
}

// @Summary 获取链接信息
// @Description 获取链接信息
// @Tags editor
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgLinkInfo true "请求"
// @Success 200 {object} model.LinkInfoResp "Token"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /editor/link_info [post]
func linkInfo(c *mars.Context) {
	arg := new(model.ArgLinkInfo)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.LinkInfo(c, arg))
}
