package usecase

import (
	"context"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/modules/repo"

	"github.com/jinzhu/copier"
	"github.com/ztrue/tracerr"

	"valerian/models"
)

type CountryCodeUsecase struct {
	sqalx.Node
	*sqlx.DB
	CountryCodeRepository interface {
		// GetAll get all records
		GetAll(ctx context.Context, node sqalx.Node) (items []*repo.CountryCode, err error)
	}
}

func (p *CountryCodeUsecase) GetAll(ctx context.Context) (items []*models.CountryCode, err error) {
	items = make([]*models.CountryCode, 0)

	data, err := p.CountryCodeRepository.GetAll(ctx, p.Node)
	if err != nil {
		err = tracerr.Wrap(err)
	}

	copier.Copy(&items, &data)
	return
}
