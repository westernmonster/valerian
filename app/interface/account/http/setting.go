package http

import "valerian/library/net/http/mars"

// @Summary 更新动态设置
// @Description 更新动态设置
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgActivitySetting true "请求"
// @Success 200 "成功revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/activity_setting/edit [post]
func updateActivitySetting(c *mars.Context) {
}

// @Summary 更新通知设置
// @Description 更新通知设置
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgNotifySetting true "请求"
// @Success 200 "成功revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/notify_setting/edit [post]
func updateNotifySetting(c *mars.Context) {
}

// @Summary 更新系统语言
// @Description 更新系统语言
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgLanguageSetting true "请求"
// @Success 200 "成功revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/language_setting/edit [post]
func updateLanguage(c *mars.Context) {
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
// @Success 200 {object} model.SettingResp "补充列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/setting/get [get]
func getSetting(c *mars.Context) {
}
