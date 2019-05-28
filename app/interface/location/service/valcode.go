package service

import (
	"context"
	"valerian/app/interface/location/model"

	"github.com/jinzhu/copier"
)

func (p *Service) GetAllCountryCodes(ctx context.Context) (items []*model.CountryCodeResp, err error) {

	var (
		data []*model.CountryCode
	)

	items = make([]*model.CountryCodeResp, 0)

	if data, err = p.d.CountryCodesCache(ctx); err == nil {
		if err = copier.Copy(&items, &data); err != nil {
			return
		}
		return
	}

	if data, err = p.d.GetAllCountryCodes(ctx, p.d.DB()); err != nil {
		return
	}

	err = copier.Copy(&items, &data)
	return
}
