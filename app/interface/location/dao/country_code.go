package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/location/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getCountryCodesSQL = "SELECT a.* FROM country_codes a WHERE a.deleted=0 ORDER BY a.id "
)

func (p *Dao) GetAllCountryCodes(c context.Context, node sqalx.Node) (items []*model.CountryCode, err error) {
	items = make([]*model.CountryCode, 0)

	if err = node.SelectContext(c, &items, _getCountryCodesSQL); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByEmail error(%+v)", err))
	}

	return
}
