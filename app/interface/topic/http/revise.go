package http

import "valerian/library/net/http/mars"

// @Summary 获取话题补充列表
// @Description 获取话题补充列表
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "topic_id"
// @Success 200 {object} model.ReviseListResp "讨论列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/list/discussions [get]
func getDiscusstions(c *mars.Context) {
}
