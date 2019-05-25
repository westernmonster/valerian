package http

import (
	"strconv"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/models"
)

// @Summary 批量更新话题成员
// @Description 批量更新话题成员
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.BatchSavedTopicMemberReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/{id}/members [patch]
func (p *TopicCtrl) BatchSavedTopicMember(ctx *mars.Context) {
	req := new(models.BatchSavedTopicMemberReq)

	ctx.Bind(req)
	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	idStr := ctx.Request.Form.Get("topic_id")
	topicID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	err = p.TopicUsecase.BulkSaveMembers(ctx.Context, bizCtx, topicID, req)
	ctx.JSON(nil, err)

}
