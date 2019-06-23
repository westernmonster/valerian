package model

type TopicSearchResp struct {
	// 封面图
	// 必须为URL
	Cover *string `json:"cover"`

	// 背景图
	// 必须为URL
	Bg *string `json:"bg"`

	// 名称
	Name string `json:"name"`

	// 简介
	Introduction string `json:"introduction"`

	// 成员数
	MembersCount int `json:"members_count"`

	// 资源数量
	ResourcesCount int `json:"resources_count"`
}
