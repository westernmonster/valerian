package model

type ReviseItem struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	Creator *Creator `json:"creator"`

	// 摘录
	Excerpt string `json:"excerpt"`

	// 图片
	ImageUrls []string `json:"img_urls"`

	// 赞数
	LikeCount int `json:"like_count"`

	// 评论数
	CommentCount int `json:"comment_count"`
}

type ReviseListResp struct {
	Items  []*ReviseItem `json:"items"`
	Paging *Paging       `json:"paging"`
}
