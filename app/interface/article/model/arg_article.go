package model

type ArgAddArticle struct {
	// 标题
	Title string `json:"title"`
	// 内容
	Content string `json:"content"`
	//  封面图
	Cover *string `json:"cover,omitempty"`
	// 简介
	Introduction string `json:"introduction"` // Introduction 话题简介
	// 是否私有
	Private bool `json:"private"`

	// 文章语言
	Locale string `json:"locale"`

	// 文章集合ID
	ArticleSetID *int64 `json:"article_set_id,string,omitempty"  swaggertype:"string"`

	// 版本名称
	VersionName string `json:"version_name"`

	Files []*AddArticleFile `json:"files"`

	Relations []*AddArticleRelation `json:"relations"`
}

type AddArticleFile struct {
	FileName string  `json:"file_name"`          // FileName 文件名
	FileURL  *string `json:"file_url,omitempty"` // FileURL 文件地址
	Seq      int     `json:"seq"`                // Seq 文件顺序
}

type AddArticleRelation struct {
	// 类目分类ID 如果根目录则传0
	ParentID int64 `json:"parent_id,string" swaggertype:"string"`

	// 所关联话题 ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`

	// 是否主话题
	Primary bool `json:"primary"`
}

type ArgUpdateArticle struct {
	// 文章ID
	ID *int64 `json:"id,string"  swaggertype:"string"`
	// 标题
	Title *string `json:"title,omitempty"`
	// 内容
	Content *string `json:"content,omitempty"`

	// 文章语言
	Locale *string `json:"locale"`

	//  封面图
	Cover *string `json:"cover,omitempty"`
	// 简介
	Introduction *string `json:"introduction,omitempty"` // Introduction 话题简介
	// 是否私有
	Private *bool `json:"private,omitempty"`

	// 版本名称
	VersionName *string `json:"version_name"`
}
