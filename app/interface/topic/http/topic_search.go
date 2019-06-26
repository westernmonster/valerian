package http

import "valerian/library/net/http/mars"

// @Summary 搜索话题
// @Description 搜索话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param query query string true "查询条件"
// @Param include query string true "支持字段: items[\*].versions,items[\*].has_catalog_taxonomy"
// @Success 200 {object} model.TopicSearchResp "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/search [get]
func searchTopics(c *mars.Context) {
	params := c.Request.Form
	query := params.Get("query")
	include := params.Get("include")
	c.JSON(srv.SearchTopics(c, query, include))
}
