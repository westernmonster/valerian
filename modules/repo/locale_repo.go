package repo

import (
	"context"
	"valerian/library/database/sqalx"

	packr "github.com/gobuffalo/packr"
	tracerr "github.com/ztrue/tracerr"
)

type Locale struct {
	ID        int64  `db:"id" json:"id,string"`          // ID ID
	Locale    string `db:"locale" json:"locale"`         // Locale 语言编码
	Name      string `db:"name" json:"name"`             // Name 语言名称
	Deleted   int    `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type LocaleRepository struct{}

// GetAll get all records
func (p *LocaleRepository) GetAll(ctx context.Context, node sqalx.Node) (items []*Locale, err error) {
	items = make([]*Locale, 0)
	sqlSelect := packr.NewBox("./sql/locale").String("GET_ALL.sql")

	err = node.SelectContext(ctx, &items, sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}
