package model

type ActivitySettingResp struct {
	// 赞
	Like bool `json:"like"`
	// 评论
	Comment bool `json:"comment"`
	// 关注话题
	FollowTopic bool `json:"follow_topic"`
	// 关注成员
	FollowMember bool `json:"follow_member"`
}

type NotifySettingResp struct {
	// 赞
	Like bool `json:"like"`
	// 评论
	Comment bool `json:"comment"`
	// 新粉丝
	NewFans bool `json:"new_fans"`
	// 话题新成员
	NewMember bool `json:"new_member"`
}

type LanguageSettingResp struct {
	Language string `json:"language"`
}

type SettingResp struct {
	// 动态设置
	Activity ActivitySettingResp `json:"activity"`
	// 通知设置
	Notify NotifySettingResp `json:"notify"`
	// 语言设置
	Language LanguageSettingResp `json:"language"`
}
