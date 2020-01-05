package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"valerian/app/service/certification/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_updateAccountIDCertSQL   = "UPDATE accounts SET id_cert=?,updated_at=? WHERE id=? AND deleted=0"
	_updateAccountWorkCertSQL = "UPDATE accounts SET work_cert=?,updated_at=? WHERE id=? AND deleted=0"
)

func (p *Dao) UpdateAccountIDCert(c context.Context, node sqalx.Node, aid int64, idCert bool) (err error) {
	idCertVal := 0
	if idCert {
		idCertVal = 1
	}
	if _, err = node.ExecContext(c, _updateAccountIDCertSQL, idCertVal, time.Now().Unix(), aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountIDCert error(%+v), id_cert(%+v) aid(%d)", err, idCert, aid))
	}
	return
}

func (p *Dao) UpdateAccountWorkCert(c context.Context, node sqalx.Node, aid int64, workCert bool) (err error) {
	workCertVal := 0
	if workCert {
		workCertVal = 1
	}
	if _, err = node.ExecContext(c, _updateAccountIDCertSQL, workCertVal, time.Now().Unix(), aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountWorkCert error(%+v), work_cert(%+v) aid(%d)", err, workCert, aid))
	}
	return
}

func (p *Dao) UpdateAccountRealName(c context.Context, node sqalx.Node, aid int64, realName string) (err error) {
	sqlUpdate := "UPDATE accounts SET user_name=?,updated_at=? WHERE id=? AND deleted=0"
	if _, err = node.ExecContext(c, sqlUpdate, realName, time.Now().Unix(), aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountRealName error(%+v), user_name(%+v) aid(%d)", err, realName, aid))
	}
	return
}

func (p *Dao) GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error) {
	item = &model.Account{}
	sqlSelect := "SELECT a.id,a.mobile,a.user_name,a.email,a.password,a.role,a.salt,a.gender,a.birth_year,a.birth_month,a.birth_day,a.location,a.introduction,a.avatar,a.source,a.ip,a.id_cert,a.work_cert,a.is_org,a.is_vip,a.deleted,a.created_at,a.updated_at,a.prefix,a.is_lock,a.deactive FROM accounts a WHERE a.id=? AND a.deleted=0"
	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByID error(%+v), id(%d)", err, id))
	}

	return
}
