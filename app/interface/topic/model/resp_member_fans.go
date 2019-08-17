package model

type TopicMemberFans struct {
	// 账户ID
	AccountID int64 `json:"account_id" swaggertype:"string"`
	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`
	// 头像
	Avatar string `json:"avatar"`

	// 关注数
	FollowCount int `json:"follow_count"`

	// 粉丝数
	FansCount int `json:"fans_count"`

	// 是否当前粉丝
	IsMember bool `json:"is_member"`

	// 是否邀请
	Invited bool `json:"invited"`
}

type TopicMemberFansResp struct {
	Items  []*TopicMemberFans `json:"items"`
	Paging *Paging            `json:"paging"`
}
