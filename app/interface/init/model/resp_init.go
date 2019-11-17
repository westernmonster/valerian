package model

type Creator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name"`
	// 头像
	Avatar string `json:"avatar"`

	// 自我介绍
	Introduction string `json:"introduction,omitempty"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}

type TargetTopic struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`
	// 头像
	// 必须为URL
	Avatar string `json:"avatar"`

	// 成员数
	MemberCount int32 `json:"member_count"`

	// 成员数
	ArticleCount int32 `json:"article_count"`

	// 讨论数
	DiscussionCount int32 `json:"discussion_count"`

	// 简介
	Introduction string `json:"introduction"`

	Creator *Creator `json:"creator,omitempty"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

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
