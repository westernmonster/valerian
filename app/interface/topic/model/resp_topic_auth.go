package model

type AuthTopicResp struct {
	ToTopicID int64 `json:"to_topic_id,string"  swaggertype:"string"` // ToTopicID To Topic ID

	EditPermission string `json:"edit_permission"` // 话题编辑权限

	Permission string `json:"permission"` // Permission 权限

	// 成员数
	MemberCount int32 `json:"member_count"`

	// 封面图
	// 必须为URL
	Avatar string `json:"avatar"`

	// 名称
	Name string `json:"name"`
}

type Auth2CurrentTopicsResp struct {
	Items  []*TargetTopic `json:"items"`
	Paging *Paging        `json:"paging"`
}
