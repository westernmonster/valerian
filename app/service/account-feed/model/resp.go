package model

type SettingResp struct {
	ActivityLike         bool   `json:"activity_like"`          // ActivityLike 动态-赞
	ActivityComment      bool   `json:"activity_comment"`       // ActivityComment 动态-评论
	ActivityFollowTopic  bool   `json:"activity_follow_topic"`  // ActivityFollowTopic 动态-关注话题
	ActivityFollowMember bool   `json:"activity_follow_member"` // ActivityFollowMember 动态-关注成员
	NotifyLike           bool   `json:"notify_like"`            // NotifyLike 通知-赞
	NotifyComment        bool   `json:"notify_comment"`         // NotifyComment 通知-评论
	NotifyNewFans        bool   `json:"notify_new_fans"`        // NotifyNewFans 通知-新粉丝
	NotifyNewMember      bool   `json:"notify_new_member"`      // NotifyNewMember 通知-新成员
	Language             string `json:"language"`               // Language 语言
}
