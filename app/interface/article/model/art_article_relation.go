package model

type UpdateArticleRelation struct {
	ID             *int64 `json:"id,string,omitempty" swaggertype:"string"`
	TopicCatalogID int64  `json:"topic_catalog_id,string" swaggertype:"string"`
	// 是否主话题
	Primary bool `json:"primary"`
}

type ArgSaveArticleRelations struct {
	ArticleID int64                    `json:"article_id,string" swaggertype:"string"`
	Items     []*UpdateArticleRelation `json:"items"`
}
