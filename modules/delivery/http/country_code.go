package http

import (
	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
	"valerian/modules/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
)

type CountryCodeCtrl struct {
	infrastructure.BaseCtrl

	CountryCodeUsecase interface {
		GetAll(bizCtx *biz.BizContext) (items []*models.CountryCode, err error)
	}
}

func NewCountryCodeCtrl(db *sqlx.DB, node sqalx.Node) *CountryCodeCtrl {
	return &CountryCodeCtrl{
		CountryCodeUsecase: &usecase.CountryCodeUsecase{
			Node:                  node,
			DB:                    db,
			CountryCodeRepository: &repo.CountryCodeRepository{},
		},
	}
}

// @Summary 获取电话国家区号
// @Description 获取电话国家区号
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {array} models.CountryCode "国家区号"
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
