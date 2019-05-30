package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/passport-register/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAccountByEmailSQL  = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.email=?"
	_getAccountByMobileSQL = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.mobile=?"
	_getAccountByIDSQL     = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.id=?"

	_addAccountSQL = "INSERT INTO accounts( id,mobile,user_name,email,password,role,salt,gender,birth_year,birth_month,birth_day,location,introduction,avatar,source,ip,id_cert,work_cert,is_org,is_vip,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
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

func (p *Dao) AddAccount(c context.Context, node sqalx.Node, item *model.Account) (err error) {
	if _, err = node.ExecContext(c, _addAccountSQL,
		item.ID,
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
		item.Deleted,
		item.CreatedAt,
		item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccount error(%+v), item(%+v)", err, item))
	}

	return
}
