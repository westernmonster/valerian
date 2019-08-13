package model

type DisucssItem struct {
	// 话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 讨论分类ID
	DiscussCategoryID int64 `json:"discuss_category_id,string" swaggertype:"string"`
	// 讨论分类名称
	DiscussCategoryName string `json:"discuss_category_name"`
	// 标题
	Title *string `json:"title"`
	// 摘录
	Excerpt string `json:"excerpt"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 评论数
	CommentCount int `json:"comment_count"`
	// 发布日期
	CreatedAt int64 `json:"created_at"`
}

type DiscussListResp struct {
	Items  []*DisucssItem `json:"items"`
	Paging *Paging        `json:"paging"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
