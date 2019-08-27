package model

type DisucssItem struct {
	// 讨论ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 讨论分类
	Category *DiscussItemCategory `json:"category,omitempty"`
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

	// 发布人
	Creator *Creator `json:"creator"`

	// 图片
	ImageUrls []string `json:"img_urls"`
}

type DiscussItemCategory struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 名称
	Name string `json:"name"`
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

type Creator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name"`
	// 头像
	Avatar string `json:"avatar"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`
}
