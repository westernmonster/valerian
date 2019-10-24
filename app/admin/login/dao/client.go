package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/admin/login/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getClientSQL = `SELECT a.* FROM clients a WHERE a.deleted=0 AND a.client_Id=? `
)

func (p *Dao) GetClient(c context.Context, node sqalx.Node, clientID string) (item *model.Client, err error) {
	item = new(model.Client)

	if err = node.GetContext(c, item, _getClientSQL, clientID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		log.For(c).Error(fmt.Sprintf("dao.GetClient error(%+v), id(%s)", err, clientID))
	}

	return
}
