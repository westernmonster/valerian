package model

type TopicResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 集合ID
	TopicSetID int64 `json:"topic_set_id,string"  swaggertype:"string"`

	// 成员
	Members []*TopicMemberResp `json:"members"`

	// 成员数
	MembersCount int `json:"members_count"`

	// 顺序
	Seq int `db:"seq" json:"seq"`

	// 资源数量
	ResourcesCount int `json:"resources_count"`

	// 版本列表
	Versions []*TopicVersionResp `json:"versions"`

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
	Catalogs []*TopicLevel1Catalog `json:"catalogs"`

	// 分类视图
	// section 章节
	// column 栏目
	CatalogViewType string `json:"catalog_view_type"`

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

	TopicMeta *TopicMeta `json:"meta,omitempty"`
}

type TopicMeta struct {
	// 是否允许关注
	CanFollow bool `json:"can_follow"`

	// 是否能编辑
	CanEdit bool `json:"can_edit"`

	// 是否能查看
	CanView bool `json:"can_view"`

	// 关注状态
	FollowStatus int `json:"follow_status"`

	IsMember bool `json:"is_member"`

	// 成员角色
	MemberRole string `json:"member_role"`
}

type TopicVersionResp struct {
	// 集合ID
	TopicSetID int64 `db:"topic_set_id" json:"topic_set_id,string"  swaggertype:"string"`
	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`

	// 顺序
	Seq int `db:"seq" json:"seq"`

	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`
	// 话题名
	TopicName string `db:"topic_name" json:"topic_name"`

	// Meta
	TopicMeta *TopicMeta `json:"meta,omitempty"`
}

type RelatedTopicShort struct {
	// 话题ID
	TopicID int64 `db:"topic_id" json:"topic_id,string" swaggertype:"string"`
	// 话题名
	TopicName string `db:"topic_name" json:"topic_name"`

	// 顺序
	Seq int `db:"seq" json:"seq"`

	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`
	// 关联类型
	// normal 普通关联
	// strong 强关联
	Type string `db:"type" json:"type"`
}
