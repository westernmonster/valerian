package http

import (
	"strconv"
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 批量更新关联话题
// @Description 批量更新关联话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgBatchSaveRelatedTopics true "请求"
// @Failure 18 "话题不存在"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/related [post]
func editTopicRelations(c *mars.Context) {
	arg := new(model.ArgBatchSaveRelatedTopics)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.BulkSaveRelations(c, arg))

}

// @Summary 获取关联话题列表
// @Description 获取关联话题列表
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Param include query string true  "目前支持：meta"
// @Success 200 {array} model.RelatedTopicResp "关联话题"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/list/related [get]
func relatedTopics(c *mars.Context) {
	include := c.Request.Form.Get("include")
	var (
		id  int64
		err error
	)

	params := c.Request.Form

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetAllRelatedTopicsWithMeta(c, id, include))
}
