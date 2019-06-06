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
		g.POST("/file/oss_token", auth.User, ossToken)
	}
}

// @Summary 获取阿里云OSS上传TOKEN
// @Description 获取阿里云OSS上传TOKEN
// @Description 阿里云文档：https://help.aliyun.com/document_detail/31926.html
// @Tags common
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgOSSToken true "请求"
// @Success 200 {object} model.PolicyToken "Token"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /file/oss_token [post]
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
