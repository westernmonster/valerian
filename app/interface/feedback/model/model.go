package model

import "valerian/library/database/sqlx/types"

type Feedback struct {
	ID           int64         `db:"id" json:"id,string"`                 // ID ID
	TargetID     int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType   int32         `db:"target_type" json:"target_type"`      // TargetType 目标类型
	TargetDesc   string        `db:"target_desc" json:"target_desc"`      // TargetDesc 目标描述（用于通知）
	FeedbackType int32         `db:"feedback_type" json:"feedback_type"`  // FeedbackType 举报类型
	FeedbackDesc string        `db:"feedback_desc" json:"feedback_desc"`  // FeedbackDesc 备注
	CreatedBy    int64         `db:"created_by" json:"created_by,string"` // CreatedBy 举报人
	Deleted      types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
	VerifyStatus int32         `db:"verify_status" json:"verify_status"`  // CheckStatus 审核状态
	VerifyDesc   string        `db:"verify_desc" json:"verify_desc"`      // CheckDesc 审核结果/原因
}

type FeedbackType struct {
	ID        int32         `db:"id" json:"id"`                 // ID ID
	Type      string        `db:"type" json:"type"`             // Type 目标类型
	Name      string        `db:"name" json:"name"`             // Name 备注
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}
