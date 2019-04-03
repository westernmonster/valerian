package http

import (
	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/models"
	"github.com/gin-gonic/gin"
)

type CountryCodeCtrl struct {
	infrastructure.BaseCtrl

	CountryCodeUsecase interface {
		GetAll() (items []*models.CountryCode, err error)
	}
}

// GetAll 获取电话国家代码
// @Summary 获取电话国家代码
// @Description  获取电话国家代码
// @Tags common
// @Accept json
// @Produce json
// @Success 200 "成功"
// @Failure 500 "服务器端错误"
// @Router /country_codes [get]
func (p *CountryCodeCtrl) GetAll(ctx *gin.Context) {

	items, err := p.CountryCodeUsecase.GetAll()
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return
}
