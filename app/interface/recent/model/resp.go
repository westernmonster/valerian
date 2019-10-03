package model

type ItemTopic struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`

	// 简介
	Introduction string `json:"introduction"`

	// 封面
	Avatar *string `json:"avatar,omitempty"`

	// 成员数
	MemberCount int `json:"member_count"`

	// 文章数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`
}

type ItemArticle struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 文章标题
	Title string `json:"title"`
	// 封面
	Avatar *string `json:"avatar,omitempty"`
	// 内容
	Excerpt string `json:"excerpt"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 补充个数
	ReviseCount int `json:"revise_count"`
	// 评论数
	CommentCount int `json:"comment_count"`

	// 图片
	ImageUrls []string `json:"images"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`

	// 更新日期
	UpdatedAt int64 `json:"updated_at"`
}

type RecentItem struct {
	// 类型
	// topic, article
	Type string `json:"type"`

	// 话题
	Topic *ItemTopic `json:"topic,omitempty"`

	// 文章
	Article *ItemArticle `json:"article,omitempty"`
}

type RecentListResp struct {
	Items  []*RecentItem `json:"items"`
	Paging *Paging       `json:"paging"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
