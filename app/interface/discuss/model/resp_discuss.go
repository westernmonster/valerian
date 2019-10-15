package model

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

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}

type DiscussItem struct {
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

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`

	// 更新日期
	UpdatedAt int64 `json:"updated_at"`

	// 发布人
	Creator *Creator `json:"creator"`

	// 图片
	ImageUrls []string `json:"images"`
}

type DiscussItemCategory struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 名称
	Name string `json:"name"`
}

type DiscussListResp struct {
	Items  []*DiscussItem `json:"items"`
	Paging *Paging        `json:"paging"`
}

type DiscussDetailResp struct {
	// 讨论ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 发布人
	Creator *Creator `json:"creator"`

	// 话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 讨论分类
	Category *DiscussItemCategory `json:"category,omitempty"`
	// 标题
	Title *string `json:"title"`
	// 内容
	Content string `json:"content"`

	Files []*DiscussFileResp `json:"files"`

	Fav bool `json:"fav"`

	Like bool `json:"like"`

	Dislike bool `json:"dislike"`

	// 喜欢数
	LikeCount int `json:"like_count"`

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	// 是否可以编辑
	// 讨论所属用户以及管理员可以编辑
	CanEdit bool `json:"can_edit"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`

	// 更新日期
	UpdatedAt int64 `json:"updated_at"`
}

type DiscussFileResp struct {
	ID       int64  `json:"id,string" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url`   // FileURL 文件地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}
