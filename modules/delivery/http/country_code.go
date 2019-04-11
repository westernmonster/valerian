package http

import (
	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/models"
	"github.com/gin-gonic/gin"
)

type CountryCodeCtrl struct {
	infrastructure.BaseCtrl

	CountryCodeUsecase interface {
		GetAll(bizCtx *biz.BizContext) (items []*models.CountryCode, err error)
	}
}

// @Summary 获取电话国家区号
// @Description 获取电话国家区号
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {array} models.CountryCode "用户资料"
// @Failure 500 "服务器端错误"
// @Router /country_codes [get]
func (p *CountryCodeCtrl) GetAll(ctx *gin.Context) {

	items, err := p.CountryCodeUsecase.GetAll(p.GetBizContext(ctx))
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return
}
