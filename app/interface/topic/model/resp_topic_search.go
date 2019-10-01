package model

type TopicSearchItem struct {
	// ID ID
	ID *int64 `json:"id,string,omitempty"  swaggertype:"string"`
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
	CreatedBy *Creator `json:"created_by,omitempty"`
	// CreatedAt 创建时间
	CreatedAt *int64 `json:"created_at,string"  swaggertype:"string"`
	// UpdatedAt 更新时间
	UpdatedAt *int64 `json:"updated_at,string"  swaggertype:"string"`

	// 成员数
	MemberCount int `json:"member_count"`

	// 是否已经授权
	IsAuthed bool `json:"is_authed"`
}

type TopicSearchResp struct {
	Items  []*TopicSearchItem `json:"items"`
	Paging *Paging            `json:"paging"`
}

type Creator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name"`
	// 头像
	Avatar string `json:"avatar"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`
}
