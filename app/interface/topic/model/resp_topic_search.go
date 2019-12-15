package model

type CanEditTopicItem struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`

	// 简介
	Introduction string `json:"introduction"`

	// 封面
	Avatar string `json:"avatar,omitempty"`

	// 成员数
	MemberCount int32 `json:"member_count"`

	// 成员数
	ArticleCount int32 `json:"article_count"`

	// 讨论数
	DiscussionCount int32 `json:"discussion_count"`

	// EditPermission 编辑权限
	EditPermission string `json:"edit_permission,omitempty"`

	HasCatalogTaxonomy bool `json:"has_catalog_taxonomy"`
	// 是否有二级目录
}

type CanEditTopicsResp struct {
	Items  []*CanEditTopicItem `json:"items"`
	Paging *Paging             `json:"paging"`
}

type Creator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty"`

	// 自我介绍
	Introduction string `json:"introduction,omitempty"`
}

type JoinedTopicItem struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`

	// 简介
	Introduction string `json:"introduction"`

	// 封面
	Avatar string `json:"avatar,omitempty"`

	// 成员数
	MemberCount int32 `json:"member_count"`

	// 成员数
	ArticleCount int32 `json:"article_count"`

	// 讨论数
	DiscussionCount int32 `json:"discussion_count"`

	// EditPermission 编辑权限
	EditPermission string `json:"edit_permission,omitempty"`
}

type JoinedTopicsResp struct {
	Items  []*JoinedTopicItem `json:"items"`
	Paging *Paging            `json:"paging"`
}
