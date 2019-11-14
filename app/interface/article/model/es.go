package model

type ESArticle struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 内容
	ContentText *string `json:"content_text,omitempty"`

	// 内容
	Excerpt *string `json:"excerpt,omitempty"`

	ChangeDesc *string `json:"change_desc,omitempty"`

	//  禁止补充
	DisableRevise *bool `json:"disable_revise,omitempty"`

	//  禁止评论
	DisableComment *bool `json:"disable_comment,omitempty"`

	Creator *ESCreator `json:"creator,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 喜欢数
	LikeCount int `json:"like_count"`
	// 反对数
	DislikeCount int `json:"dislike_count"`
	// 补充个数
	ReviseCount int `json:"revise_count"`
	// 评论数
	CommentCount int `json:"comment_count"`
}

type ESCreator struct {
	// ID ID
	ID int64 `json:"id,string,omitempty"  swaggertype:"string"`
	// 用户名
	UserName *string `json:"user_name,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty"`

	Introduction *string `json:"introduction,omitempty"`
}
