package model

import "valerian/library/database/sqlx/types"

type TopicFeed struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	TopicID    int64         `db:"topic_id" json:"topic_id,string"`       // TopicID 话题ID
	ActionType string        `db:"action_type" json:"action_type"`        // ActionType 行为类型
	ActionTime int64         `db:"action_time" json:"action_time,string"` // ActionTime 行为发起时间
	ActionText string        `db:"action_text" json:"action_text"`        // ActionText 行为文字内容
	ActorID    int64         `db:"actor_id" json:"actor_id,string"`       // ActorID 行为发起者
	ActorType  string        `db:"actor_type" json:"actor_type"`          // ActorType 行为发起者类型
	TargetID   int64         `db:"target_id" json:"target_id,string"`     // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`        // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}
