package http

import (
	"strconv"
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取成员列表
// @Description 获取成员列表
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Param page query integer true "页码"
// @Param page_size query integer true "每页大小"
// @Success 200 {object} model.TopicMembersPagedResp "话题成员"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/list/members [get]
func topicMembers(c *mars.Context) {
	var (
		id       int64
		err      error
		page     int
		pageSize int
	)

	params := c.Request.Form

	if page, err = strconv.Atoi(params.Get("page")); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if page <= 0 {
		page = 1
	}

	if pageSize, err = strconv.Atoi(params.Get("page_size")); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if pageSize <= 0 {
		pageSize = 10
	}

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetTopicMembersPaged(c, id, page, pageSize))
}

// @Summary 批量更新话题成员
// @Description 批量更新话题成员
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgBatchSavedTopicMember true "请求"
// @Success 200 "成功"
// @Failure 18 "话题不存在"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/members [post]
func editTopicMembers(c *mars.Context) {
	arg := new(model.ArgBatchSavedTopicMember)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.BulkSaveMembers(c, arg))

}

// @Summary 退出话题
// @Description 退出话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgLeave true "请求"
// @Success 200 "成功"
// @Failure 68 "主理人不可退出，只可转让后再退出"
// @Failure 34 "不是话题成员"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/leave [post]
func leave(c *mars.Context) {
	arg := new(model.ArgLeave)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Leave(c, arg.TopicID))
}
