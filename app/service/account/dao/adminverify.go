package dao

import (
	"context"
	"fmt"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/log"
)

func (p *Dao) AccountSetLock(c context.Context, node sqalx.Node, accountID int64, isLock bool) (err error) {
	sqlUpdate := "UPDATE accounts SET is_lock=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, types.BitBool(isLock), accountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AccountSetLock err(%+v)", err))
		return
	}

	return
}
