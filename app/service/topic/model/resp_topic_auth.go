package model

type AuthTopicResp struct {
	// 授权话题ID
	ToTopicID int64 `json:"to_topic_id,string"  swaggertype:"string"` // ToTopicID To Topic ID

	// 话题编辑权限
	EditPermission string `json:"edit_permission"`

	// 授权权限
	Permission string `json:"permission"` // Permission 权限

	// 成员数
	MemberCount int `json:"member_count"`

	// 封面图
	// 必须为URL
	Avatar *string `db:"avatar" json:"avatar"`

	// 话题名称
	Name string `db:"name" json:"name"`
}
