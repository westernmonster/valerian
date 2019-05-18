package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/passport-login/model"
	"valerian/library/database/sqalx"
)

const (
	_getAccountByEmailSQL  = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.email=?"
	_getAccountByMobileSQL = "SELECT a.* FROM accounts a WHERE a.deleted=0 AND a.mobile=?"
)

func (p *Dao) GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error) {
	item = &model.Account{}
	if e := node.GetContext(c, item, _getAccountByEmailSQL, email); e != nil {
		if e == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		p.logger.For(c).Error(fmt.Sprintf("dao.GetAccountByEmail error(%+v), email(%s)", err, email))
	}

	return
}

func (p *Dao) GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error) {
	item = &model.Account{}
	if e := node.GetContext(c, item, _getAccountByMobileSQL, mobile); e != nil {
		if e == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		p.logger.For(c).Error(fmt.Sprintf("dao.GetAccountByMobile error(%+v), email(%s)", err, mobile))
	}

	return
}
