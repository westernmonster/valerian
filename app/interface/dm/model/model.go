package model

import "valerian/library/database/sqlx/types"

type Message struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"`   // AccountID 账户ID
	ActionType string        `db:"action_type" json:"action_type"`        // ActionType 行为类型
	ActionTime int64         `db:"action_time" json:"action_time,string"` // ActionTime 行为发起时间
	ActionText string        `db:"action_text" json:"action_text"`        // ActionText 行为文字内容
	Actors     string        `db:"actors" json:"actors"`                  // Actors 行为发起者
	Extend     string        `db:"extend" json:"extend"`                  // Extend 扩展内容
	MergeCount int           `db:"merge_count" json:"merge_count"`        // MergeCount 合并条数
	ActorType  string        `db:"actor_type" json:"actor_type"`          // ActorType 行为发起者类型
	TargetID   int64         `db:"target_id" json:"target_id,string"`     // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`        // TargetType 目标类型
	IsRead     types.BitBool `db:"is_read" json:"is_read"`                // IsRead 是否阅读
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}

type MessageStat struct {
	AccountID   int64 `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	UnreadCount int   `db:"unread_count" json:"unread_count"`    // UnreadCount 未读数
	CreatedAt   int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type TopicMember struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"`     // TopicID 分类ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 成员ID
	Role      string        `db:"role" json:"role"`                    // Role 成员角色
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type TopicInviteRequest struct {
	ID            int64         `db:"id" json:"id,string"`                           // ID ID
	AccountID     int64         `db:"account_id" json:"account_id,string"`           // AccountID 用户ID
	TopicID       int64         `db:"topic_id" json:"topic_id,string"`               // TopicID 话题ID
	Status        int           `db:"status" json:"status"`                          // Status 状态
	Deleted       types.BitBool `db:"deleted" json:"deleted"`                        // Deleted 是否删除
	CreatedAt     int64         `db:"created_at" json:"created_at"`                  // CreatedAt 创建时间
	UpdatedAt     int64         `db:"updated_at" json:"updated_at"`                  // UpdatedAt 更新时间
	FromAccountID int64         `db:"from_account_id" json:"from_account_id,string"` // FromAccountID 邀请人
}

type TopicFollowRequest struct {
	ID            int64         `db:"id" json:"id,string"`                    // ID ID
	AccountID     int64         `db:"account_id" json:"account_id,string"`    // AccountID 用户ID
	TopicID       int64         `db:"topic_id" json:"topic_id,string"`        // TopicID 话题ID
	Status        int32         `db:"status" json:"status"`                   // Status 状态
	Deleted       types.BitBool `db:"deleted" json:"deleted"`                 // Deleted 是否删除
	CreatedAt     int64         `db:"created_at" json:"created_at"`           // CreatedAt 创建时间
	UpdatedAt     int64         `db:"updated_at" json:"updated_at"`           // UpdatedAt 更新时间
	Reason        string        `db:"reason" json:"reason"`                   // Reason 原因
	AllowViewCert types.BitBool `db:"allow_view_cert" json:"allow_view_cert"` // AllowViewCert 允许查看认证
	RejectReason  string        `db:"reject_reason" json:"reject_reason"`     // RejectReason 原因
}
