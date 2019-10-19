package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/certification/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAccountByIDSQL = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.id=?"
	_updateAccountSQL  = "UPDATE accounts SET mobile=?,user_name=?,email=?,password=?,role=?,salt=?,gender=?,birth_year=?,birth_month=?,birth_day=?,location=?,introduction=?,avatar=?,source=?,ip=?,id_cert=?,work_cert=?,is_org=?,is_vip=?,updated_at=? WHERE id=? AND deleted=0"
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

func (p *Dao) UpdateAccount(c context.Context, node sqalx.Node, item *model.Account) (err error) {
	if _,
		err = node.ExecContext(c,
		_updateAccountSQL,
		item.Mobile,
		item.UserName,
		item.Email,
		item.Password,
		item.Role,
		item.Salt,
		item.Gender,
		item.BirthYear,
		item.BirthMonth,
		item.BirthDay,
		item.Location,
		item.Introduction,
		item.Avatar,
		item.Source,
		item.IP,
		item.IDCert,
		item.WorkCert,
		item.IsOrg,
		item.IsVIP,
		item.UpdatedAt,
		item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccount error(%+v), item(%+v)", err, item))
	}
	return
}
