package http

import (
	"strconv"
	"valerian/app/interface/dm/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取消息列表
// @Description 获取消息列表
// @Tags dm
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param type query string false "消息类型"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.dm.model.NotificationResp "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /dm/list [get]
func getActivites(c *mars.Context) {
	var (
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

	aType := params.Get("type")
	switch aType {
	case "all":
	case "topic":
	case "article":
		break
	default:
		aType = "all"
		// c.JSON(nil, ecode.RequestErr)
		// return
	}

	c.JSON(srv.GetUserMessagesPaged(c, aType, limit, offset))
}

// @Summary 设置为所有已读
// @Description 设置为所有已读
// @Tags dm
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /dm/mark_all_read [post]
func markAllRead(c *mars.Context) {
	c.JSON(nil, srv.MarkAllRead(c))
}

// @Summary 设置单条消息已读
// @Description 设置单条消息已读
// @Tags dm
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.dm.model.ArgMarkRead true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /dm/mark_read [post]
func markRead(c *mars.Context) {
	arg := new(model.ArgMarkRead)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.MarkRead(c, arg.ID))
}
