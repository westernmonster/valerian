package model

type CommentResp struct {
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
	Creator *CommentCreator `json:"creator"`

	// 是否点赞
	Like bool `json:"like"`
	// 是否反对
	Dislike bool `json:"dislike"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 喜欢数
	DislikeCount int `json:"dislike_count"`
	// 子评论数
	ChildCommentsCount int `json:"child_comments_count"`
	// 子评论
	ChildComments []*ChildCommentItem `json:"child_comments,omitempty"`
}

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
	Creator *CommentCreator `json:"creator"`

	// 是否点赞
	Like bool `json:"like"`
	// 是否反对
	Dislike bool `json:"dislike"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 喜欢数
	DislikeCount int `json:"dislike_count"`
	// 子评论数
	ChildCommentsCount int `json:"child_comments_count"`
	// 子评论
	ChildComments []*ChildCommentItem `json:"child_comments,omitempty"`
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
	Creator *CommentCreator `json:"creator"`

	// 是否点赞
	Like bool `json:"like"`
	// 是否反对
	Dislike bool `json:"dislike"`

	// 喜欢数
	LikeCount int `json:"like_count"`
	// 喜欢数
	DislikeCount int `json:"dislike_count"`

	// 被回复人
	ReplyTo *CommentCreator `json:"reply_to"`
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

type CommentCreator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name"`
	// 头像
	Avatar string `json:"avatar"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`

	// 是否该文章/补充/讨论的作者
	IsAuthor bool `json:"is_author"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
