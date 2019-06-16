package http

import (
	"strconv"
	"valerian/app/interface/draft/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取用户所有草稿列表
// @Description 获取用户所有草稿列表
// @Tags draft
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param category_id query string true "分类ID"
// @Success 200 {array} model.DraftResp "草稿分类列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /draft/list [get]
func drafts(c *mars.Context) {
	var categoryID *int64

	if id, err := strconv.ParseInt(c.Request.Form.Get("category_id"), 10, 64); err != nil {
		categoryID = nil
	} else if id != 0 {
		categoryID = &id
	}

	c.JSON(srv.GetUserDrafts(c, categoryID))
}

// @Summary 添加草稿
// @Description 添加草稿
// @Tags draft
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgAddDraft true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /draft/add [post]
func addDraft(c *mars.Context) {
	arg := new(model.ArgAddDraft)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.AddDraft(c, arg))
}

// @Summary 更新草稿
// @Description 更新草稿
// @Tags draft
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgUpdateDraft true "请求"
// @Success 200 "成功"
// @Failure 43 "草稿分类不存在"
// @Failure 44 "不属于你"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /draft/edit [post]
func updateDraft(c *mars.Context) {
	arg := new(model.ArgUpdateDraft)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateDraft(c, arg))
}

// @Summary 删除草稿分类
// @Description 删除草稿分类
// @Tags draft
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id query string true "ID"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /draft/del [post]
func delDraft(c *mars.Context) {
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(nil, srv.DelDraft(c, id))
	}
}
