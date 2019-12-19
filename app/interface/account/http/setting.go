package http

import (
	"valerian/app/interface/account/model"
	"valerian/library/net/http/mars"
)

// @Summary 更新动态设置
// @Description 更新动态设置
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgActivitySetting true "请求"
// @Success 200 "成功revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/activity_setting/edit [post]
func updateActivitySetting(c *mars.Context) {
	arg := new(model.ArgActivitySetting)
	if e := c.Bind(arg); e != nil {
		return
	}

	c.JSON(nil, srv.UpdateActivitySetting(c, arg))
}

// @Summary 更新通知设置
// @Description 更新通知设置
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgNotifySetting true "请求"
// @Success 200 "成功revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/notify_setting/edit [post]
func updateNotifySetting(c *mars.Context) {
	arg := new(model.ArgNotifySetting)
	if e := c.Bind(arg); e != nil {
		return
	}

	c.JSON(nil, srv.UpdateNotifySetting(c, arg))
}

// @Summary 更新系统语言
// @Description 更新系统语言
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgLanguageSetting true "请求"
// @Success 200 "成功revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/language_setting/edit [post]
func updateLanguage(c *mars.Context) {
	arg := new(model.ArgLanguageSetting)
	if e := c.Bind(arg); e != nil {
		return
	}

	c.JSON(nil, srv.UpdateLanguageSetting(c, arg))
}

// @Summary 获取设置
// @Description 获取设置
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param article_id query string true "article_id"
// @Success 200 {object}  app.interface.account.model.SettingResp "补充列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/setting/get [get]
func getSetting(c *mars.Context) {
	c.JSON(srv.GetAccountSetting(c))
}

// @Summary 注销账号
// @Description  输入用户密码，注销账号
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgAnnulAccount true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/annul [post]
func annulAccount(c *mars.Context) {
	arg := new(model.ArgAnnulAccount)
	if e := c.Bind(arg); e != nil {
		return
	}
	c.JSON(nil, srv.AnnulAccount(c, arg))
}
