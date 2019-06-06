package model

import "valerian/library/database/sqlx/types"

type CountryCode struct {
	ID        int64         `db:"id" json:"id,string"`    // ID ID
	Name      string        `db:"name" json:"name"`       // Name 国家名
	Emoji     string        `db:"emoji" json:"emoji"`     // Emoji 国旗
	CnName    string        `db:"cn_name" json:"cn_name"` // CnName 国家中文名
	Code      string        `db:"code" json:"code"`       // Code 编码
	Prefix    string        `db:"prefix" json:"prefix"`   // Prefix 前缀
	Deleted   types.BitBool `db:"deleted" json:"-"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"-"`    // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"-"`    // UpdatedAt 更新时间
}

type Area struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Name      string        `db:"name" json:"name"`             // Name 名称
	Code      string        `db:"code" json:"code"`             // Code 编码
	Type      string        `db:"type" json:"type"`             // Type 编码
	Parent    int64         `db:"parent" json:"parent,string"`  // Parent 父级ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}
