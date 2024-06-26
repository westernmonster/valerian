package model

// 补充
type ReviseItem struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	Creator *Creator `json:"creator"`

	// 摘录
	Excerpt string `json:"excerpt"`

	// 图片
	ImageUrls []string `json:"images"`

	// 赞数
	LikeCount int `json:"like_count"`

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type ReviseListResp struct {
	Items       []*ReviseItem `json:"items"`
	Paging      *Paging       `json:"paging"`
	ReviseCount int32         `json:"revise_count"`
}

type ReviseFileResp struct {
	ID       int64  `json:"id,string" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url"`  // FileURL 文件地址
	FileType string `json:"file_type"` // 文件类型
	PdfURL   string `json:"pdf_url"`   // PDF地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}

type ReviseDetailResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	ArticleID int64 `json:"article_id,string" swaggertype:"string"` // ArticleID 文章ID

	Creator *Creator `json:"creator"`

	Files []*ReviseFileResp `json:"files"`

	// 文章标题
	Title string `json:"title"`

	// 内容
	Content string `json:"content"`

	CanEdit bool `json:"can_edit"`

	// 是否收藏
	Fav bool `json:"fav"`

	// 是否点赞
	Like bool `json:"like"`

	// 是否反对
	Dislike bool `json:"dislike"`

	// 赞数
	LikeCount int `json:"like_count"`

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type ReviseDetailListCacheResp struct {
	Items []*ReviseDetailResp `json:"items"`
}

type Paging struct {
	// // 统计数量
	// Total int `json:"total,omitempty"`
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
