package model

type RelatedTopicResp struct {
	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`
	// 话题名
	Name string `db:"name" json:"name"`
	// 封面图
	// 必须为URL
	Cover string `db:"cover" json:"cover"`
	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`
	// 关联类型
	// normal 普通关联
	// strong 强关联
	Type string `db:"type" json:"type"`
	// 简介
	Introduction string `db:"introduction" json:"introduction"`

	// 成员数
	MembersCount int `json:"members_count"`

	// 资源数量
	ResourcesCount int `json:"resources_count"`

	// 是否关注
	IsFollowed bool `json:"is_followed"`

	// 能否关注
	// 根据当前用户匹配权限返回
	CanFollow bool `json:"can_follow"`
}
