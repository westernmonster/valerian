package model

type MemberTopic struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`
	// 头像
	// 必须为URL
	Avatar *string `json:"avatar"`
	// 成员数
	MembersCount int `json:"members_count"`

	// 简介
	Introduction string `json:"introduction"`
}

type MemberTopicResp struct {
	Items  []*MemberTopic `json:"items"`
	Paging *Paging        `json:"paging"`
}
