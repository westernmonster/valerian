package model

type MemberInfo struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 用户名
	UserName string `json:"user_name" format:"user_name"`
	// 性别 1为男， 2为女
	Gender int32 `json:"gender,omitempty"`

	// 所在地区值
	Location int64 `json:"location,string,omitempty"`

	// 所在地区名，地区是层级结构，这里将国家、州/省、市、区全部获取出来
	LocationString string `json:"location_string,omitempty"`

	// 自我介绍
	Introduction string `json:"introduction,omitempty"`

	// 头像
	Avatar string `json:"avatar"`

	// 是否身份认证
	IDCert bool `json:"id_cert"`

	// 是否工作认证
	WorkCert bool `json:"work_cert"`

	// 是否机构用户
	IsOrg bool `json:"is_org"`

	// 是否VIP
	IsVIP bool `json:"is_vip"`

	Company  string `json:"company"`
	Position string `json:"position"`

	// 状态
	Stat *MemberInfoStat `json:"stat"`
}

type MemberInfoStat struct {
	// 关注数
	FollowingCount int32 `json:"following_count"`

	// 粉丝数
	FansCount int32 `json:"fans_count"  db:"-"`

	// 话题数
	TopicCount int32 `json:"topic_count"`

	// 文章数
	ArticleCount int32 `json:"article_count"`

	// 讨论数
	DiscussionCount int32 `json:"discussion_count"`
}

type MajorListResp struct {
	// 大类话题
	Items []*TargetTopic `json:"items"`
}

type RelatedListResp struct {
	Items []*TargetTopic `json:"items"`
}

type MemberListResp struct {
	Items []*MemberInfo `json:"items"`
}
