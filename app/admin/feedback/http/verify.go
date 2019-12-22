package http

import (
	"strconv"
	"valerian/app/admin/feedback/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 审核反馈/举报
// @Description 审核反馈、举报
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.feedback.model.ArgVerifyFeedback true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/feedback/verify [post]
func verifyFeedback(c *mars.Context) {
	arg := new(model.ArgVerifyFeedback)
	if e := c.Bind(arg); e != nil {
		return
	}
	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, srv.VerifyFeedback(c, arg))
}

// @Summary 反馈列表
// @Description 反馈列表
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param target_type query string false "目标类型"
// @Param feedback_type query string false "反馈类型"
// @Param verify_status query string false "审核类型"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.admin.feedback.model.FeedbackListResp "反馈列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/feedback/list [get]
func listFeedback(c *mars.Context) {
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

	cond := map[string]interface{}{}
	targetType := params.Get("target_type")
	if len(targetType) > 0 {
		cond["target_type"] = targetType
	}

	feedback_type := params.Get("feedback_type")
	if len(feedback_type) > 0 {
		cond["feedback_type"] = feedback_type
	}

	verify_status := params.Get("verify_status")
	if len(verify_status) > 0 {
		cond["verify_status"] = verify_status
	}

	c.JSON(srv.GetFeedbacksByCondPaged(c, cond, limit, offset))
}

// @Summary 举报历史
// @Description 举报历史
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param created_by query string true "选中用户的 id"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.admin.feedback.model.FeedbackListResp "举报历史列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/feedback/report/list [get]
func reportHistory(c *mars.Context) {
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

	cond := map[string]interface{}{}
	createdBy := params.Get("created_by")
	if len(createdBy) > 0 {
		cond["created_by"] = createdBy
	}

	c.JSON(srv.GetReportByCondPaged(c, cond, limit, offset))
}

// @Summary 被举报历史
// @Description 被举报历史
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param target_user_id query string true "选中用户的 id"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.admin.feedback.model.FeedbackListResp "被举报历史列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/feedback/be-reported/list [get]
func beReportedHistory(c *mars.Context) {
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

	cond := map[string]interface{}{}
	targetUserId := params.Get("target_user_id")
	if len(targetUserId) > 0 {
		cond["target_user_id"] = targetUserId
	}

	c.JSON(srv.GetBeReportedByCondPaged(c, cond, limit, offset))
}
