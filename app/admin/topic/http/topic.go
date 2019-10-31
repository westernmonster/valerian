package http

import (
	"strconv"
	"valerian/app/admin/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
)

// @Summary 新增话题
// @Description 新增话题
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.topic.model.ArgCreateTopic true "请求"
// @Success 200 "成功,返回topic_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/topic/add [post]
func createTopic(c *mars.Context) {
	arg := new(model.ArgCreateTopic)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if topicID, err := srv.CreateTopic(c, arg); err != nil {
		c.JSON(nil, err)
	} else {
		c.JSON(strconv.FormatInt(topicID, 10), err)
	}

}

// @Summary 更新话题
// @Description 更新话题
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.topic.model.ArgUpdateTopic true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/topic/edit [post]
func editTopic(c *mars.Context) {
	arg := new(model.ArgUpdateTopic)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateTopic(c, arg))
}

// @Summary 删除话题
// @Description 删除话题
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.topic.model.ArgDelete true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/topic/del [post]
func deleteTopic(c *mars.Context) {
	arg := new(model.ArgDelete)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.DelTopic(c, arg.ID))
}

// @Summary 获取话题
// @Description 获取话题
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Param include query string true  "目前支持：members,catalogs,auth_topics,catalogs,meta"
// @Success 200 {object}  app.admin.topic.model.TopicResp "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/topic/get [get]
func getTopic(c *mars.Context) {
	include := c.Request.Form.Get("include")
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(srv.GetTopic(c, id, include))
	}
}

// @Summary 更改主理人
// @Description 更改主理人,需要用户都为话题成员，并且发起操作用户必须为当前主理人
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.topic.model.ArgChangeOwner true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/topic/owner [post]
func changeOwner(c *mars.Context) {
	arg := new(model.ArgChangeOwner)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.ChangeOwner(c, arg))
}

// @Summary 有编辑权限的话题列表
// @Description 有编辑权限的话题列表
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param query query string true "查询条件"
// @Param pn query integer false "页码 1开始"
// @Param ps query integer false "每页大小"
// @Success 200 {object}  app.admin.topic.model.CanEditTopicsResp "成员"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/topic/list/has_edit_permission [get]
func topicsWithEditPermission(c *mars.Context) {
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

	c.JSON(srv.GetUserCanEditTopics(c, params.Get("query"), pn, ps))
}

// @Summary 已经加入话题列表
// @Description 已经加入话题列表
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param query query string true "查询条件"
// @Param ps query integer false "每页大小"
// @Param pn query integer false "页码 1开始"
// @Success 200 {object}  app.admin.topic.model.JoinedTopicsResp "成员"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/topic/list/followed [get]
func followedTopics(c *mars.Context) {
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

	c.JSON(srv.FollowedTopics(c, params.Get("query"), pn, ps))
}

// @Summary 获取话题Meta信息
// @Description 获取话题Meta信息
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Success 200 {object}  app.admin.topic.model.TopicMeta "Meta"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/topic/meta [get]
func topicMeta(c *mars.Context) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		c.JSON(nil, ecode.AcquireAccountIDFailed)
		return
	}
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(srv.GetTopicMeta(c, aid, id))
	}
}
