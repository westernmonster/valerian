package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgSaveAuthTopics struct {
	// 话题ID
	TopicID int64 `json:"topic_id,string"  swaggertype:"string"`

	// 授权话题
	AuthTopics []*ArgAuthTopic `json:"auth_topics,omitempty"`
}

func (p *ArgSaveAuthTopics) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.AuthTopics),
	)
}

type ArgAuthTopic struct {
	// 授权话题ID
	TopicID int64 `json:"topic_id,string"  swaggertype:"string"`

	// 类型
	// view  // 只允许查看
	// edit // 允许所有成员编辑
	// admin_edit // 只允许管理员和主理人编辑
	Permission string `json:"permission"`
}

func (p *ArgAuthTopic) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Permission, validation.Required, validation.In(AuthPermissionView, AuthPermissionEdit, AuthPermissionAdminEdit)),
	)
}
