package repo

import (
	"context"

	"valerian/library/database/sqalx"

	tracerr "github.com/ztrue/tracerr"
)

type CountryCode struct {
	ID        int64  `db:"id" json:"id,string"`          // ID ID
	Name      string `db:"name" json:"name"`             // Name 国家名
	Emoji     string `db:"emoji" json:"emoji"`           // Emoji 国旗
	CnName    string `db:"cn_name" json:"cn_name"`       // CnName 国家中文名
	Code      string `db:"code" json:"code"`             // Code 编码
	Prefix    string `db:"prefix" json:"prefix"`         // Prefix 前缀
	Deleted   int    `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type CountryCodeRepository struct{}

// GetAll get all records
func (p *CountryCodeRepository) GetAll(ctx context.Context, node sqalx.Node) (items []*CountryCode, err error) {
	items = make([]*CountryCode, 0)
	sqlSelect := "SELECT a.* FROM country_codes a WHERE a.deleted=0 ORDER BY a.id "

	err = node.SelectContext(ctx, &items, sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}
