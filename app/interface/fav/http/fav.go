package http

import (
	"valerian/app/interface/fav/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 收藏列表
// @Description  收藏列表
// @Tags fav
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param target_type query string true "类型：topic,article,revise,discussion"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object} model.FavListResp "讨论列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /fav/list/all [get]
func favList(c *mars.Context) {
}

// @Summary 收藏
// @Description 收藏
// @Tags fav
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgAddFav true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /fav/fav [post]
func fav(c *mars.Context) {
	arg := new(model.ArgAddFav)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Fav(c, arg))
}

// @Summary 取消收藏
// @Description 取消收藏
// @Tags fav
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgDelFav true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /fav/unfav [post]
func unfav(c *mars.Context) {
	arg := new(model.ArgDelFav)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Unfav(c, arg))
}
