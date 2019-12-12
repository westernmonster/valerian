package dao

import (
	"context"
	"fmt"
	"valerian/app/service/search/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Insert insert a new record
func (p *Dao) AddSearchStat(c context.Context, node sqalx.Node, item *model.SearchStat) (err error) {
	sqlInsert := "INSERT INTO search_stats( id,keywords,created_by,hits,enterpoint,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Keywords, item.CreatedBy, item.Hits, item.Enterpoint, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddSearchStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}