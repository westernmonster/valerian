package model

type ArgActivitySetting struct {
	// 赞
	Like *bool `json:"like"`
	// 评论
	Comment *bool `json:"comment"`
	// 关注话题
	FollowTopic *bool `json:"follow_topic"`
	// 关注成员
	FollowMember *bool `json:"follow_member"`
}

type ArgNotifySetting struct {
	// 赞
	Like *bool `json:"like"`
	// 评论
	Comment *bool `json:"comment"`
	// 新粉丝
	NewFans *bool `json:"new_fans"`
	// 话题新成员
	NewMember *bool `json:"new_member"`
}

type ArgLanguageSetting struct {
	// 系统语言
	Language *string `json:"language"`
}

type ArgSetting struct {
	// 系统语言
	Language *string `json:"language"`

	// 动态设置
	Activity *ArgActivitySetting `json:"activity,omitempty"`

	// 通知设置
	Notify *ArgNotifySetting `json:"notify,omitempty"`
}

// 注销账户
// swagger:model
type ArgAnnulAccount struct {
	Valcode string `json:"valcode"`
}
