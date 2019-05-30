package dao

import (
	"context"
	"database/sql"
	"valerian/app/interface/account/model"
	"valerian/library/database/sqalx"
)

const (
	_getAreaSQL = "SELECT a.id, a.name, a.code, a.type, a.parent, a.deleted, a.created_at, a.updated_at FROM areas a WHERE a.id=? AND a.deleted=0"
)

func (p *Dao) GetArea(ctx context.Context, node sqalx.Node, id int64) (item *model.Area, err error) {
	item = new(model.Area)
	if err = node.GetContext(ctx, item, _getAreaSQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
	}
	return
}
