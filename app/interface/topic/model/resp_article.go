package model

type ArticleResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title string `json:"title"`

	// 内容
	Content string `json:"content"`

	// 集合ID
	ArticleSetID int64 `json:"article_set_id,string"  swaggertype:"string"`

	// 文章语言
	Locale string `json:"locale"`

	//  封面图
	Cover *string `json:"cover,omitempty"`
	// 简介
	Introduction string `json:"introduction"` // Introduction 话题简介

	// 是否私有
	Private bool `json:"private"`

	// 版本名称
	VersionName string `json:"version_name"`

	// 版本顺序
	Seq int `json:"seq"`

	Files []*ArticleFileResp `json:"files"`

	Relations []*ArticleRelationResp `json:"relations"`

	Versions []*ArticleVersionResp `json:"versions"`

	// 属性
	ArticleMeta *ArticleMeta `json:"meta,omitempty"`
}

type ArticleRelationResp struct {
	// 话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 话题名
	TopicName string `json:"topic_name"`
	// 话题在目的中的ID
	TopicCatalogID int64 `json:"topic_catalog_id,string" swaggertype:"string"`
	// 全路径名
	CatalogFullPath string `json:"catalog_full_path"`
	// 是否主话题
	Primary bool `json:"primary"`
}

type ArticleFileResp struct {
	ID       int64  `json:"id,string" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url`   // FileURL 文件地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}

type ArticleMeta struct {
	// 是否能编辑
	CanEdit bool `json:"can_edit"`

	Like bool `json:"like"`

	LikeCount int `json:"like_count"`

	Read bool `json:"read"`

	Fav bool `json:"fav"`
}
