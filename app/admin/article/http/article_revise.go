package http

import (
	"strconv"
	"valerian/app/admin/article/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取文章补充列表
// @Description 获取文章补充列表
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param article_id query string true "article_id"
// @Param sort query string true "hot,recent"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.admin.article.model.ReviseListResp "补充列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/article/list/revises [get]
func getRevises(c *mars.Context) {
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
	s := "hot"
	sortStr := params.Get("sort")
	if sortStr == "" {
		s = "hot"
	} else {
		s = sortStr
	}

	switch s {
	case "recent":
	case "hot":
		break
	default:
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetArticleRevisesPaged(c, id, s, limit, offset))
}

// @Summary 新增文章补充
// @Description 新增文章补充
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.article.model.ArgAddRevise true "请求"
// @Success 200 "成功,返回revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/article/revise/add [post]
func addRevise(c *mars.Context) {
	arg := new(model.ArgAddRevise)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if id, err := srv.AddRevise(c, arg); err != nil {
		c.JSON(nil, err)
	} else {
		c.JSON(strconv.FormatInt(id, 10), err)
	}
}

// @Summary 删除文章补充
// @Description 删除文章补充
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.article.model.ArgDelete true "请求"
// @Success 200 "成功,返回revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/article/revise/del [post]
func delRevise(c *mars.Context) {
	arg := new(model.ArgDelete)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.DelRevise(c, arg.ID))
}

// @Summary 更新文章补充
// @Description 更新文章补充
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.article.model.ArgUpdateRevise true "请求"
// @Success 200 "成功,返回revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/article/revise/edit [post]
func updateRevise(c *mars.Context) {
	arg := new(model.ArgUpdateRevise)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateRevise(c, arg))
}

// @Summary 获取文章补充详情
// @Description 获取文章补充详情
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "id"
// @Success 200 {object}  app.admin.article.model.ReviseDetailResp "补充列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/article/revise/get [get]
func getRevise(c *mars.Context) {
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(srv.GetRevise(c, id))
	}
}
