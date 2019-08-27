package model

type ArticleResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title string `json:"title"`

	// 内容
	Content string `json:"content"`

	//  禁止补充
	DisableRevise bool `json:"disable_revise"`

	//  禁止评论
	DisableComment bool `json:"disable_comment"`

	Creator *Creator `json:"creator"`

	Files []*ArticleFileResp `json:"files"`

	Relations []*ArticleRelationResp `json:"relations"`

	// 属性
	ArticleMeta *ArticleMeta `json:"meta,omitempty"`

	// 创建时间
	CreatedAt int64 `json:"created_at"`

	CreatedBy int64 `json:"-"`
}

type ArticleFileResp struct {
	ID       int64  `json:"id,string" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url`   // FileURL 文件地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}

type ArticleMeta struct {
	Like bool `json:"like"`

	LikeCount int `json:"like_count"`

	Read bool `json:"read"`

	Fav bool `json:"fav"`

	// 补充数
	ReviseCount int `json:"revise_count"`

	// 评论数
	CommentCount int `json:"comment_count"`
}
