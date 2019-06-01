package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAccountByEmailSQL  = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.email=?"
	_getAccountByMobileSQL = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.mobile=?"
	_getAccountByIDSQL     = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.id=?"
	_changePasswordSQL     = "UPDATE accounts SET password=?, salt =? WHERE id=?"
	_updateAccountSQL      = "UPDATE accounts SET mobile=?,user_name=?,email=?,password=?,role=?,salt=?,gender=?,birth_year=?,birth_month=?,birth_day=?,location=?,introduction=?,avatar=?,source=?,ip=?,id_cert=?,work_cert=?,is_org=?,is_vip=?,updated_at=? WHERE id=?"
)

func (p *Dao) GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error) {
	item = &model.Account{}
	if err = node.GetContext(c, item, _getAccountByEmailSQL, email); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByEmail error(%+v), email(%s)", err, email))
	}

	return
}

func (p *Dao) GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error) {
	item = &model.Account{}
	if err = node.GetContext(c, item, _getAccountByMobileSQL, mobile); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByMobile error(%+v), email(%s)", err, mobile))
	}

	return
}

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

func (p *Dao) SetPassword(c context.Context, node sqalx.Node, password, salt string, id int64) (err error) {
	if _, err = node.ExecContext(c, _changePasswordSQL, password, salt, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetPassword error(%+v), id(%d)", err, id))
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
