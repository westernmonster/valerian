package http

import (
	"strconv"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 最近发布
// @Description 最近发布
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Param type query string true "类型：all,article,revise,discussion"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.account.model.RecentPublishResp "个人资料"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/recent [get]
func recent(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	aType := params.Get("type")
	switch aType {
	case "all":
	case "revise":
	case "article":
	case "discussion":
		break
	default:
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetMemberRecentPubsPaged(c, id, aType, limit, offset))
}

// @Summary 获取用户资料
// @Description 获取用户资料
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Param include query string true  "目前支持"
// @Success 200 {object}  app.interface.account.model.MemberInfo "个人资料"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/info [get]
func memberInfo(c *mars.Context) {
	var id int64
	var err error
	params := c.Request.Form

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(srv.GetMemberInfo(c, id))
}

// @Summary 获取用户认证信息
// @Description 获取用户认证信息
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Success 200 {object}  app.interface.account.model.MemberCertInfo "认证资料"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/cert [get]
func memberCert(c *mars.Context) {
	var id int64
	var err error
	params := c.Request.Form

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(srv.GetMemberCert(c, id))
}

// @Summary 获取用户动态
// @Description 获取用户动态
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "用户ID"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.account.model.FeedResp "用户动态"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/activities [get]
func memberActivites(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetMemberActivitiesPaged(c, id, limit, offset))
}

// @Summary 获取用户管理话题
// @Description 获取用户管理话题
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "用户ID"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.account.model.MemberTopicResp "用户动态"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/managed_topics [get]
func memberManagedTopics(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetMemberManageTopicsPaged(c, id, limit, offset))
}

// @Summary 获取用户话题
// @Description 获取用户话题
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "用户ID"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.account.model.MemberTopicResp "用户动态"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/followed_topics [get]
func memberFollowedTopics(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetMemberFollowedTopicsPaged(c, id, limit, offset))
}

// @Summary 获取用户讨论
// @Description 获取用户讨论
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "用户ID"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.account.model.MemberDiscussResp "用户动态"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/discussions [get]
func memberDiscuss(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetMemberDiscussionsPaged(c, id, limit, offset))
}

// @Summary 获取用户文章
// @Description 获取用户文章
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "用户ID"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.account.model.MemberArticleResp "用户动态"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/articles [get]
func memberArticles(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetMemberArticlesPaged(c, id, limit, offset))
}

// @Summary 拉黑
// @Description 拉黑
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body app.interface.account.model.ArgBlockMember true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/block [post]
func block(c *mars.Context) {
	// TODO: block user
}
