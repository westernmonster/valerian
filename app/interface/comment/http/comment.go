package http

import (
	"strconv"
	"valerian/app/interface/comment/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 评论列表
// @Description 评论列表
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param type query string true "类型"
// @Param resource_id query string true "目标ID"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object} model.CommentListResp "评论列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /comment/list/comments [get]
func commentList(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("resource_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	ctype := params.Get("type")
	if !model.IsValidTargetType(ctype) {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetCommentsPaged(c, id, ctype, limit, offset))
}

// @Summary 新增评论
// @Description 新增评论
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgAddComment true "请求"
// @Success 200 "成功,返回comment_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /comment/add [post]
func addComment(c *mars.Context) {
	arg := new(model.ArgAddComment)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if id, err := srv.AddComment(c, arg); err != nil {
		c.JSON(nil, err)
	} else {
		c.JSON(strconv.FormatInt(id, 10), err)
	}
}

// @Summary 删除评论
// @Description 删除评论
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgDelete true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /comment/del [post]
func delComment(c *mars.Context) {
	arg := new(model.ArgDelete)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.DelComment(c, arg.ID))

}

// @Summary 获取单个评论
// @Description 如果当前评论是子评论，则获取父级评论并将起所有子评论返回
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Success 200 "成功,返回comment_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /comment/get [get]
func getComment(c *mars.Context) {
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(srv.GetComment(c, id))
	}

}
