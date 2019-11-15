package http

import (
	"valerian/library/net/http/mars"
)

// @Summary 冷启动(获取大分类)
// @Description 冷启动(获取大分类)
// @Tags init
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {object}  app.interface.init.model.MajorListResp "数据"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /init/list/major [get]
func getMajorData(c *mars.Context) {
}

// @Summary 冷启动(获取关联话题)
// @Description 冷启动(获取关联话题)
// @Tags init
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param major_ids query string true  "话题ID,逗号分隔"
// @Success 200 {object}  app.interface.init.model.RelatedListResp "数据"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /init/list/related [get]
func getRelatedData(c *mars.Context) {
}

// @Summary 冷启动(获取成员)
// @Description 冷启动(获取成员)
// @Tags init
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param major_ids query string true  "大分类话题ID,逗号分隔"
// @Param related_ids query string true  "关联话题ID,逗号分隔"
// @Success 200 {object}  app.interface.init.model.MemberListResp "数据"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /init/list/members [get]
func getMembersData(c *mars.Context) {
}
