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

type ReviseFileResp struct {
	ID       int64  `json:"id,string" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url`   // FileURL 文件地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}

type ReviseDetailResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	Creator *Creator `json:"creator"`

	Files []*ReviseFileResp `json:"files"`

	// 内容
	Content string `json:"content"`

	// 赞数
	LikeCount int `json:"like_count"`

	// 评论数
	CommentCount int `json:"comment_count"`
}
