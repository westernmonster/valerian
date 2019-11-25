package http

import (
	"valerian/app/interface/file/conf"
	"valerian/app/interface/file/model"
	"valerian/app/interface/file/service"
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
	g := e.Group("/api/v1/file")
	{
		g.POST("/oss_token", authSvc.User, ossToken)
		g.GET("/sts_cred", authSvc.User, stsCred)
	}
}

// @Summary 获取阿里云OSS STS上传凭证
// @Description 获取阿里云OSS STS上传凭证
// @Tags common
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {object}  app.interface.file.model.STSResp "Token"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /file/sts_cred [post]
func stsCred(c *mars.Context) {
	c.JSON(srv.AssumeRole(c))
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
// @Param req body  app.interface.file.model.ArgOSSToken true "请求"
// @Success 200 {object}  app.interface.file.model.PolicyToken "Token"
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
