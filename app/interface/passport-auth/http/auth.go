package http

import (
	"valerian/app/interface/passport-auth/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取新Access Token
// @Description 获取新Access Token
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.passport-auth.model.ArgRenewToken true "请求"
// @Success 200 "成功"
// @Failure 116 "Client 不存在"
// @Failure 17 "RefreshToken 不存在或已过期"
// @Failure 400 "验证请求失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/renew [post]
func renewToken(c *mars.Context) {
	arg := new(model.ArgRenewToken)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.RenewToken(c, arg))
}

// @Summary 退出登录
// @Description 退出登录
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/logout [post]
func logout(c *mars.Context) {
	arg := new(model.ArgLogout)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Logout(c, arg))
}
