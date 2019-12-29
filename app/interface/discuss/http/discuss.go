package http

import (
	"fmt"
	"strconv"
	"strings"

	"valerian/app/interface/discuss/model"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

// @Summary 获取话题讨论列表(按话题)
// @Description 获取话题讨论列表,category_id 不传则是全部，-1代表问答，其他值则为自定义分类
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "topic_id"
// @Param category_id query string false "category_id"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.discuss.model.DiscussListResp "讨论列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/list/by_topic [get]
func getDiscusstionsByTopic(c *mars.Context) {
	var (
		id         int64
		categoryID int64
		err        error
		offset     int
		limit      int
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

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	strCategoryID := strings.TrimSpace(params.Get("category_id"))
	if strCategoryID == "" {
		strCategoryID = "0"
	}
	if categoryID, err = strconv.ParseInt(strCategoryID, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetTopicDiscussionsPaged(c, id, categoryID, limit, offset))
}

// @Summary 获取话题讨论列表（按用户)
// @Description 获取话题讨论列表,category_id 不传则是全部，-1代表问答，其他值则为自定义分类
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param account_id query string true "account_id"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.discuss.model.DiscussListResp "讨论列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/list/by_account [get]
func getDiscusstionsByAccount(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("account_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetUserDiscussionsPaged(c, id, limit, offset))
}

// @Summary 新增话题讨论
// @Description 新增话题讨论
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.discuss.model.ArgAddDiscuss true "请求"
// @Success 200 "成功,返回discussion_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/add [post]
func addDiscussion(c *mars.Context) {
	arg := new(model.ArgAddDiscuss)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		log.For(c).Error(fmt.Sprintf("validation error(%+v)", e))
		c.JSON(nil, ecode.RequestErr)
		return
	}

	id, err := srv.AddDiscussion(c, arg)
	c.JSON(strconv.FormatInt(id, 10), err)
}

// @Summary 更新话题讨论
// @Description 更新话题讨论
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.discuss.model.ArgUpdateDiscuss true "请求"
// @Success 200 "成功,返回discussion_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/edit [post]
func updateDiscussion(c *mars.Context) {
	arg := new(model.ArgUpdateDiscuss)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		log.For(c).Error(fmt.Sprintf("validation error(%+v)", e))
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateDiscussion(c, arg))
}

// @Summary 删除话题讨论
// @Description 删除话题讨论
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.discuss.model.ArgDelete true "请求"
// @Success 200 "成功,返回discussion_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/del [post]
func delDiscussion(c *mars.Context) {
	arg := new(model.ArgDelete)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.DelDiscussion(c, arg.ID))
}

// @Summary 获取话题讨论详情
// @Description 获取话题讨论详情
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "id"
// @Success 200 {object}  app.interface.discuss.model.DiscussDetailResp "讨论列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/get [get]
func getDiscussion(c *mars.Context) {
	var id int64
	var err error
	params := c.Request.Form

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		log.For(c).Error(fmt.Sprintf("req error(%+v)", err))
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetDiscussion(c, id))
}
