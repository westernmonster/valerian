package model

type MemberDiscuss struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 标题
	Title *string `json:"title,omitempty"`
	// 评论内容
	Excerpt string `json:"excerpt"`
	// 评论数
	CommentCount int `json:"comment_count"`

	// 图片
	ImageUrls []string `json:"img_urls"`
}

type MemberDiscussResp struct {
	Items  []*MemberDiscuss `json:"items"`
	Paging *Paging          `json:"paging"`
}
