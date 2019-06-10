package service

import (
	"context"
	"valerian/app/interface/location/model"
)

func (p *Service) GetAllCountryCodes(ctx context.Context) (items []*model.CountryCode, err error) {
	var addCache = true

	items = make([]*model.CountryCode, 0)
	if items, err = p.d.CountryCodesCache(ctx); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetAllCountryCodes(ctx, p.d.DB()); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetCountryCodesCache(context.TODO(), items)
		})
	}
	return
}
