package usecase

import (
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
)

type CountryCodeUsecase struct {
	sqalx.Node
	*sqlx.DB
	CountryCodeRepository interface {
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.CountryCode, err error)
	}
}

func (p *CountryCodeUsecase) GetAll(ctx *biz.BizContext) (items []*models.CountryCode, err error) {
	items = make([]*models.CountryCode, 0)

	data, err := p.CountryCodeRepository.GetAll(p.Node)
	if err != nil {
		err = tracerr.Wrap(err)
	}

	copier.Copy(&items, &data)
	return
}
