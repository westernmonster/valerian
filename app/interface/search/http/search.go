package http

import (
	"valerian/app/interface/search/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 搜索话题
// @Description 搜索话题
// @Tags search
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param kw query string true "关键词"
// @Param kw_fields query string false "查询关键词的字段列表，逗号分隔"
// @Param order query string false "sort中字段排序的顺序(desc,asc) 以逗号分隔"
// @Param sort query string false "排序的字段，逗号分隔"
// @Param pn query integer true "页码"
// @Param ps query integer true "每页大小"
// @Param debug query bool false "debug"
// @Param source query bool false "所需要展示的字段，逗号分隔"
// @Success 200 {object} model.TopicSearchResult "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/topics [get]
func searchTopics(c *mars.Context) {
	arg := new(model.TopicSearchParams)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.TopicSearch(c, arg))
}

// @Summary 搜索账户
// @Description 搜索账户
// @Tags search
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param kw query string true "关键词"
// @Param kw_fields query string false "查询关键词的字段列表，逗号分隔"
// @Param order query string false "sort中字段排序的顺序(desc,asc) 以逗号分隔"
// @Param sort query string false "排序的字段，逗号分隔"
// @Param pn query integer true "页码"
// @Param ps query integer true "每页大小"
// @Param debug query bool false "debug"
// @Param source query bool false "所需要展示的字段，逗号分隔"
// @Success 200 {object} model.AccountSearchResult "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/accounts [get]
func searchAccounts(c *mars.Context) {
	arg := new(model.AccountSearchParams)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.AccountSearch(c, arg))
}

// @Summary 搜索文章
// @Description 搜索文章
// @Tags search
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param kw query string true "关键词"
// @Param kw_fields query string false "查询关键词的字段列表，逗号分隔"
// @Param order query string false "sort中字段排序的顺序(desc,asc) 以逗号分隔"
// @Param sort query string false "排序的字段，逗号分隔"
// @Param pn query integer true "页码"
// @Param ps query integer true "每页大小"
// @Param debug query bool false "debug"
// @Param source query bool false "所需要展示的字段，逗号分隔"
// @Success 200 {object} model.ArticleSearchResult "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/articles [get]
func searchArticles(c *mars.Context) {
	arg := new(model.ArticleSearchParams)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.ArticleSearch(c, arg))
}

// @Summary 搜索讨论
// @Description 搜索讨论
// @Tags search
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param kw query string true "关键词"
// @Param kw_fields query string false "查询关键词的字段列表，逗号分隔"
// @Param order query string false "sort中字段排序的顺序(desc,asc) 以逗号分隔"
// @Param sort query string false "排序的字段，逗号分隔"
// @Param pn query integer true "页码"
// @Param ps query integer true "每页大小"
// @Param debug query bool false "debug"
// @Param source query bool false "所需要展示的字段，逗号分隔"
// @Success 200 {object} model.DiscussSearchResult "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/discussions [get]
func searchDiscusstions(c *mars.Context) {
	arg := new(model.DiscussionSearchParams)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.DiscussionSearch(c, arg))
}

// @Summary 综合搜索
// @Description 综合搜索
// @Tags search
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param kw query string true "关键词"
// @Success 200 {object} model.AllSearchResult "结果"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/all [get]
func searchAll(c *mars.Context) {
}
