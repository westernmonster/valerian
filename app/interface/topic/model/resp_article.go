package model

type ArticleResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title string `json:"title"`

	// 内容
	Content string `json:"content"`

	CreatedBy int64 `json:"created_by,string"`

	Creator *BasicAccountResp `json:"creator"`

	//  封面图
	Avatar *string `json:"avatar,omitempty"`
	// 简介
	Introduction string `json:"introduction"` // Introduction 话题简介

	// 是否私有
	Private bool `json:"private"`

	Files []*ArticleFileResp `json:"files"`

	Relations []*ArticleRelationResp `json:"relations"`

	// 属性
	ArticleMeta *ArticleMeta `json:"meta,omitempty"`

	PrimaryTopicMeta *TopicMeta `json:"primary_topic_meta,omitempty"`
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
	Like bool `json:"like"`

	LikeCount int `json:"like_count"`

	Read bool `json:"read"`

	Fav bool `json:"fav"`

	FavCount int `json:"fav_count"`
}
