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

	// 更新时间
	UpdatedAt int64 `json:"updated_at"`

	CreatedBy int64 `json:"created_by"`

	Updator *Creator `json:"updator,omitempty"` // Updator 更新人

	ChangeDesc string `json:"change_desc"` // ChangeDesc 修订说明
}

type ArticleListCacheResp struct {
	Items  []*ArticleResp `json:"items"`
}

type ArticleFileResp struct {
	ID        int64  `json:"id,string" swaggertype:"string"`
	FileName  string `json:"file_name"` // FileName 文件名
	FileURL   string `json:"file_url"`  // FileURL 文件地址
	FileType  string `json:"file_type"` // 文件类型
	PdfURL    string `json:"pdf_url"`   // PDF地址
	Seq       int    `json:"seq"`       // Seq 文件顺序
	CreatedAt int64  `json:"created_at"`
}

type ArticleMeta struct {
	Like bool `json:"like"`

	LikeCount int32 `json:"like_count"`

	// 反对数
	DislikeCount int32 `json:"dislike_count"`

	Read bool `json:"read"`

	CanEdit bool `json:"can_edit"`

	Fav bool `json:"fav"`

	Dislike bool `json:"dislike"`

	// 补充数
	ReviseCount int32 `json:"revise_count"`

	// 评论数
	CommentCount int32 `json:"comment_count"`
}

type Creator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name"`
	// 头像
	Avatar string `json:"avatar"`

	// 自我介绍
	Introduction string `json:"introduction,omitempty"`
}

type ArticleItem struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 文章标题
	Title string `json:"title"`
	// 内容
	Excerpt string `json:"excerpt"`
	// 图片
	ImageUrls []string `json:"images"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 反对数
	DislikeCount int `json:"dislike_count"`
	// 补充个数
	ReviseCount int `json:"revise_count"`
	// 评论数
	CommentCount int `json:"comment_count"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type ArticleListResp struct {
	Items  []*ArticleItem `json:"items"`
	Paging *Paging        `json:"paging"`
}
