package http

import (
	"github.com/gin-gonic/gin"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/infrastructure/cloudauth"
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
func (p *IDCertificationCtrl) Request(ctx *gin.Context) {
	token, err := p.IDCertificationUsecase.Request(p.GetBizContext(ctx))
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, token)

	return
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
func (p *IDCertificationCtrl) GetStatus(ctx *gin.Context) {
	status, err := p.IDCertificationUsecase.GetStatus(p.GetBizContext(ctx))
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, status)

	return
}
