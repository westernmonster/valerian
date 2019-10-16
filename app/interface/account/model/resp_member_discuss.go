package model

type MemberDiscuss struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 评论内容
	Excerpt string `json:"excerpt"`

	// 喜欢数
	LikeCount int `json:"like_count"`

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	// 图片
	ImageUrls []string `json:"images"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type MemberDiscussResp struct {
	Items  []*MemberDiscuss `json:"items"`
	Paging *Paging          `json:"paging"`
}
