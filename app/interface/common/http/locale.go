package http

import "valerian/library/net/http/mars"

// @Summary 获取系统支持语言
// @Description 获取系统支持语言
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {array}  app.interface.common.model.Locale "国家区号"
// @Failure 500 "服务器端错误"
// @Router /list/locales [get]
func locales(c *mars.Context) {
	c.JSON(srv.GetAllLocales(c))
}
