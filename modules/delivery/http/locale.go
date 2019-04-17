package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
	"git.flywk.com/flywiki/api/modules/usecase"
)

type LocaleCtrl struct {
	infrastructure.BaseCtrl

	LocaleUsecase interface {
		GetAll(bizCtx *biz.BizContext) (items []*models.Locale, err error)
	}
}

func NewLocaleCtrl(db *sqlx.DB, node sqalx.Node) *LocaleCtrl {
	return &LocaleCtrl{
		LocaleUsecase: &usecase.LocaleUsecase{
			Node:             node,
			DB:               db,
			LocaleRepository: &repo.LocaleRepository{},
		},
	}

}

// @Summary 语言编码
// @Description 语言编码
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {array} models.Locale "语言编码"
// @Failure 500 "服务器端错误"
// @Router /locales [get]
func (p *LocaleCtrl) GetAll(ctx *gin.Context) {

	items, err := p.LocaleUsecase.GetAll(p.GetBizContext(ctx))
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return
}
