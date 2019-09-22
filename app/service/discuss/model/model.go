package model

import "valerian/library/database/sqlx/types"

type DiscussCategory struct {
	ID        int64         `db:"id" json:"id,string"`             // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"` // TopicID 话题ID
	Seq       int           `db:"seq" json:"seq"`                  // Seq 顺序
	Name      string        `db:"name" json:"name"`                // Name 话题名
	Deleted   types.BitBool `db:"deleted" json:"deleted"`          // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`    // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`    // UpdatedAt 更新时间
}

type Discussion struct {
	ID          int64         `db:"id" json:"id,string"`                   // ID ID
	TopicID     int64         `db:"topic_id" json:"topic_id,string"`       // TopicID 话题ID
	CategoryID  int64         `db:"category_id" json:"category_id,string"` // CategoryID 分类ID
	CreatedBy   int64         `db:"created_by" json:"created_by,string"`   // CreatedBy 创建用户ID
	Title       *string       `db:"title" json:"title,omitempty"`          // Title 标题
	Content     string        `db:"content" json:"content"`                // Content 内容HTML
	ContentText string        `db:"content_text" json:"content_text"`      // ContentText 内容纯文本
	Deleted     types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}

type AccountResStat struct {
	AccountID       int64 `db:"account_id" json:"account_id,string"`      // AccountID 用户ID
	TopicCount      int   `db:"topic_count" json:"topic_count"`           // TopicCount 话题数
	ArticleCount    int   `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int   `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type TopicResStat struct {
	TopicID         int64 `db:"topic_id" json:"topic_id,string"`          // TopicID 话题ID
	ArticleCount    int   `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int   `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}
