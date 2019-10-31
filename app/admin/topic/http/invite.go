package http

import (
	"strconv"
	"valerian/app/admin/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 邀请
// @Description 邀请
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.topic.model.ArgTopicInvite true "请求"
// @Success 200 "成功"
// @Failure 18 "话题不存在"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/invite [post]
func inviteFans(c *mars.Context) {
	arg := new(model.ArgTopicInvite)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Invite(c, arg))
}

// @Summary 话题邀请处理
// @Description 话题邀请处理
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.topic.model.ArgProcessInvite true "请求"
// @Success 200  "返回关注状态"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/process_invite [post]
func processInvite(c *mars.Context) {
	arg := new(model.ArgProcessInvite)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.ProcessInvite(c, arg))
}

// @Summary 成员粉丝列表
// @Description 成员粉丝列表
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Param query query string false "查询条件"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.admin.topic.model.TopicMemberFansResp "粉丝列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/list/member_fans [get]
func memberFansList(c *mars.Context) {
	var (
		id    int64
		err   error
		pn    int
		ps    int
		query string
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

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	query = params.Get("query")

	c.JSON(srv.GetMemberFansList(c, id, query, pn, ps))
}
