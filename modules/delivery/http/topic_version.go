package http

import (
	"strconv"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/models"
)

// @Summary 获取话题版本
// @Description 获取话题版本
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id path string true  "话题集合ID"
// @Success 200 {array} models.TopicVersion "话题成员"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_sets/{id}/versions [get]
func (p *TopicCtrl) GetTopicVersions(ctx *mars.Context) {
	idStr := ctx.Request.Form.Get("topic_id")
	topicSetID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.GetTopicVersions(ctx.Context, bizCtx, topicSetID)
	ctx.JSON(items, err)

}

// @Summary 新增话题版本
// @Description 新增话题版本
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id path string true  "话题集合ID"
// @Param req body models.ArgNewVersion true "请求"
// @Success 200 "成功,返回topic_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_sets/{id}/versions [post]
func (p *TopicCtrl) CreateNewVersion(ctx *mars.Context) {
	req := new(models.ArgNewVersion)

	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	id, err := p.TopicUsecase.CreateNewVersion(ctx.Context, bizCtx, req)
	ctx.JSON(strconv.FormatInt(id, 10), err)

	return

}
