package usecase

import (
	"context"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"

	"github.com/jinzhu/copier"
	"github.com/ztrue/tracerr"

	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
)

type LocaleUsecase struct {
	sqalx.Node
	*sqlx.DB
	LocaleRepository interface {
		// GetAll get all records
		GetAll(ctx context.Context, node sqalx.Node) (items []*repo.Locale, err error)
	}
}

func (p *LocaleUsecase) GetAll(c context.Context, ctx *biz.BizContext) (items []*models.Locale, err error) {
	items = make([]*models.Locale, 0)

	data, err := p.LocaleRepository.GetAll(c, p.Node)
	if err != nil {
		err = tracerr.Wrap(err)
	}

	copier.Copy(&items, &data)
	return
}
