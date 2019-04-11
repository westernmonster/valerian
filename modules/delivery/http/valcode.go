package http

import (
	"net/http"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/models"
	"github.com/gin-gonic/gin"
)

type ValcodeCtrl struct {
	infrastructure.BaseCtrl

	ValcodeUsecase interface {
		RequestEmailValcode(bizCtx *biz.BizContext, req *models.RequestEmailValcodeReq) (createdTime int64, err error)
		RequestMobileValcode(bizCtx *biz.BizContext, req *models.RequestMobileValcodeReq) (createdTime int64, err error)
	}
}

// @Summary 请求邮件验证码
// @Description 请求邮件验证码
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.RequestEmailValcodeReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 500 "服务器端错误"
// @Router /valcodes/email [post]
func (p *ValcodeCtrl) RequestEmailValcode(ctx *gin.Context) {
	req := new(models.RequestEmailValcodeReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	createdTime, err := p.ValcodeUsecase.RequestEmailValcode(p.GetBizContext(ctx), req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, createdTime)

	return
}

// @Summary 请求短信验证码
// @Description 请求短信验证码
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.RequestMobileValcodeReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 500 "服务器端错误"
// @Router /valcodes/mobile [post]
func (p *ValcodeCtrl) RequestMobileValcode(ctx *gin.Context) {
	req := new(models.RequestMobileValcodeReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	createdTime, err := p.ValcodeUsecase.RequestMobileValcode(p.GetBizContext(ctx), req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, createdTime)

	return
}
