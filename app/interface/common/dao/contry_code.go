package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/common/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getCountryCodesSQL = "SELECT a.id,a.name,a.cn_name,a.code,a.emoji,a.prefix,a.deleted,a.created_at,a.updated_at  FROM country_codes a WHERE a.deleted=0 ORDER BY a.id "
)

func (p *Dao) GetAllCountryCodes(c context.Context, node sqalx.Node) (items []*model.CountryCode, err error) {
	items = make([]*model.CountryCode, 0)

	if err = node.SelectContext(c, &items, _getCountryCodesSQL); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByEmail error(%+v)", err))
	}

	return
}
