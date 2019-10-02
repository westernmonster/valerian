package http

import (
	"strconv"
	"valerian/app/interface/article/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 新增关联话题
// @Description 新增关联话题
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgAddArticleRelation true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/relations/add [post]
func addArticleRelation(c *mars.Context) {
	arg := new(model.ArgAddArticleRelation)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.AddArticleRelation(c, arg))
}

// @Summary 更改关联话题
// @Description 更改关联话题
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgUpdateArticleRelation true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/relations/edit [post]
func editArticleRelation(c *mars.Context) {
	arg := new(model.ArgUpdateArticleRelation)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateArticleRelation(c, arg))
}

// @Summary 删除关联话题
// @Description 删除关联话题
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgDelArticleRelation true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/relations/del [post]
func delArticleRelation(c *mars.Context) {
	arg := new(model.ArgDelArticleRelation)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.DelArticleRelation(c, arg))
}

// @Summary 设置主话题
// @Description 设置主话题
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgSetPrimaryArticleRelation true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/relations/primary [post]
func setArticleRelationPrimary(c *mars.Context) {
	arg := new(model.ArgSetPrimaryArticleRelation)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.SetPrimary(c, arg))
}

// @Summary 获取关联话题列表
// @Description 获取关联话题列表
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param article_id query string true "文章ID"
// @Success 200 {array} model.ArticleRelationResp "关联话题"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/list/relations [get]
func articleRelations(c *mars.Context) {
	var (
		articleID int64
		err       error
	)

	params := c.Request.Form
	if articleID, err = strconv.ParseInt(params.Get("article_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetArticleRelations(c, articleID))
}
