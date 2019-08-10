package model

type AuthTopicResp struct {
	ToTopicID int64 `db:"to_topic_id" json:"to_topic_id,string"  swaggertype:"string"` // ToTopicID To Topic ID

	Permission string `db:"permission" json:"permission"` // Permission 权限

	// 成员数
	MembersCount int `db:"-" json:"members_count"`

	// 封面图
	// 必须为URL
	Avatar *string `db:"avatar" json:"avatar"`

	// 名称
	Name string `db:"name" json:"name"`
}
