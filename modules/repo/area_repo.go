package repo

import (
	"context"
	"database/sql"
	"valerian/library/database/sqalx"

	tracerr "github.com/ztrue/tracerr"
)

type Area struct {
	ID        int64  `db:"id" json:"id,string"`          // ID ID
	Name      string `db:"name" json:"name"`             // Name 名称
	Code      string `db:"code" json:"code"`             // Code 编码
	Type      string `db:"type" json:"type"`             // Type 编码
	Parent    int64  `db:"parent" json:"parent,string"`  // Parent 父级ID
	Deleted   int    `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type AreaRepository struct{}

// GetByID get record by ID
func (p *AreaRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *Area, exist bool, err error) {
	item = new(Area)
	sqlSelect := `SELECT a.id, a.name, a.code, a.type, a.parent, a.deleted, a.created_at, a.updated_at FROM areas a WHERE a.id=? AND a.deleted=0 `
	if e := node.GetContext(ctx, item, sqlSelect, id); e != nil {
		if e == sql.ErrNoRows {
			item = nil
			return
		}
		err = tracerr.Wrap(e)
		return
	}

	exist = true
	return
}
