package http

import (
	"valerian/app/admin/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 禁用用户
// @Description 禁用用户
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgAccountLock true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/account/lock [post]
func setAccountLock(c *mars.Context) {
	arg := new(model.ArgAccountLock)
	if e := c.Bind(arg); e != nil {
		return
	}
	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, srv.AccountLock(c, arg))
}

// @Summary 解锁用户
// @Description 解锁被冻结的用户
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgAccountLock true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/account/unlock [post]
func setAccountUnlock(c *mars.Context) {
	arg := new(model.ArgAccountUnlock)
	if e := c.Bind(arg); e != nil {
		return
	}
	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, srv.AccountUnlock(c, arg))
}
