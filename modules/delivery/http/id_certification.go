package http

import (
	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/library/cloudauth"
	"valerian/library/net/http/mars"
)

type IDCertificationCtrl struct {
	infrastructure.BaseCtrl

	IDCertificationUsecase interface {
		Request(ctx *biz.BizContext) (token cloudauth.VerifyTokenData, err error)
		GetStatus(ctx *biz.BizContext) (status int, err error)
	}
}

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
// @Router /me/certification/idcard [post]
func (p *IDCertificationCtrl) Request(ctx *mars.Context) {
	token, err := p.IDCertificationUsecase.Request(p.GetBizContext(ctx))
	ctx.JSON(token, err)
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
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/certification/idcard/status [get]
func (p *IDCertificationCtrl) GetStatus(ctx *mars.Context) {
	status, err := p.IDCertificationUsecase.GetStatus(p.GetBizContext(ctx))

	ctx.JSON(status, err)

	return
}
