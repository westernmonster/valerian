package model

import "valerian/library/database/sqlx/types"

type Comment struct {
	ID         int64         `db:"id" json:"id,string"`                       // ID ID
	Content    string        `db:"content" json:"content"`                    // Content 内容
	TargetType string        `db:"target_type" json:"target_type"`            // TargetType 目标类型
	OwnerID    int64         `db:"owner_id" json:"owner_id,string"`           // OwnerID 资源ID (discussion, article, revise)
	ResourceID int64         `db:"resource_id" json:"resource_id,string"`     // ResourceID 所属对象ID (discussion, article, revise, comment)
	Featured   types.BitBool `db:"featured" json:"featured"`                  // Featured 是否精选
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                    // Deleted 是否删除
	ReplyTo    *int64        `db:"reply_to" json:"reply_to,omitempty,string"` // ReplyTo 被回复人
	CreatedBy  int64         `db:"created_by" json:"created_by,string"`       // CreatedBy 创建人
	CreatedAt  int64         `db:"created_at" json:"created_at"`              // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`              // UpdatedAt 更新时间
}

type CommentStat struct {
	CommentID     int64 `db:"comment_id" json:"comment_id,string"`  // CommentID 账户ID
	LikeCount     int   `db:"like_count" json:"like_count"`         // LikeCount 喜欢数
	DislikeCount  int   `db:"dislike_count" json:"dislike_count"`   // DislikeCount 反对数
	ChildrenCount int   `db:"children_count" json:"children_count"` // ChildrenCount 子评论数
	CreatedAt     int64 `db:"created_at" json:"created_at"`         // CreatedAt 创建时间
	UpdatedAt     int64 `db:"updated_at" json:"updated_at"`         // UpdatedAt 更新时间
}
