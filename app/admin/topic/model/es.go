package model

type ESCreator struct {
	// ID ID
	ID int64 `json:"id,string,omitempty"  swaggertype:"string"`
	// 用户名
	UserName *string `json:"user_name,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty"`

	Introduction *string `json:"introduction,omitempty"`
}

type ESTopic struct {
	// ID ID
	ID int64 `json:"id,string,omitempty"  swaggertype:"string"`
	// Name 话题名
	Name *string `json:"name,omitempty" `
	// Avatar 话题头像
	Avatar *string `json:"avatar,omitempty"`
	// Bg 背景图
	Bg *string `json:"bg,omitempty"`
	// Introduction 话题简介
	Introduction *string `json:"introduction,omitempty"`
	// AllowDiscuss 允许讨论
	AllowDiscuss *bool `json:"allow_discuss,omitempty"`
	// AllowChat 开启群聊
	AllowChat *bool `json:"allow_chat,omitempty"`
	// IsPrivate 是否私密
	IsPrivate *bool `json:"is_private,omitempty"`
	// ViewPermission 查看权限
	ViewPermission *string `json:"view_permission,omitempty"`
	// EditPermission 编辑权限
	EditPermission *string `json:"edit_permission,omitempty"`
	// JoinPermission 加入权限
	JoinPermission *string `json:"join_permission,omitempty"`
	// CatalogViewType 分类视图
	CatalogViewType *string `json:"catalog_view_type,omitempty"`
	// CreatedBy 创建人
	Creator *ESCreator `json:"creator,omitempty"`
	// CreatedAt 创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`
	// UpdatedAt 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 成员数
	MemberCount int `json:"member_count"`

	// 成员数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`
}
