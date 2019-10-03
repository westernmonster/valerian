package http

import (
	"strconv"
	"valerian/library/ecode"
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
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {array} model.ArticleHistoryListResp "历史记录"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/list/histories [get]
func articleHistories(c *mars.Context) {
	var (
		id     int64
		err    error
		offset int
		limit  int
	)

	params := c.Request.Form

	if offset, err = strconv.Atoi(params.Get("offset")); err != nil {
		offset = 0
	} else if offset < 0 {
		offset = 0
	}

	if limit, err = strconv.Atoi(params.Get("limit")); err != nil {
		limit = 10
	} else if limit < 0 {
		limit = 10
	}

	if id, err = strconv.ParseInt(params.Get("article_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetArticleHistoriesResp(c, id, limit, offset))
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
	var (
		id  int64
		err error
	)

	params := c.Request.Form
	if id, err = strconv.ParseInt(params.Get("article_history_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetArticleHistoryResp(c, id))
}
