package model

type RelatedTopicResp struct {
	// 话题版本ID
	TopicVersionID int64 `db:"topic_version_id" json:"topic_version_id,string" swaggertype:"string"`

	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`

	// 话题名
	TopicName string `db:"topic_name" json:"topic_name"`
	// 封面图
	// 必须为URL
	Cover *string `db:"cover" json:"cover"`

	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`
	// 关联类型
	// normal 普通关联
	// strong 强关联
	Type string `db:"type" json:"type"`

	// 顺序
	Seq int `db:"seq" json:"seq"`

	// 简介
	Introduction string `db:"introduction" json:"introduction"`

	// 成员数
	MembersCount int `json:"members_count"`

	// 资源数量
	ResourcesCount int `json:"resources_count"`

	TopicMeta *TopicMeta `json:"meta,omitempty"`
}
