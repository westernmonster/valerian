package model

type TopicStat struct {
	TopicID         int64 `db:"topic_id" json:"topic_id,string"`          // TopicID 话题ID
	MemberCount     int32 `db:"member_count" json:"member_count"`         // MemberCount 成员数
	ArticleCount    int32 `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int32 `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type ArticleStat struct {
	ArticleID    int64 `db:"article_id" json:"article_id,string"` // ArticleID 讨论ID
	LikeCount    int32 `db:"like_count" json:"like_count"`        // LikeCount 喜欢数
	DislikeCount int32 `db:"dislike_count" json:"dislike_count"`  // DislikeCount 反对数
	ReviseCount  int32 `db:"revise_count" json:"revise_count"`    // ReviseCount 补充数
	CommentCount int32 `db:"comment_count" json:"comment_count"`  // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type ReviseStat struct {
	ReviseID     int64 `db:"revise_id" json:"revise_id,string"`  // ReviseID 补充ID
	LikeCount    int32 `db:"like_count" json:"like_count"`       // LikeCount 喜欢数
	DislikeCount int32 `db:"dislike_count" json:"dislike_count"` // DislikeCount 反对数
	CommentCount int32 `db:"comment_count" json:"comment_count"` // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`       // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`       // UpdatedAt 更新时间
}

type DiscussionStat struct {
	DiscussionID int64 `db:"discussion_id" json:"discussion_id,string"` // DiscussionID 讨论ID
	LikeCount    int32 `db:"like_count" json:"like_count"`              // LikeCount 喜欢数
	DislikeCount int32 `db:"dislike_count" json:"dislike_count"`        // DislikeCount 反对数
	CommentCount int32 `db:"comment_count" json:"comment_count"`        // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`              // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`              // UpdatedAt 更新时间
}

type AccountStat struct {
	AccountID       int64 `db:"account_id" json:"account_id,string"`      // AccountID 用户ID
	Following       int32 `db:"following" json:"following"`               // Following 关注数
	Fans            int32 `db:"fans" json:"fans"`                         // Fans 粉丝数
	ArticleCount    int32 `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int32 `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	TopicCount      int32 `db:"topic_count" json:"topic_count"`           // TopicCount 讨论数
	Black           int32 `db:"black" json:"black"`                       // Black 黑名单数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}
