package http

import (
	"strings"
	"valerian/app/interface/feedback/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取反馈类型
// @Description 获取反馈类型,需要传入type参数，accuse_people,accuse_content,feedback
// @Tags feedback
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param type query string true "accuse_people,accuse_content,feedback"
// @Success 200 {array}  app.interface.feedback.model.FeedbackTypeResp "反馈类型"
// @Failure 500 "服务器端错误"
// @Router /list/feedback_types [get]
func feedbackTypes(c *mars.Context) {
	fType := strings.ToLower(c.Request.Form.Get("type"))

	switch fType {
	case model.FeedbackTypeAccusePeople:
		break
	case model.FeedbackTypeAccuseContent:
		break
	case model.FeedbackTypeFeedback:
		break
	default:
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.FeedbackTypeList(c, fType))
}

// @Summary 提交反馈、投诉
// @Description 提交反馈、投诉
// @Tags feedback
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.feedback.model.ArgAddFeedback true "请求"
// @Success 200 "成功"
// @Failure 18 "话题不存在"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /feedback [post]
func addFeedback(c *mars.Context) {
	arg := new(model.ArgAddFeedback)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.AddFeedback(c, arg))

}
