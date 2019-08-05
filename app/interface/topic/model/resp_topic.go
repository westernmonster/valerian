package model

type TopicResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 成员
	Members []*TopicMemberResp `json:"members"`

	// 成员数
	MembersCount int `json:"members_count"`

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

	// 分类视图
	// section 章节
	// column 栏目
	CatalogViewType string `json:"catalog_view_type"`

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

	// 消息免打扰
	MuteNotification bool `db:"mute_notification" json:"mute_notification"`

	// 编辑权限
	// "member"
	// "admin"
	EditPermission string `json:"edit_permission"`

	// 加入权限
	// 默认为 member
	// member 注册用户
	// member_approve 用户需审批
	// cert_approve 认证需审批
	// manual_add 仅管理员手工添加
	JoinPermission string `json:"join_permission"`

	// 查看权限
	// public 公开
	// join 加入
	ViewPermission string `json:"view_permission"`

	// 重要标记
	Important bool `db:"important" json:"important"`

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
