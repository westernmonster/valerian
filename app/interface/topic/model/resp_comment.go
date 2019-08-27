package model

type CommentItem struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 内容
	Content string `json:"content"`
	// 是否精选
	Featured bool `json:"featured"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 发布日期
	CreatedAt int64 `json:"created_at"`
	// 发布人
	Author *Creator `json:"creator"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 子评论数
	ChildCommentsCount int `json:"child_comments_count"`
	// 子评论
	ChildComments []*ChildCommentItem `json:"child_comments"`
}

type ChildCommentItem struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 内容
	Content string `json:"content"`
	// 是否精选
	Featured bool `json:"featured"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 发布日期
	CreatedAt int64 `json:"created_at"`
	// 发布人
	Author *Creator `json:"creator"`
	// 喜欢数
	LikeCount int `json:"like_count"`

	// 被回复人
	ReplyToAuthor *Creator `json:"reply_to_author"`
}

type CommentListResp struct {
	Items  []*CommentItem `json:"items"`
	Paging *Paging        `json:"paging"`
	Total  int            `json:"total"`
	// 精选数
	FeaturedCount int `json:"featured_count"`
	// 评论数
	CommentsCount int `json:"comments_count"`
}
