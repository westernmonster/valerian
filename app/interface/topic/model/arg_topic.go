package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgCreateTopic struct {
	// 成员
	Members []*ArgCreateTopicMember `json:"members,omitempty"`

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

	// 关联话题
	RelatedTopics []*ArgRelatedTopic `json:"related_topics"`

	// 话题分类
	Versions []*ArgCreateTopicVersion `json:"versions"`

	// 分类视图
	// section 章节
	// column 栏目
	CatalogViewType string `json:"catalog_view_type"`

	// 话题类型
	TopicType int `json:"topic_type"`

	// 话题首页
	// introduction 简介
	// feed 动态
	// catalog 目录
	// discussion 讨论
	// chat 群聊
	TopicHome string `json:"topic_home"`

	// 是否私密
	IsPrivate bool `json:"is_private"`

	// 开启群聊
	AllowChat bool `json:"allow_chat"`

	// 允许讨论
	AllowDiscuss bool `json:"allow_discuss"`

	// 编辑权限
	// "id_cert"
	// "work_cert"
	// "id_cert_joined"
	// "work_cert_joined"
	// "approved_id_cert_joined"
	// "approved_work_cert_joined"
	// "only_admin"
	EditPermission string `json:"edit_permission"`

	// 查看权限
	// public 公开
	// join 加入
	ViewPermission string `json:"view_permission"`

	// 加入权限
	// 默认为 member
	// member 注册用户
	// id_cert 身份认证自由关注
	// work_cert 业内认证自由关注
	// member_approve 注册用户，需审批
	// id_cert_approve 身份认证自由关注，需审批
	// work_cert_approve 业内认证自由关注，需审批
	// admin_add 仅管理员手工添加
	// purchase 付费购买
	// vip Pro 用户自由关注
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
		validation.Field(&p.Versions, validation.Required),
		validation.Field(&p.Introduction, validation.Required, validation.RuneLength(0, 1000)),
		validation.Field(&p.CatalogViewType, validation.Required, validation.In(CatalogViewTypeColumn, CatalogViewTypeSection)),
		validation.Field(&p.TopicHome, validation.Required,
			validation.In(TopicHomeIntroduction, TopicHomeFeed, TopicHomeCataglog, TopicHomeDiscussion, TopicHomeChat),
		),
		validation.Field(&p.EditPermission,
			validation.Required,
			validation.In(EditPermissionIDCert, EditPermissionWorkCert, EditPermissionIDCertJoined, EditPermissionWorkCertJoined, EditPermissionApprovedIDCertJoined, EditPermissionApprovedWorkCertJoined, EditPermissionAdmin),
		),
		validation.Field(&p.ViewPermission,
			validation.Required,
			validation.In(ViewPermissionJoin, ViewPermissionPublic),
		),
		validation.Field(&p.JoinPermission,
			validation.Required,
			validation.In(JoinPermissionMember, JoinPermissionIDCert, JoinPermissionWorkCert, JoinPermissionMemberApprove, JoinPermissionIDCertApprove, JoinPermissionWorkCertApprove, JoinPermissionAdminAdd, JoinPermissionPurchase, JoinPermissionVIP),
		),
	)
}

type ArgCreateTopicVersion struct {
	// 顺序
	Seq int `json:"seq"`

	// 版本名称
	Name string `json:"name"`

	// 话题分类
	Catalogs []*TopicLevel1Catalog `json:"catalogs"`
}

func (p *ArgCreateTopicVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Catalogs),
	)
}

type ArgCreateTopicMember struct {
	// 成员ID
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
	// 角色
	// user 普通用户
	// admin 管理员
	Role string `json:"role"`
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
	// introduction 简介
	// feed 动态
	// catalog 目录
	// discussion 讨论
	// chat 群聊
	TopicHome *string `json:"topic_home,omitempty"`

	// 是否私密
	IsPrivate *bool `json:"is_private,omitempty"`

	// 开启群聊
	AllowChat *bool `json:"allow_chat,omitempty"`

	// 编辑权限
	// "id_cert"
	// "work_cert"
	// "id_cert_joined"
	// "work_cert_joined"
	// "approved_id_cert_joined"
	// "approved_work_cert_joined"
	// "only_admin"
	EditPermission *string `json:"edit_permission,omitempty"`

	// 查看权限
	// public 公开
	// join 加入
	ViewPermission *string `json:"view_permission,omitempty"`

	// 加入权限
	// free  自由加入
	// auth_free 认证用户自由加入
	// approve 经批准
	// auth_approve 认证用户经批准
	// admin 仅管理员添加
	// purchase 购买
	JoinPermission *string `json:"join_permission,omitempty"`

	// 重要标记
	Important *bool `json:"important,omitempty"`

	// 消息免打扰
	MuteNotification *bool `json:"mute_notification,omitempty"`
}

func (p *ArgUpdateTopic) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Cover,
			is.URL.Error("封面图格式不正确"),
		),
		validation.Field(&p.Bg,
			is.URL.Error("背景图格式不正确"),
		),
		validation.Field(&p.Name,
			validation.RuneLength(0, 250).Error(`话题名最大长度为250个字符`),
		),
		// TODO: Web安全性
		validation.Field(&p.Introduction,
			validation.RuneLength(0, 1000).Error(`话题简介最大长度为1000个字符`),
		),
		validation.Field(&p.CatalogViewType,
			validation.In(CatalogViewTypeColumn, CatalogViewTypeSection).Error("分类视图不正确"),
		),
		validation.Field(&p.TopicHome,
			validation.In(TopicHomeIntroduction, TopicHomeFeed, TopicHomeCataglog, TopicHomeDiscussion, TopicHomeChat).Error("话题首页不正确"),
		),
		validation.Field(&p.EditPermission,
			validation.In(EditPermissionIDCert, EditPermissionWorkCert, EditPermissionIDCertJoined, EditPermissionWorkCertJoined, EditPermissionApprovedIDCertJoined, EditPermissionApprovedWorkCertJoined, EditPermissionAdmin).Error("编辑权限不正确"),
		),
		validation.Field(&p.ViewPermission,
			validation.In(ViewPermissionJoin, ViewPermissionPublic).Error("查看权限不正确"),
		),
		validation.Field(&p.JoinPermission,
			validation.In(JoinPermissionMember, JoinPermissionIDCert, JoinPermissionWorkCert, JoinPermissionMemberApprove, JoinPermissionIDCertApprove, JoinPermissionWorkCertApprove, JoinPermissionAdminAdd, JoinPermissionPurchase, JoinPermissionVIP).Error("加入权限不正确"),
		),
	)
}
