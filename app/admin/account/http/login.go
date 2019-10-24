package http

import "valerian/library/net/http/mars"

// @Summary 获取个人资料
// @Description 获取个人资料
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Param include query string true  "目前支持：settings"
// @Success 200 {object} model.SelfProfile "个人资料"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/me/profile [get]
func getProfile(c *mars.Context) {
	c.JSON(srv.GetSelfProfile(c))
}
