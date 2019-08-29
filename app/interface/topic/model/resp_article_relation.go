package model

type ArticleRelationResp struct {
	// 在话题目录中的ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题ID
	TopicID int64 `json:"topic_id,string"  swaggertype:"string"`
	// 是否主话题
	Primary bool `json:"primary"`
	// 话题名称
	TopicName string `json:"topic_name"`
	// 话题头像
	TopicAvatar *string `json:"topic_avatar,omitempty"`
	// 权限
	Permission string `json:"permission"`

	// 话题编辑权限
	TopicEditPermission string `json:"topic_edit_permission"`

	// 路径
	CatalogFullPath string `json:"catalog_full_path"`
}
