package usecase

import (
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

type LocaleUsecase struct {
	sqalx.Node
	*sqlx.DB
	LocaleRepository interface {
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.Locale, err error)
	}
}

func (p *LocaleUsecase) GetAll(ctx *biz.BizContext) (items []*models.Locale, err error) {
	items = make([]*models.Locale, 0)

	data, err := p.LocaleRepository.GetAll(p.Node)
	if err != nil {
		err = tracerr.Wrap(err)
	}

	copier.Copy(&items, &data)
	return
}
