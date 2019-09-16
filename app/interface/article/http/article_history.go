package http

import (
	"valerian/library/net/http/mars"
)

// @Summary 获取历史记录列表
// @Description 获取历史记录列表
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param article_id query string true "文章ID"
// @Success 200 {array} model.ArticleHistoryResp "历史记录"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/list/histories [get]
func articleHistories(c *mars.Context) {
}

// @Summary 获取历史记录
// @Description 获取历史记录
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param article_history_id query string true "历史记录ID"
// @Success 200 {object} model.ArticleHistoryResp "历史记录"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/history [get]
func articleHistory(c *mars.Context) {
}
