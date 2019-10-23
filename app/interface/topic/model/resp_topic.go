package model

type TopicResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 成员
	Members []*TopicMemberResp `json:"members,omitempty"`

	// 成员数
	MemberCount int `json:"member_count,omitempty"`

	// 目录
	Catalogs []*TopicLevel1Catalog `json:"catalogs,omitempty"`

	// 授权话题
	AuthTopics []*AuthTopicResp `json:"auth_topics,omitempty"`

	// 讨论分类
	DiscussCategories []*DiscussCategoryResp `json:"discuss_categories,omitempty"`

	// 头像
	// 必须为URL
	Avatar string `json:"avatar,omitempty" extensions:"x-nullable,x-abc=def"`

	// 背景图
	// 必须为URL
	Bg string `json:"bg"`

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
	// discussion 讨论
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

	// 是否收藏
	Fav bool `json:"fav"`

	// 是否能查看
	CanView bool `json:"can_view"`

	// 关注状态
	FollowStatus int `json:"follow_status"`

	IsMember bool `json:"is_member"`

	// 成员角色
	MemberRole string `json:"member_role"`
}

type TopicItem struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`

	// 简介
	Introduction string `json:"introduction"`

	// 封面
	Avatar string `json:"avatar,omitempty" extensions:"x-nullable,x-abc=def"`

	// 成员数
	MemberCount int `json:"member_count"`

	// 文章数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`
}

type TopicListResp struct {
	Items  []*TopicItem `json:"items"`
	Paging *Paging      `json:"paging"`
}
