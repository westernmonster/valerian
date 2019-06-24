package model

type UpdateArticleRelation struct {
	// ID 文章在话题目录中的ID
	// 当只是更新的情况，去除当前主话题，设置另一个关联为主话题时候传入
	// 如果是增加的关联，则无需传入，还是仅仅传入 parent_id, topic_id, primary 属性就可以
	ID *int64 `json:"id,string,omitempty" swaggertype:"string"`
	// 类目分类ID 如果根目录则传0
	ParentID int64 `json:"parent_id,string" swaggertype:"string"`

	// 所关联话题 ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 是否主话题
	Primary bool `json:"primary"`
}

type ArgSaveArticleRelations struct {
	ArticleID int64                    `json:"article_id,string" swaggertype:"string"`
	Items     []*UpdateArticleRelation `json:"items"`
}
