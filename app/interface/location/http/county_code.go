package http

import "valerian/library/net/http/mars"

// @Summary 获取电话国家区号
// @Description 获取电话国家区号
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {array}  app.interface.location.model.CountryCodeResp "国家区号"
// @Failure 500 "服务器端错误"
// @Router /list/country_codes [get]
func countryCodes(c *mars.Context) {
	c.JSON(srv.GetAllCountryCodes(c))
}
