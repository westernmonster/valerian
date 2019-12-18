package http

import (
	"strconv"
	"valerian/app/admin/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 工作认证
// @Description  工作认证
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.account.model.ArgWorkCert true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/account/workcert [post]
func setWorkCert(c *mars.Context) {
	arg := new(model.ArgWorkCert)
	if e := c.Bind(arg); e != nil {
		return
	}
	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, srv.WorkCert(c, arg))
}

// @Summary 工作认证列表
// @Description 工作认证列表
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
// @Success 200 {object}  app.interface.comment.model.CommentListResp "评论列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/workcert/list [get]
func listWorkCert(c *mars.Context) {
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
	
	c.JSON(srv.GetWorkCertsByCondPaged(c, cond, limit, offset))
}
