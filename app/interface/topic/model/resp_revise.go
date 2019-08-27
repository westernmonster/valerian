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

type ReviseResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	Content string `json:"string"`

	Files []*ReviseFileResp `json:"files"`

	Creator *Creator `json:"creator"`
}

type ReviseFileResp struct {
	ID       int64  `json:"id,string" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url`   // FileURL 文件地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}
