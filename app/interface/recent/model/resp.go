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

	// 资源数量
	ResourcesCount int `json:"resources_count"`

	// 成员数
	MembersCount int `json:"members_count"`
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
}

type RecentTopicListResp struct {
	Items  []*ItemTopic `json:"items"`
	Paging *Paging      `json:"paging"`
}

type RecentArticleListResp struct {
	Items  []*ItemArticle `json:"items"`
	Paging *Paging        `json:"paging"`
}

type Paging struct {
	// 统计数量
	Total *int `json:"total,omitempty"`
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
