package http

import (
	"strconv"
	"valerian/app/interface/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 更改密码
// @Description 更改密码
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgChangePassword true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/password [put]
func changePassword(c *mars.Context) {
	arg := new(model.ArgChangePassword)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	aid, _ := c.Get("aid")
	c.JSON(nil, srv.ChangePassword(c, aid.(int64), arg))
}

// @Summary 获取个人资料
// @Description 获取个人资料
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Param include query string true  "目前支持：settings"
// @Success 200 {object}  app.interface.account.model.Profile "个人资料"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/profile [get]
func getProfile(c *mars.Context) {
	aid, _ := c.Get("aid")
	c.JSON(srv.GetProfile(c, aid.(int64)))
}

// @Summary 更改用户资料
// @Description 更改用户资料
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgUpdateProfile true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/edit [put]
func updateProfile(c *mars.Context) {
	arg := new(model.ArgUpdateProfile)
	if e := c.Bind(arg); e != nil {
		return
	}

	aid, _ := c.Get("aid")
	c.JSON(nil, srv.UpdateProfile(c, aid.(int64), arg))
}

// @Summary 忘记密码
// @Description 忘记密码
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgForgetPassword true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/password/reset [put]
func forgetPassword(c *mars.Context) {
	arg := new(model.ArgForgetPassword)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.ForgetPassword(c, arg))
}

// @Summary 重设密码
// @Description 重设密码
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.account.model.ArgResetPassword true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/password/reset/confirm [put]
func resetPassword(c *mars.Context) {
	arg := new(model.ArgResetPassword)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.ResetPassword(c, arg))
}

// @Summary 获取当前用户话题
// @Description 获取当前用户话题
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param cate query string true "类型：managed, followed, faved, viewed"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object}  app.interface.account.model.MemberTopicResp "用户动态"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /account/list/my_topics [get]
func accountTopics(c *mars.Context) {
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
	c.JSON(srv.GetUserTopicsPaged(c, params.Get("cate"), int32(limit), int32(offset)))
}
