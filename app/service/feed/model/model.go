package model

import "valerian/library/database/sqlx/types"

type Feed struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"`   // AccountID 账户ID
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

type AccountSetting struct {
	AccountID            int64         `db:"account_id" json:"account_id,string"`                  // AccountID 账户ID
	ActivityLike         types.BitBool `db:"activity_like" json:"activity_like"`                   // ActivityLike 动态-赞
	ActivityComment      types.BitBool `db:"activity_comment" json:"activity_comment"`             // ActivityComment 动态-评论
	ActivityFollowTopic  types.BitBool `db:"activity_follow_topic" json:"activity_follow_topic"`   // ActivityFollowTopic 动态-关注话题
	ActivityFollowMember types.BitBool `db:"activity_follow_member" json:"activity_follow_member"` // ActivityFollowMember 动态-关注成员
	NotifyLike           types.BitBool `db:"notify_like" json:"notify_like"`                       // NotifyLike 通知-赞
	NotifyComment        types.BitBool `db:"notify_comment" json:"notify_comment"`                 // NotifyComment 通知-评论
	NotifyNewFans        types.BitBool `db:"notify_new_fans" json:"notify_new_fans"`               // NotifyNewFans 通知-新粉丝
	NotifyNewMember      types.BitBool `db:"notify_new_member" json:"notify_new_member"`           // NotifyNewMember 通知-新成员
	Language             string        `db:"language" json:"language"`                             // Language 语言
	CreatedAt            int64         `db:"created_at" json:"created_at"`                         // CreatedAt 创建时间
	UpdatedAt            int64         `db:"updated_at" json:"updated_at"`                         // UpdatedAt 更新时间
}
