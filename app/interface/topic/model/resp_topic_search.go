package model

type TopicSearchItem struct {
	ID int64 `json:"id,string" swaggertype:"string"`
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

	// 集合ID
	TopicSetID int64 `json:"topic_set_id,string"  swaggertype:"string"`

	VersionName string `json:"version_name"`

	// 版本列表
	Versions []*TopicVersionResp `json:"versions"`
}

type TopicSearchPaging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
}

type TopicSearchResp struct {
	Items  []*TopicSearchItem `json:"items"`
	Paging *TopicSearchPaging `json:"paging"`
}
