package http

import (
	"valerian/app/interface/account/model"
	"valerian/library/net/http/mars"
)

// @Summary 管理员更改用户资料
// @Description 管理员更改用户资料
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgAdminUpdateProfile true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/account/profile [post]
func adminUpdateProfile(c *mars.Context) {
	arg := new(model.ArgAdminUpdateProfile)
	if e := c.Bind(arg); e != nil {
		return
	}

	c.JSON(nil, srv.AdminUpdateAccount(c, arg))
}

// @Summary 管理员锁定账户
// @Description 管理员锁定账户
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgAdminLockAccount true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/account/lock [post]
func adminLockAccount(c *mars.Context) {
	arg := new(model.ArgAdminLockAccount)
	if e := c.Bind(arg); e != nil {
		return
	}

	c.JSON(nil, srv.AdminLockAccount(c, arg))
}

// @Summary 管理员解锁账户
// @Description 管理员解锁账户
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgAdminLockAccount true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/account/unlock [post]
func adminUnlockAccount(c *mars.Context) {
	arg := new(model.ArgAdminLockAccount)
	if e := c.Bind(arg); e != nil {
		return
	}

	c.JSON(nil, srv.AdminUnlockAccount(c, arg))
}

// @Summary 管理员添加账户
// @Description 管理员添加账户
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgAdminAddAccount true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/account/add [post]
func adminAddAccount(c *mars.Context) {
	arg := new(model.ArgAdminAddAccount)
	if e := c.Bind(arg); e != nil {
		return
	}

	c.JSON(nil, srv.AdminAddAccount(c, arg))
}

// @Summary 管理员注销账户
// @Description 管理员注销账户，不可撤销
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgAdminAddAccount true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/account/deactive [post]
func adminDeactiveAccount(c *mars.Context) {
	arg := new(model.ArgAdminDeactiveAccount)
	if e := c.Bind(arg); e != nil {
		return
	}

	c.JSON(nil, srv.AdminDeactive(c, arg))
}
