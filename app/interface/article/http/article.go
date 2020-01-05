package http

import (
	"strconv"
	"strings"
	"valerian/app/interface/article/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 新增文章
// @Description 新增文章
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.article.model.ArgAddArticle true "请求"
// @Success 200 "成功,返回article_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/add [post]
func addArticle(c *mars.Context) {
	arg := new(model.ArgAddArticle)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if id, err := srv.AddArticle(c, arg); err != nil {
		c.JSON(nil, err)
	} else {
		c.JSON(strconv.FormatInt(id, 10), err)
	}
}

// @Summary 更新文章
// @Description 更新文章
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.article.model.ArgUpdateArticle true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/edit [post]
func editArticle(c *mars.Context) {
	arg := new(model.ArgUpdateArticle)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateArticle(c, arg))
}

// @Summary 删除文章
// @Description 删除文章
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.article.model.ArgDelete true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/del [post]
func delArticle(c *mars.Context) {
	arg := new(model.ArgDelete)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.DelArticle(c, arg.ID))
}

// @Summary 获取文章
// @Description 获取文章
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Param cur_topic_id query string true "当前话题ID，用于决定是否返回上下条连续分页"
// @Param include query string true  "目前支持：files,relations,histories,meta"
// @Param updated_at query string false  "app缓存的更新时间戳"
// @Success 200 {object}  app.interface.article.model.ArticleResp "文章"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/get [get]
func getArticle(c *mars.Context) {
	include := c.Request.Form.Get("include")
	idStr := c.Request.Form.Get("id")
	curTopicStr := c.Request.Form.Get("cur_topic_id")
	updatedAt := c.Request.Form.Get("updated_at")
	updatedAtTimeStamp, _ := strconv.ParseInt(updatedAt, 10, 64)

	var err error
	var id int64
	var curTopicID int64

	if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if strings.TrimSpace(curTopicStr) != "" {
		if curTopicID, err = strconv.ParseInt(curTopicStr, 10, 64); err != nil {
			c.JSON(nil, ecode.RequestErr)
			return
		}
	}

	articleResp, err := srv.GetArticle(c, id, curTopicID, include)
	// 传入updatedAtTimeStamp 的时候判断是否已经修改过用于更新 app 的缓存
	if err == nil && updatedAtTimeStamp == articleResp.UpdatedAt {
		c.JSON(nil, ecode.NotModified)
		return
	}
	c.JSON(articleResp, err)
}

// @Summary 有编辑权限的文章列表
// @Description 有编辑权限的文章列表
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param query query string true "查询条件"
// @Param ps query integer false "每页大小"
// @Param pn query integer false "页码 1开始"
// @Success 200 {object}  app.interface.article.model.ArticleListResp "文章"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/list/has_edit_permission [get]
func getHasEditPermissionArticles(c *mars.Context) {
	var (
		err error
		pn  int
		ps  int
	)

	params := c.Request.Form

	if pn, err = strconv.Atoi(params.Get("pn")); err != nil {
		pn = 1
	} else if pn < 0 {
		pn = 1
	}

	if ps, err = strconv.Atoi(params.Get("ps")); err != nil {
		ps = 10
	} else if ps < 0 {
		ps = 10
	}

	c.JSON(srv.GetUserCanEditArticles(c, params.Get("query"), pn, ps))
}
