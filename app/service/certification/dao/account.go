package dao

import (
	"context"
	"fmt"
	"time"

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
