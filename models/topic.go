package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type TopicMember struct {
	// 账户ID
	AccountID int64 `db:"account_id" json:"account_id,string" swaggertype:"string"`
	// 成员角色
	// user 普通用户
	// admin 管理员
	Role string `db:"role" json:"role"`
	// 头像
	Avatar string `db:"avatar" json:"avatar"`
	// 用户名
	UserName string `db:"user_name" json:"user_name"`
}

type RelatedTopic struct {
	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`
	// 话题名
	Name string `db:"name" json:"name"`
	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`
	// 版本语言
	VersionLanguage string `db:"version_lang" json:"version_lang"`
	// 关联类型
	// normal 普通关联
	// strong 强关联
	Type string `db:"type" json:"type"`
}

type TopicSearchResult struct {
	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`
	// 话题名
	Name string `db:"name" json:"name"`
	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`
	// 版本语言
	VersionLanguage string `db:"version_lang" json:"version_lang"`
}

type Topic struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 集合ID
	TopicSetID int64 `json:"topic_set_id,string"  swaggertype:"string"`

	// 成员
	Members []*TopicMember `json:"members"`

	// 封面图
	// 必须为URL
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
	// section 章节
	// column 栏目
	CategoryViewType string `json:"category_view_type"`

	// 话题类型
	TopicType int `json:"topic_type"`

	TopicTypeName string `json:"topic_type_name"`

	// 话题首页
	// introduction 简介
	// feed 动态
	// catalog 目录
	// discussion 讨论
	// chat 群聊
	TopicHome string `json:"topic_home"`

	// 版本名
	VersionName string `json:"version_name"`

	// 版本语言
	VersionLanguage string `json:"version_lang"`

	// 是否私密
	IsPrivate bool `json:"is_private"`

	// 开启群聊
	AllowChat bool `json:"allow_chat"`

	// 编辑权限
	// auth 认证用户
	// auth_join 认证且加入
	// auth_join_audit 认证且加入，需审核
	// admin 仅管理员
	EditPermission string `json:"edit_permission"`

	// 查看权限
	// public 公开
	// join 加入
	ViewPermission string `json:"view_permission"`

	// 加入权限
	// free  自由加入
	// auth_free 认证用户自由加入
	// approve 经批准
	// auth_approve 认证用户经批准
	// admin 仅管理员添加
	// purchase 购买
	JoinPermission string `json:"join_permission"`

	// 重要标记
	Important bool `db:"important" json:"important"`

	// 消息免打扰
	MuteNotification bool `db:"mute_notification" json:"mute_notification"`

	CreatedAt int64 `json:"created_at" swaggertype:"integer"`
}

type CreateTopicReq struct {

	// 集合ID
	// 如果是一个全新的创建，无需传入这个字段
	// 在当前话题版本上创建新的版本时候，需要将这个ID传入进来，
	// 后台方可在内部将不同版本归属于一个集合
	TopicSetID *int64 `json:"topic_set_id,string,omitempty"`

	// 成员
	Members []*TopicMemberReq `json:"members"`

	// 封面图
	// 必须为URL
	Cover string `json:"cover"`

	// 名称
	Name string `json:"name"`

	// 简介
	Introduction string `json:"introduction"`

	// 关联话题
	RelatedTopics []*RelatedTopicReq `json:"related_topics"`

	// 话题分类
	Categories []*TopicCategoryParentItem `json:"categories"`

	// 分类视图
	// section 章节
	// column 栏目
	CategoryViewType string `json:"category_view_type"`

	// 话题类型
	TopicType int `json:"topic_type"`

	// 话题首页
	// introduction 简介
	// feed 动态
	// catalog 目录
	// discussion 讨论
	// chat 群聊
	TopicHome string `json:"topic_home"`

	// 版本名
	VersionName string `json:"version_name"`

	// 版本语言
	VersionLanguage string `json:"version_lang"`

	// 是否私密
	IsPrivate bool `json:"is_private"`

	// 开启群聊
	AllowChat bool `json:"allow_chat"`

	// 编辑权限
	// auth 认证用户
	// auth_join 认证且加入
	// auth_join_audit 认证且加入，需审核
	// admin 仅管理员
	EditPermission string `json:"edit_permission"`

	// 查看权限
	// public 公开
	// join 加入
	ViewPermission string `json:"view_permission"`

	// 加入权限
	// free  自由加入
	// auth_free 认证用户自由加入
	// approve 经批准
	// auth_approve 认证用户经批准
	// admin 仅管理员添加
	// purchase 购买
	JoinPermission string `json:"join_permission"`

	// 重要标记
	Important bool `db:"important" json:"important"`

	// 消息免打扰
	MuteNotification bool `db:"mute_notification" json:"mute_notification"`
}

func (p *CreateTopicReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Members, validation.Required.Error(`请选择成员`)),
		validation.Field(&p.TopicType, validation.Required.Error(`请选择话题类型`)),
		validation.Field(&p.Cover,
			validation.Required.Error(`请选择封面图`),
			is.URL.Error("封面图格式不正确"),
		),
		validation.Field(&p.Name,
			validation.Required.Error(`请输入话题名`),
			validation.RuneLength(0, 250).Error(`话题名最大长度为250个字符`),
		),
		// TODO: Web安全性
		validation.Field(&p.Introduction,
			validation.Required.Error(`请输入话题简介`),
			validation.RuneLength(0, 1000).Error(`话题简介最大长度为1000个字符`),
		),
		validation.Field(&p.CategoryViewType,
			validation.Required.Error(`请输入分类视图`),
			validation.In(CategoryViewTypeiColumn, CategoryViewTypeiSection).Error("分类视图不正确"),
		),
		validation.Field(&p.TopicHome,
			validation.Required.Error(`请输入话题首页`),
			validation.In(TopicHomeIntroduction, TopicHomeFeed, TopicHomeCataglog, TopicHomeDiscussion, TopicHomeChat).Error("话题首页不正确"),
		),
		validation.Field(&p.EditPermission,
			validation.Required.Error(`请输入编辑权限`),
			validation.In(EditPermissionAuth, EditPermissionAdmin, EditPermissionAuthJoin, EditPermissionAuthJoinAudit).Error("编辑权限不正确"),
		),
		validation.Field(&p.ViewPermission,
			validation.Required.Error(`请输入查看权限`),
			validation.In(ViewPermissionJoin, ViewPermissionPublic).Error("查看权限不正确"),
		),
		validation.Field(&p.JoinPermission,
			validation.Required.Error(`请输入加入权限`),
			validation.In(JoinPermissionFree, JoinPermissionAdmin, JoinPermissionApprove, JoinPermissionAuthFree, JoinPermissionPurchase, JoinPermissionAuthApprove).Error("加入权限不正确"),
		),
		validation.Field(&p.VersionName,
			validation.Required.Error(`请输入版本名`),
			validation.RuneLength(0, 250).Error(`版本名最大长度为250个字符`),
		),
		validation.Field(&p.VersionLanguage,
			validation.Required.Error(`请输入版本语言`),
			validation.RuneLength(0, 50).Error(`版本语言最大长度为50个字符`),
		),
	)
}

type RelatedTopicReq struct {
	// 关联话题ID
	TopicID int64 `json:"topic_id,string"`

	// 类型
	// normal
	// strong
	Type string `json:"type"`
}

func (p *RelatedTopicReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required.Error(`请传入关联话题ID`)),
		validation.Field(&p.Type, validation.Required.Error(`请传入关联话题类型`),
			validation.In(TopicRelationStrong, TopicRelationNormal).Error("话题类型"),
		),
	)
}

type TopicMemberReq struct {
	// 成员ID
	AccountID int64 `json:"account_id,string"`
	// 角色
	// user 普通用户
	// admin 管理员
	Role string `json:"role"`
}

func (p *TopicMemberReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required.Error(`请传入成员ID`)),
		validation.Field(&p.Role, validation.Required.Error(`请传入成员角色`),
			validation.In(MemberRoleUser, MemberRoleAdmin).Error("角色值不正确")),
	)
}

type BulkSaveTopicMembersReq struct {
	// 成员
	Members []*TopicMemberReq `json:"members"`
}

type UpdateTopicReq struct {
	// 成员
	Members []*TopicMemberReq `json:"members,omitempty"`

	// 封面图
	// 必须为URL
	Cover *string `json:"cover,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 简介
	Introduction *string `json:"introduction,omitempty"`

	// 关联话题
	RelatedTopics []*RelatedTopicReq `json:"related_topics,omitempty"`

	// 话题分类
	Categories []*TopicCategoryParentItem `json:"categories,omitempty"`

	// 分类视图
	// section 章节
	// column 栏目
	CategoryViewType *string `json:"category_view_type,omitempty"`

	// 话题类型
	TopicType *int `json:"topic_type,omitempty"`

	// 话题首页
	// introduction 简介
	// feed 动态
	// catalog 目录
	// discussion 讨论
	// chat 群聊
	TopicHome *string `json:"topic_home,omitempty"`

	// 版本名
	VersionName *string `json:"version_name,omitempty"`

	// 版本语言
	VersionLanguage *string `json:"version_lang,omitempty"`

	// 是否私密
	IsPrivate *bool `json:"is_private,omitempty"`

	// 开启群聊
	AllowChat *bool `json:"allow_chat,omitempty"`

	// 编辑权限
	// auth 认证用户
	// auth_join 认证且加入
	// auth_join_audit 认证且加入，需审核
	// admin 仅管理员
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

func (p *UpdateTopicReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Cover,
			is.URL.Error("封面图格式不正确"),
		),
		validation.Field(&p.Name,
			validation.RuneLength(0, 250).Error(`话题名最大长度为250个字符`),
		),
		// TODO: Web安全性
		validation.Field(&p.Introduction,
			validation.RuneLength(0, 1000).Error(`话题简介最大长度为1000个字符`),
		),
		validation.Field(&p.CategoryViewType,
			validation.In(CategoryViewTypeiColumn, CategoryViewTypeiSection).Error("分类视图不正确"),
		),
		validation.Field(&p.TopicHome,
			validation.In(TopicHomeIntroduction, TopicHomeFeed, TopicHomeCataglog, TopicHomeDiscussion, TopicHomeChat).Error("话题首页不正确"),
		),
		validation.Field(&p.EditPermission,
			validation.In(EditPermissionAuth, EditPermissionAdmin, EditPermissionAuthJoin, EditPermissionAuthJoinAudit).Error("编辑权限不正确"),
		),
		validation.Field(&p.ViewPermission,
			validation.In(ViewPermissionJoin, ViewPermissionPublic).Error("查看权限不正确"),
		),
		validation.Field(&p.JoinPermission,
			validation.In(JoinPermissionFree, JoinPermissionAdmin, JoinPermissionApprove, JoinPermissionAuthFree, JoinPermissionPurchase, JoinPermissionAuthApprove).Error("加入权限不正确"),
		),
		validation.Field(&p.VersionName,
			validation.RuneLength(0, 250).Error(`版本名最大长度为250个字符`),
		),
		validation.Field(&p.VersionLanguage,
			validation.RuneLength(0, 50).Error(`版本语言最大长度为50个字符`),
		),
	)
}
