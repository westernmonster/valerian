package http

import "valerian/library/net/http/mars"

// @Summary 获取实名认证Token
// @Description 获取实名认证Token
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {object} cloudauth.VerifyTokenData "Token"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/certification/id [post]
func idCertificationRequest(c *mars.Context) {
	c.JSON(srv.RequestIDCertification(c))
}

// @Summary 获取实名认证状态
// @Description 获取实名认证状态
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 "返回实名认证状态"
// @Failure 42 "尚未发起身份认证"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/certification/id/status [get]
func idCertificationStatus(c *mars.Context) {
	c.JSON(srv.GetIDCertificationStatus(c))
}

// @Summary 获取身份信息信息
// @Description 获取身份信息信息
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {object} model.IDCertResp "认证信息"
// @Failure 42 "尚未发起身份认证"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/certification/id [get]
func idCert(c *mars.Context) {
}
