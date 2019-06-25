package model

import "valerian/library/database/sqlx/types"

type Locale struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Locale    string        `db:"locale" json:"locale"`         // Locale 语言编码
	Name      string        `db:"name" json:"name"`             // Name 语言名称
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}
