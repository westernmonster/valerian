package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CreateTopicReq struct {
	// 成员
	Members []*TopicMember `json:"members"`
	// 封面图
	Cover string `json:"cover"`
	// 名称
	Name string `json:"name"`
	// 简介
	Introduction string `json:"introduction"`
	// 关联话题
	RelatedTopics []*RelatedTopic `json:"related_topics"`
	// 话题分类
	Categories []*TopicCategoryParentItem `json:"categories"`
	// 分类视图
	CategoryViewType int `db:"category_view_type" json:"category_view_type"`

	// TODO: 优先显示

	// 是否私密
	IsPrivate bool `json:"is_private"`
	// 允许讨论
	AllowDiscuss bool `json:"allow_discuss"`
	// 编辑权限
	EditPermission int `json:"edit_permission"`
	// 查看权限
	ViewPermission int `json:"view_permission"`
	// 加入权限
	JoinPermission int `json:"join_permission"`
	// 重要标记
	Important bool `db:"important" json:"important"`
	// 消息免打扰
	MuteNotification bool `db:"mute_notification" json:"mute_notification"`
}

func (p *CreateTopicReq) Validate() error {
	// TODO: 其他字段验证
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Members, validation.Required.Error(`请选择成员`)),
		validation.Field(&p.Cover,
			validation.Required.Error(`请选择封面图`),
			is.URL.Error("封面图格式不正确"),
		),
		validation.Field(&p.Name,
			validation.Required.Error(`请输入话题名`),
			validation.RuneLength(0, 250).Error(`话题名最大长度为250个字符`),
		),
		validation.Field(&p.Introduction,
			validation.Required.Error(`请输入话题简介`),
			validation.RuneLength(0, 1000).Error(`话题简介最大长度为1000个字符`),
		),
	)
}

type RelatedTopic struct {
	// 关联话题ID
	TopicID int64 `json:"topic_id"`
	// 类型：1:强关系 2:弱关系
	Type int `json:"type"`
}

func (p *RelatedTopic) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required.Error(`请传入关联话题ID`)),
		validation.Field(&p.Type, validation.Required.Error(`请传入关联话题类型`)),
	)
}

type TopicMember struct {
	// 成员ID
	AccountID int64 `json:"account_id"`
	// 角色
	Role int `json:"role"`
}

func (p *TopicMember) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required.Error(`请传入成员ID`)),
		validation.Field(&p.Role, validation.Required.Error(`请传入成员角色`)),
	)
}

type BulkSaveTopicMembersReq struct {
	// 成员
	Members []*TopicMember `json:"members"`
}

type UpdateTopicReq struct {
	// 成员
	Members []*TopicMember `json:"members,omitempty"`
	// 封面图
	Cover *string `json:"cover,omitempty"`
	// 名称
	Name *string `json:"name,omitempty"`
	// 简介
	Introduction *string `json:"introduction,omitempty"`
	// 关联话题
	RelatedTopics []*RelatedTopic `json:"related_topics,omitempty"`
	// 话题分类
	Categories []*TopicCategoryParentItem `json:"categories,omitempty"`
	// 分类视图
	CategoryViewType *int `json:"category_view_type,omitempty"`

	// TODO: 优先显示

	// 是否私密
	IsPrivate *bool `json:"is_private,omitempty"`
	// 允许讨论
	AllowDiscuss *bool `json:"allow_discuss,omitempty"`
	// 编辑权限
	EditPermission *int `json:"edit_permission,omitempty"`
	// 查看权限
	ViewPermission *int `json:"view_permission,omitempty"`
	// 加入权限
	JoinPermission *int `json:"join_permission,omitempty"`
	// 重要标记
	Important *bool `json:"important,omitempty"`
	// 消息免打扰
	MuteNotification *bool `json:"mute_notification,omitempty"`
}
