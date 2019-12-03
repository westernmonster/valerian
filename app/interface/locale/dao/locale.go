package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/locale/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getLocalesSQL = "SELECT a.id,a.locale,a.name,a.deleted,a.created_at,a.updated_at FROM locales a WHERE a.deleted=0 ORDER BY a.id "
)

func (p *Dao) GetAllLocales(c context.Context, node sqalx.Node) (items []*model.Locale, err error) {
	items = make([]*model.Locale, 0)

	if err = node.SelectContext(c, &items, _getLocalesSQL); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllLocales error(%+v)", err))
	}

	return
}
