package http

import (
	"strconv"
	"valerian/app/interface/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 审核工作认证
// @Description  审核工作认证
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.certification.model.ArgAuditWorkCert true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/certification/workcert/audit [post]
func auditWorkCert(c *mars.Context) {
	arg := new(model.ArgAuditWorkCert)
	if e := c.Bind(arg); e != nil {
		return
	}
	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, srv.AuditWorkCert(c, arg))
}

// @Summary 工作认证列表
// @Description 工作认证列表
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param status query string false "状态"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.certification.model.WorkCertsPagedResp "工作认证列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/certification/workcert/list [get]
func listWorkCert(c *mars.Context) {
	var (
		err    error
		offset int
		limit  int
		status int
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

	if status, err = strconv.Atoi(params.Get("status")); err != nil {
		status = 0
	} else if limit < 0 {
		status = 0
	}

	c.JSON(srv.GetWorkCertsPaged(c, status, limit, offset))
}

// @Summary 工作认证审核列表
// @Description 工作认证审核列表
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param account_id query string true "用户 id"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.certification.model.WorkCertHistoriesResp "工作认证审核列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/certification/workcert/history/list [get]
func workCertHistory(c *mars.Context) {
	var (
		err    error
		offset int
		limit  int
		id     int64
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

	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetWorkCertHistoriesPaged(c, id, limit, offset))
}
