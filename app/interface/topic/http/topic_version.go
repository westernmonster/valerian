package http

import (
	"strconv"
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取话题版本列表
// @Description 获取话题版本列表
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_set_id query string true "话题集合ID"
// @Success 200 {array} model.TopicVersionResp "话题版本"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/list/versions [get]
func topicVersions(c *mars.Context) {
	var (
		topicSetID int64
		err        error
	)

	params := c.Request.Form
	if topicSetID, err = strconv.ParseInt(params.Get("topic_set_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetTopicVersions(c, topicSetID))
}

// @Summary 新建话题版本
// @Description 新建话题版本
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgNewTopicVersion true "请求"
// @Failure 20 "获取用户ID失败，一般是因为未登录造成"
// @Failure 18 "话题不存在"
// @Failure 34 "不属于话题成员"
// @Failure 35 "不是话题主理人或管理员"
// @Failure 31 "话题名已被占用"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/versions [post]
func addTopicVersion(c *mars.Context) {
	arg := new(model.ArgNewTopicVersion)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.AddTopicVersion(c, arg))
}

// @Summary 合并话题版本
// @Description 合并话题版本，需要当前用户为合并两个话题集中，是所有话题的成员并且是管理员或主理人
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgMergeTopicVersion true "请求"
// @Failure 20 "获取用户ID失败，一般是因为未登录造成"
// @Failure 18 "话题不存在"
// @Failure 34 "不属于话题成员"
// @Failure 35 "不是话题主理人或管理员"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/versions/merge [post]
func mergeTopicVersion(c *mars.Context) {
	arg := new(model.ArgMergeTopicVersion)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.MergeTopicVersions(c, arg))
}

// @Summary 保存话题版本（排序和重命名）
// @Description 保存话题版本（排序和重命名）
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body ArgSaveTopicVersions true "请求"
// @Failure 20 "获取用户ID失败，一般是因为未登录造成"
// @Failure 18 "话题不存在"
// @Failure 34 "不属于话题成员"
// @Failure 35 "不是话题主理人或管理员"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/versions/save [post]
func saveTopicVersions(c *mars.Context) {
	arg := new(model.ArgSaveTopicVersions)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.SaveTopicVersions(c, arg))
}
