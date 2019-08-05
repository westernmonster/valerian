package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgInnerTopicMember struct {
	// 成员ID
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
	// 角色
	// user 普通用户
	// admin 管理员
	Role string `json:"role"`
}

type ArgCreateTopic struct {
	// 成员
	Members []*ArgInnerTopicMember `json:"members,omitempty"`

	// 话题分类
	Catalogs []*TopicLevel1Catalog `json:"catalogs,omitempty"`

	// 封面图
	// 必须为URL
	Cover *string `json:"cover,omitempty"`

	// 背景图
	// 必须为URL
	Bg *string `json:"bg,omitempty"`

	// 名称
	Name string `json:"name"`

	// 简介
	Introduction string `json:"introduction"`

	// 分类视图
	// section 章节
	// column 栏目
	CatalogViewType string `json:"catalog_view_type"`

	// 话题类型
	TopicType int `json:"topic_type"`

	// 话题首页
	// feed 动态
	// catalog 目录
	// discuss 讨论
	TopicHome string `json:"topic_home"`

	// 是否私密
	IsPrivate bool `json:"is_private"`

	// 开启群聊
	AllowChat bool `json:"allow_chat"`

	// 允许讨论
	AllowDiscuss bool `json:"allow_discuss"`

	// 查看权限
	// public 公开
	// join 加入
	ViewPermission string `json:"view_permission"`

	// 编辑权限
	// "member"
	// "only_admin"
	EditPermission string `json:"edit_permission"`

	// 加入权限
	// 默认为 member
	// all 注册用户
	// only_approve 用户需审批
	// cert_approve 认证需审批
	// manual_add 仅管理员手工添加
	JoinPermission string `json:"join_permission"`

	// 重要标记
	Important bool `db:"important" json:"important"`

	// 消息免打扰
	MuteNotification bool `db:"mute_notification" json:"mute_notification"`
}

func (p *ArgCreateTopic) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Members, validation.Required),
		validation.Field(&p.TopicType, validation.Required),
		validation.Field(&p.Cover, is.URL),
		validation.Field(&p.Bg, is.URL),
		validation.Field(&p.Name, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.Introduction, validation.Required, validation.RuneLength(0, 1000)),
		validation.Field(&p.CatalogViewType, validation.Required, validation.In(CatalogViewTypeColumn, CatalogViewTypeSection)),
		validation.Field(&p.TopicHome, validation.Required, validation.In(TopicHomeFeed, TopicHomeCataglog, TopicHomeDiscuss)),
		validation.Field(&p.ViewPermission, validation.Required, validation.In(ViewPermissionJoin, ViewPermissionPublic)),
		validation.Field(&p.EditPermission, validation.Required, validation.In(EditPermissionMember, EditPermissionAdmin)),
		validation.Field(&p.JoinPermission,
			validation.Required,
			validation.In(JoinPermissionMember, JoinPermissionMemberApprove, JoinPermissionCertApprove, JoinPermissionManualAdd),
		),
	)
}

type ArgUpdateTopic struct {
	// 话题ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 封面图
	// 必须为URL
	Cover *string `json:"cover,omitempty"`

	// 背景图
	// 必须为URL
	Bg *string `json:"bg,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 简介
	Introduction *string `json:"introduction,omitempty"`

	// 分类视图
	// section 章节
	// column 栏目
	CatalogViewType *string `json:"catalog_view_type,omitempty"`

	// 话题类型
	TopicType *int `json:"topic_type,omitempty"`

	// 话题首页
	// feed 动态
	// catalog 目录
	// discuss 讨论
	TopicHome *string `json:"topic_home,omitempty"`

	// 是否私密
	IsPrivate *bool `json:"is_private,omitempty"`

	// 开启群聊
	AllowChat *bool `json:"allow_chat,omitempty"`

	// 查看权限
	// public 公开
	// join 加入
	ViewPermission *string `json:"view_permission"`

	// 编辑权限
	// "member"
	// "admin"
	EditPermission *string `json:"edit_permission"`

	// 加入权限
	// 默认为 member
	// member 注册用户
	// member_approve 用户需审批
	// cert_approve 认证需审批
	// manual_add 仅管理员手工添加
	JoinPermission *string `json:"join_permission"`

	// 重要标记
	Important *bool `json:"important,omitempty"`

	// 消息免打扰
	MuteNotification *bool `json:"mute_notification,omitempty"`
}

func (p *ArgUpdateTopic) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Cover, is.URL), validation.Field(&p.Bg, is.URL),
		validation.Field(&p.Name, validation.RuneLength(0, 250)),
		// TODO: Web安全性
		validation.Field(&p.Introduction, validation.RuneLength(0, 1000)),
		validation.Field(&p.CatalogViewType, validation.In(CatalogViewTypeColumn, CatalogViewTypeSection)),
		validation.Field(&p.TopicHome, validation.In(TopicHomeFeed, TopicHomeCataglog, TopicHomeDiscuss)),
		validation.Field(&p.ViewPermission, validation.Required, validation.In(ViewPermissionJoin, ViewPermissionPublic)),
		validation.Field(&p.EditPermission, validation.Required, validation.In(EditPermissionMember, EditPermissionAdmin)),
		validation.Field(&p.JoinPermission, validation.Required, validation.In(JoinPermissionMember, JoinPermissionMemberApprove, JoinPermissionCertApprove, JoinPermissionManualAdd)),
	)
}
