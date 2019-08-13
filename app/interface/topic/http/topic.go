package http

import (
	"fmt"
	"strconv"
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 新增话题
// @Description 新增话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgCreateTopic true "请求"
// @Success 200 "成功,返回topic_id"
// @Failure 20 "获取用户ID失败"
// @Failure 30 "话题类型不存在"
// @Failure 31 "话题版本已经存在"
// @Failure 19 "主理人只允许一个"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/add [post]
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
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgUpdateTopic true "请求"
// @Success 200 "成功"
// @Failure 18 "话题不存在"
// @Failure 20 "获取用户ID失败"
// @Failure 30 "话题类型不存在"
// @Failure 31 "话题版本已经存在"
// @Failure 34 "你不是话题成员"
// @Failure 35 "你不是话题管理员"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/edit [post]
func editTopic(c *mars.Context) {
	arg := new(model.ArgUpdateTopic)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		fmt.Println(e)
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateTopic(c, arg))
}

// @Summary 删除话题
// @Description 删除话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/del [post]
func deleteTopic(c *mars.Context) {
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(nil, srv.DelTopic(c, id))
	}
}

// @Summary 获取话题
// @Description 获取话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Param include query string true  "目前支持：members,versions,related_topics,versions[*].catalogs,meta"
// @Success 200 {object} model.TopicResp "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/get [get]
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
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgChangeOwner true "请求"
// @Success 200 "成功"
// @Failure 20 "获取用户ID失败"
// @Failure 34 "不是话题成员"
// @Failure 46 "不是话题主理人"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/owner [post]
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

// @Summary 收藏话题
// @Description 收藏话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Success 200 "成功后返回bool值"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/fav [post]
func favTopic(c *mars.Context) {
}

// @Summary 举报话题
// @Description 举报话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgReportTopic true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/report [post]
func reportTopic(c *mars.Context) {
}
