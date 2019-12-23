package dao

import (
	"context"
	"fmt"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) GetFansIDs(c context.Context, node sqalx.Node, aid int64) (items []int64, err error) {
	// items: = make([]*model.AccountFans, 0)
	items = make([]int64, 0)
	sqlSelect := "SELECT a.target_account_id FROM account_fans a WHERE a.deleted=0 AND a.account_id =?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFansIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetFansIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}
