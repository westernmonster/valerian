package model

import "valerian/library/database/sqlx/types"

type RecentPub struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type RecentView struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
