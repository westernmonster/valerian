package http

import "valerian/library/net/http/mars"

// @Summary 获取消息列表
// @Description 获取消息列表
// @Tags dm
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param type query string false "消息类型"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object} model.NotificationResp "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /dm/list [get]
func getActivites(c *mars.Context) {
}
