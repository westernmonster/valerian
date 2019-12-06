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
	Title       string        `db:"title" json:"title,omitempty"`          // Title 标题
	Content     string        `db:"content" json:"content"`                // Content 内容HTML
	ContentText string        `db:"content_text" json:"content_text"`      // ContentText 内容纯文本
	Deleted     types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}

type AccountStat struct {
	AccountID       int64 `db:"account_id" json:"account_id,string"`      // AccountID 用户ID
	Following       int   `db:"following" json:"following"`               // Following 关注数
	Fans            int   `db:"fans" json:"fans"`                         // Fans 粉丝数
	ArticleCount    int   `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int   `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	TopicCount      int   `db:"topic_count" json:"topic_count"`           // TopicCount 讨论数
	Black           int   `db:"black" json:"black"`                       // Black 黑名单数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type TopicStat struct {
	TopicID         int64 `db:"topic_id" json:"topic_id,string"`          // TopicID 话题ID
	MemberCount     int   `db:"member_count" json:"member_count"`         // MemberCount 成员数
	ArticleCount    int   `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int   `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type DiscussionFile struct {
	ID           int64         `db:"id" json:"id,string"`                       // ID ID
	FileName     string        `db:"file_name" json:"file_name"`                // FileName 文件名
	FileURL      string        `db:"file_url" json:"file_url"`                  // FileURL 文件地址
	Seq          int32         `db:"seq" json:"seq"`                            // Seq 文件顺序
	DiscussionID int64         `db:"discussion_id" json:"discussion_id,string"` // DiscussionID 讨论ID
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                    // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`              // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`              // UpdatedAt 更新时间
	FileType     string        `db:"file_type" json:"file_type"`                // FileType 文件类型
	PdfURL       string        `db:"pdf_url" json:"pdf_url"`                    // PdfURL PDF文件
}

type DiscussionStat struct {
	DiscussionID int64 `db:"discussion_id" json:"discussion_id,string"` // DiscussionID 讨论ID
	LikeCount    int   `db:"like_count" json:"like_count"`              // LikeCount 喜欢数
	DislikeCount int   `db:"dislike_count" json:"dislike_count"`        // DislikeCount 反对数
	CommentCount int   `db:"comment_count" json:"comment_count"`        // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`              // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`              // UpdatedAt 更新时间
}

type ImageURL struct {
	ID         int64         `db:"id" json:"id,string"`               // ID ID
	TargetID   int64         `db:"target_id" json:"target_id,string"` // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`    // TargetType 目标类型
	URL        string        `db:"url" json:"url"`                    // URL 图片路径
	Deleted    types.BitBool `db:"deleted" json:"deleted"`            // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`      // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`      // UpdatedAt 更新时间
}
