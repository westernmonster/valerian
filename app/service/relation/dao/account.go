package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAccountByIDSQL = "SELECT a.id,a.mobile,a.user_name,a.email,a.password,a.role,a.salt,a.gender,a.birth_year,a.birth_month,a.birth_day,a.location,a.introduction,a.avatar,a.source,a.ip,a.id_cert,a.work_cert,a.is_org,a.is_vip,a.deleted,a.created_at,a.updated_at,a.prefix  FROM accounts a WHERE a.deleted=0 AND a.id=?"
)

func (p *Dao) GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error) {
	item = &model.Account{}
	if err = node.GetContext(c, item, _getAccountByIDSQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByID error(%+v), id(%d)", err, id))
	}

	return
}
