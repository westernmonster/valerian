package model

type MemberTopic struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`
	// 头像
	// 必须为URL
	Avatar string `json:"avatar"`
	// 成员数
	MemberCount int32 `json:"member_count"`

	// 简介
	Introduction string `json:"introduction"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type MemberTopicResp struct {
	Items  []*TargetTopic `json:"items"`
	Paging *Paging        `json:"paging"`
}
