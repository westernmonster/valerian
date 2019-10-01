package http

import (
	"strconv"

	"valerian/app/interface/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取粉丝列表
// @Description 获取粉丝列表
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param account_id query string true "用户ID"
// @Param query query string true "查询条件"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object} model.MemberResp "成员"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/fans [get]
func fans(c *mars.Context) {
	var (
		id     int64
		err    error
		offset int
		limit  int
		query  string
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

	if id, err = strconv.ParseInt(params.Get("account_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	query = params.Get("query")

	c.JSON(srv.FansPaged(c, id, query, limit, offset))
}

// @Summary 获取关注用户列表
// @Description 获取关注用户列表
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param account_id query string true "用户ID"
// @Param query query string true "查询条件"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object} model.MemberResp "成员"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/followings [get]
func followed(c *mars.Context) {
	var (
		id     int64
		err    error
		offset int
		limit  int
		query  string
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

	if id, err = strconv.ParseInt(params.Get("account_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	query = params.Get("query")

	c.JSON(srv.FollowPaged(c, id, query, limit, offset))
}

// @Summary 关注用户
// @Description 关注用户
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgFollow true "请求"
// @Success 200 "成功revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/follow [post]
func follow(c *mars.Context) {
	arg := new(model.ArgFollow)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Follow(c, arg))
}

// @Summary 取关用户
// @Description 取关用户
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgUnfollow true "请求"
// @Success 200 "成功revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/unfollow [post]
func unfollow(c *mars.Context) {
	arg := new(model.ArgUnfollow)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Unfollow(c, arg))
}
