package model

type Creator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name"`
	// 头像
	Avatar string `json:"avatar"`

	// 自我介绍
	Introduction string `json:"introduction,omitempty"`
}

type TargetArticle struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 文章标题
	Title string `json:"title"`
	// 内容
	Excerpt string `json:"excerpt"`

	// 图片
	ImageUrls []string `json:"images"`

	// 补充个数
	ReviseCount int `json:"revise_count"`

	// 喜欢数
	LikeCount int `json:"like_count"`

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	Creator *Creator `json:"creator,omitempty"`
}

type TargetRevise struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 文章标题
	Title string `json:"title,omitempty"`
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

	Creator *Creator `json:"creator,omitempty"`
}

type TargetDiscuss struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 标题
	Title string `json:"title,omitempty"`
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

	Creator *Creator `json:"creator,omitempty"`
}

type TargetTopic struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`
	// 头像
	// 必须为URL
	Avatar string `json:"avatar"`

	// 成员数
	MemberCount int `json:"member_count"`

	// 成员数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`

	// 简介
	Introduction string `json:"introduction"`

	Creator *Creator `json:"creator,omitempty"`
}

type FavItem struct {
	ID int64 `json:"id,string" swaggertype:"string"`
	// 类型
	// topic, article, revise, discussion
	TargetType string `json:"target_type"`

	// 话题
	Topic *TargetTopic `json:"topic,omitempty"`

	// 文章
	Article *TargetArticle `json:"article,omitempty"`

	// 文章补充
	Revise *TargetRevise `json:"revise,omitempty"`

	// 讨论
	Discussion *TargetDiscuss `json:"discussion,omitempty"`
}

type FavListResp struct {
	Items  []*FavItem `json:"items"`
	Paging *Paging    `json:"paging"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
