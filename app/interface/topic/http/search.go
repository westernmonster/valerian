package http

import (
	"valerian/app/interface/topic/model"
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
// @Param req body model.TopicSearchParams true "请求"
// @Success 200 {object} model.TopicSearchResult "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/topics [post]
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
// @Param req body model.AccountSearchParams true "请求"
// @Success 200 {object} model.AccountSearchResult "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/accounts [post]
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
// @Param req body model.ArticleSearchParams true "请求"
// @Success 200 {object} model.ArticleSearchResult "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/articles [post]
func searchArticles(c *mars.Context) {
}

// @Summary 搜索讨论
// @Description 搜索讨论
// @Tags search
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.DiscussSearchParams true "请求"
// @Success 200 {object} model.DiscussSearchResult "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /search/discussions [post]
func searchDiscusstions(c *mars.Context) {
}
