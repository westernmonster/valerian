package http

import (
	"context"
	"valerian/infrastructure"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/models"
	"valerian/modules/repo"
	"valerian/modules/usecase"

	"valerian/library/database/sqalx"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CountryCodeCtrl struct {
	infrastructure.BaseCtrl
	tracer opentracing.Tracer
	logger log.Factory

	CountryCodeUsecase interface {
		GetAll(ctx context.Context) (items []*models.CountryCode, err error)
	}
}

func NewCountryCodeCtrl(node sqalx.Node, logger log.Factory) *CountryCodeCtrl {
	return &CountryCodeCtrl{
		logger: logger,
		CountryCodeUsecase: &usecase.CountryCodeUsecase{
			Node:                  node,
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
func (p *CountryCodeCtrl) GetAll(ctx *mars.Context) {
	p.logger.For(ctx.Request.Context()).Info("HTTP", zap.String("method", ctx.Request.Method), zap.Stringer("url", ctx.Request.URL))

	items, err := p.CountryCodeUsecase.GetAll(ctx.Request.Context())

	ctx.JSON(items, err)

}
