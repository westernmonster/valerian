package http

import (
	"valerian/app/interface/certification/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 提交工作认证
// @Description 提交工作认证
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.certification.model.ArgWorkCert true "请求"
// @Success 200 "提交成功"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /certification/work [post]
func reqWorkCert(c *mars.Context) {
	arg := new(model.ArgWorkCert)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.RequestWorkCert(c, arg))
}

// @Summary 获取工作认证信息
// @Description 获取工作认证信息
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {object}  app.interface.certification.model.WorkCertResp "认证信息"
// @Failure 42 "尚未发起身份认证"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /certification/work [get]
func workCert(c *mars.Context) {
	c.JSON(srv.GetWorkCert(c))
}
