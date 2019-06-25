package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getLocaleByCondition = "SELECT a.* FROM locales a WHERE a.deleted=0 %s"
)

func (p *Dao) GetLocaleByCondition(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Locale, err error) {
	item = new(model.Locale)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["locale"]; ok {
		clause += " AND a.locale =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf(_getLocaleByCondition, clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		log.For(c).Error(fmt.Sprintf("dao.GetLocaleByCondition error(%+v), condition(%+v)", err, cond))
	}

	return
}
