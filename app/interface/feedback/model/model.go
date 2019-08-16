package model

import "valerian/library/database/sqlx/types"

type Feedback struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType int           `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Type       int           `db:"type" json:"type"`                    // Type 举报类型
	Desc       string        `db:"desc" json:"desc"`                    // Desc 备注
	CreatedBy  int64         `db:"created_by" json:"created_by,string"` // CreatedBy 举报人
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type FeedbackType struct {
	ID        int           `db:"id" json:"id"`                 // ID ID
	Type      string        `db:"type" json:"type"`             // Type 目标类型
	Name      string        `db:"name" json:"name"`             // Name 备注
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}
