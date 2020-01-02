package http

import (
	"valerian/app/interface/article/model"
	"valerian/library/net/http/mars"
)

// @Summary app端缓存Aritcle数据拉取
// @Description app端缓存Aritcle数据拉取
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.article.model.ArgArticleAppCache true "请求"
// @Success 200 {object}  app.interface.article.model.ArticleListCacheResp "文章列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/app_cache/pull [post]
func appArticleCachePull(c *mars.Context) {
	arg := new(model.ArgArticleAppCache)
	if e := c.Bind(arg); e != nil {
		return
	}


	c.JSON( srv.AppArticleCachePull(c, arg))
}


// @Summary app端缓存Revises数据拉取
// @Description app端缓存Revises数据拉取
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.article.model.ArgReviseAppCache true "请求"
// @Success 200 {object}  app.interface.article.model.ReviseDetailListCacheResp "文章补充列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/revise/app_cache/pull [post]
func appReviseCachePull(c *mars.Context) {
	arg := new(model.ArgReviseAppCache)
	if e := c.Bind(arg); e != nil {
		return
	}
	c.JSON(srv.AppReviseCachePull(c, arg))
}