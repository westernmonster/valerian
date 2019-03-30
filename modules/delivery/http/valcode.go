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

// Request 请求验证码
// @Summary 请求验证码
// @Description  请求验证码
// @Tags common
// @Accept json
// @Produce json
// @Param req body models.RequestValcodeReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /valcodes [post]
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
