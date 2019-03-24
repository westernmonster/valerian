package http

import (
	"net/http"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/models"
	"github.com/gin-gonic/gin"
)

type ValcodeCtrl struct {
	infrastructure.BaseCtrl

	ValcodeUsecase interface {
		Request(req *models.RequestValcodeReq) (createdTime int64, err error)
	}
}

func (p *ValcodeCtrl) Request(ctx *gin.Context) {
	req := new(models.RequestValcodeReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "验证失败，请检查您的输入",
		})

		return
	}

	createdTime, err := p.ValcodeUsecase.Request(req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, createdTime)

	return
}
