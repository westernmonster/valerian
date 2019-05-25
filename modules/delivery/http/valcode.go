package http

import (
	"context"

	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/email"
	"valerian/library/net/http/mars"
	"valerian/library/sms"

	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
	"valerian/modules/usecase"
)

type ValcodeCtrl struct {
	infrastructure.BaseCtrl

	ValcodeUsecase interface {
		RequestEmailValcode(c context.Context, ctx *biz.BizContext, req *models.RequestEmailValcodeReq) (createdTime int64, err error)
		RequestMobileValcode(c context.Context, ctx *biz.BizContext, req *models.RequestMobileValcodeReq) (createdTime int64, err error)
	}
}

func NewValcodeCtrl(smsClient *sms.SMSClient, emailClient *email.EmailClient, node sqalx.Node) *ValcodeCtrl {
	return &ValcodeCtrl{
		ValcodeUsecase: &usecase.ValcodeUsecase{
			Node:              node,
			SMSClient:         smsClient,
			EmailClient:       emailClient,
			ValcodeRepository: &repo.ValcodeRepository{},
		},
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
func (p *ValcodeCtrl) RequestEmailValcode(ctx *mars.Context) {
	req := new(models.RequestEmailValcodeReq)

	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	createdTime, err := p.ValcodeUsecase.RequestEmailValcode(ctx.Context, p.GetBizContext(ctx), req)
	ctx.JSON(createdTime, err)
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
func (p *ValcodeCtrl) RequestMobileValcode(ctx *mars.Context) {
	req := new(models.RequestMobileValcodeReq)

	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	createdTime, err := p.ValcodeUsecase.RequestMobileValcode(ctx.Context, p.GetBizContext(ctx), req)
	ctx.JSON(createdTime, err)
}
