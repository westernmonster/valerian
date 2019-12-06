package model

type ArticleRelationResp struct {
	// 在话题目录中的ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 路径
	CatalogFullPath string `json:"catalog_full_path"`

	// 是否主话题
	Primary bool `json:"primary"`

	// 话题ID
	ToTopicID int64 `json:"to_topic_id,string"  swaggertype:"string"`
	// 话题名称
	Name string `json:"name"`
	// 话题头像
	Avatar string `json:"avatar,omitempty"`
	// 权限
	Permission string `json:"permission"`

	// 话题简介
	Introduction string `json:"introduction"`

	// 话题编辑权限
	EditPermission string `json:"edit_permission"`

	// 成员数
	MemberCount int32 `json:"member_count"`

	// 成员数
	ArticleCount int32 `json:"article_count"`

	// 讨论数
	DiscussionCount int32 `json:"discussion_count"`
}
