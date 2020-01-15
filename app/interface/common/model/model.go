package model

import "valerian/library/database/sqlx/types"

type Like struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type Dislike struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type Locale struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Locale    string        `db:"locale" json:"locale"`         // Locale 语言编码
	Name      string        `db:"name" json:"name"`             // Name 语言名称
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

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

type Fav struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
