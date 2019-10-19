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

type TopicStat struct {
	TopicID         int64 `db:"topic_id" json:"topic_id,string"`          // TopicID 话题ID
	MemberCount     int   `db:"member_count" json:"member_count"`         // MemberCount 成员数
	ArticleCount    int   `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int   `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type ArticleStat struct {
	ArticleID    int64 `db:"article_id" json:"article_id,string"` // ArticleID 讨论ID
	LikeCount    int   `db:"like_count" json:"like_count"`        // LikeCount 喜欢数
	DislikeCount int   `db:"dislike_count" json:"dislike_count"`  // DislikeCount 反对数
	ReviseCount  int   `db:"revise_count" json:"revise_count"`    // ReviseCount 补充数
	CommentCount int   `db:"comment_count" json:"comment_count"`  // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type ReviseStat struct {
	ReviseID     int64 `db:"revise_id" json:"revise_id,string"`  // ReviseID 补充ID
	LikeCount    int   `db:"like_count" json:"like_count"`       // LikeCount 喜欢数
	DislikeCount int   `db:"dislike_count" json:"dislike_count"` // DislikeCount 反对数
	CommentCount int   `db:"comment_count" json:"comment_count"` // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`       // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`       // UpdatedAt 更新时间
}

type DiscussionStat struct {
	DiscussionID int64 `db:"discussion_id" json:"discussion_id,string"` // DiscussionID 讨论ID
	LikeCount    int   `db:"like_count" json:"like_count"`              // LikeCount 喜欢数
	DislikeCount int   `db:"dislike_count" json:"dislike_count"`        // DislikeCount 反对数
	CommentCount int   `db:"comment_count" json:"comment_count"`        // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`              // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`              // UpdatedAt 更新时间
}

type CommentStat struct {
	CommentID     int64 `db:"comment_id" json:"comment_id,string"`  // CommentID 账户ID
	LikeCount     int   `db:"like_count" json:"like_count"`         // LikeCount 喜欢数
	DislikeCount  int   `db:"dislike_count" json:"dislike_count"`   // DislikeCount 反对数
	ChildrenCount int   `db:"children_count" json:"children_count"` // ChildrenCount 子评论数
	CreatedAt     int64 `db:"created_at" json:"created_at"`         // CreatedAt 创建时间
	UpdatedAt     int64 `db:"updated_at" json:"updated_at"`         // UpdatedAt 更新时间
}
