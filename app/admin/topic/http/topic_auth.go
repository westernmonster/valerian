package http

import (
	"strconv"
	"valerian/app/admin/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取授权话题列表
// @Description 获取授权话题列表
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Success 200 {array}  app.admin.topic.model.AuthTopicResp "授权话题"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/list/auth_topics [get]
func authTopics(c *mars.Context) {
	var (
		id  int64
		err error
	)

	params := c.Request.Form

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetAuthTopics(c, id))
}

// @Summary 批量更新授权话题
// @Description 批量更新授权话题
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.topic.model.ArgSaveAuthTopics true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/auth_topics [post]
func editAuthTopics(c *mars.Context) {
	arg := new(model.ArgSaveAuthTopics)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.SaveAuthTopics(c, arg))
}
