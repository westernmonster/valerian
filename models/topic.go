package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type TopicMember struct {
	// 账户ID
	AccountID int64 `db:"account_id" json:"account_id,string" swaggertype:"string"`
	// 成员角色
	// owner 所有者
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

type RelatedTopicShort struct {
	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`
	// 话题名
	Name string `db:"name" json:"name"`
	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`
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
}

type Topic struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 集合ID
	TopicSetID int64 `json:"topic_set_id,string"  swaggertype:"string"`

	// 成员
	Members []*TopicMember `json:"members"`

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

	// 关联话题
	RelatedTopics []*RelatedTopicShort `json:"related_topics"`

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

	// 是否私密
	IsPrivate bool `json:"is_private"`

	// 开启群聊
	AllowChat bool `json:"allow_chat"`

	// 允许讨论
	AllowDiscuss bool `json:"allow_discuss"`

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

	CreatedAt int64 `json:"created_at" swaggertype:"integer"`

	// 是否关注
	IsFollowed bool `json:"is_followed"`

	// 能否关注
	// 根据当前用户匹配权限返回
	CanFollow bool `json:"can_follow"`
}

type CreateTopicMemberArg struct {
	// 成员ID
	AccountID int64 `json:"account_id,string"`
	// 角色
	// user 普通用户
	// admin 管理员
	Role string `json:"role"`
}

type CreateTopicReq struct {

	// 集合ID
	// 如果是一个全新的创建，无需传入这个字段
	// 在当前话题版本上创建新的版本时候，需要将这个ID传入进来，
	// 后台方可在内部将不同版本归属于一个集合
	TopicSetID *int64 `json:"topic_set_id,string,omitempty"`

	// 成员
	Members []*CreateTopicMemberArg `json:"members,omitempty"`

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

	// 是否私密
	IsPrivate bool `json:"is_private"`

	// 开启群聊
	AllowChat bool `json:"allow_chat"`

	// 允许讨论
	AllowDiscuss bool `json:"allow_discuss"`

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

func (p *CreateTopicReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Members, validation.Required.Error(`请选择成员`)),
		validation.Field(&p.TopicType, validation.Required.Error(`请选择话题类型`)),
		validation.Field(&p.Cover,
			is.URL.Error("封面图格式不正确"),
		),
		validation.Field(&p.Bg,
			is.URL.Error("背景图格式不正确"),
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
			validation.In(EditPermissionIDCert, EditPermissionWorkCert, EditPermissionIDCertJoined, EditPermissionWorkCertJoined, EditPermissionApprovedIDCertJoined, EditPermissionApprovedWorkCertJoined, EditPermissionAdmin).Error("编辑权限不正确"),
		),
		validation.Field(&p.ViewPermission,
			validation.Required.Error(`请输入查看权限`),
			validation.In(ViewPermissionJoin, ViewPermissionPublic).Error("查看权限不正确"),
		),
		validation.Field(&p.JoinPermission,
			validation.Required.Error(`请输入加入权限`),
			validation.In(JoinPermissionMember, JoinPermissionIDCert, JoinPermissionWorkCert, JoinPermissionMemberApprove, JoinPermissionIDCertApprove, JoinPermissionWorkCertApprove, JoinPermissionAdminAdd, JoinPermissionPurchase, JoinPermissionVIP).Error("加入权限不正确"),
		),
		validation.Field(&p.VersionName,
			validation.Required.Error(`请输入版本名`),
			validation.RuneLength(0, 250).Error(`版本名最大长度为250个字符`),
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
	// owner 主理人
	Role string `json:"role"`

	// 操作
	// C U D
	Opt string `json:"opt"`
}

func (p *TopicMemberReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required.Error(`请传入成员ID`)),
		validation.Field(&p.Role, validation.Required.Error(`请传入成员角色`),
			validation.In(MemberRoleUser, MemberRoleAdmin).Error("角色值不正确")),
		validation.Field(&p.Opt, validation.Required.Error(`请传入操作`),
			validation.In("C", "U", "D").Error("操作值不正确")),
	)
}

type BulkSaveTopicMembersReq struct {
	// 成员
	Members []*TopicMemberReq `json:"members"`
}

type UpdateTopicReq struct {
	// 成员
	Members []*CreateTopicMemberArg `json:"members,omitempty"`

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
		validation.Field(&p.CategoryViewType,
			validation.In(CategoryViewTypeiColumn, CategoryViewTypeiSection).Error("分类视图不正确"),
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
		validation.Field(&p.VersionName,
			validation.RuneLength(0, 250).Error(`版本名最大长度为250个字符`),
		),
	)
}

type TopicMembersPagedResp struct {
	Count    int            `json:"count"`
	PageSize int            `json:"page_size"`
	Data     []*TopicMember `json:"data"`
}

type BatchSavedTopicMemberReq struct {
	// 成员列表
	Members []*TopicMemberReq `json:"members"`
}

func (p *BatchSavedTopicMemberReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Members, validation.Required.Error(`请传入成员`)),
	)
}

type TopicFollower struct {
	// 是否关注
	IsFollowed bool `json:"is_followed"`
	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`
}

func (p *TopicFollower) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required.Error(`请传入话题ID`)),
	)
}

type TopicVersion struct {
	// 集合ID
	TopicSetID int64 `db:"topic_set_id" json:"topic_set_id,string"  swaggertype:"string"`
	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`
	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`
}

type ChangeOwnerArg struct {
	NewOwnerID int64 `json:"new_owner_id,string"  swaggertype:"string"`
}

func (p *ChangeOwnerArg) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.NewOwnerID, validation.Required.Error(`请传入新主理人ID`)),
	)
}

type TopicType struct {
	ID   int    `db:"id" json:"id"`     // ID ID
	Name string `db:"name" json:"name"` // Name 话题类型
}
