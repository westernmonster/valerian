package model

type SelfProfile struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 手机
	Mobile string `json:"mobile" format:"mobile"`
	// 邮件地址
	Email string `json:"email" format:"email"`

	// 用户名
	UserName string `json:"user_name" format:"user_name"`
	// 性别 1为男， 2为女
	Gender int `json:"gender,omitempty"`
	// 出生年
	BirthYear int `json:"birth_year,omitempty"`
	// 出生月
	BirthMonth int `json:"birth_month,omitempty"`
	// 出生日
	BirthDay int `json:"birth_day,omitempty"`
	// 所在地区值
	Location int64 `json:"location,string,omitempty"`
	// 所在地区名，地区是层级结构，这里将国家、州/省、市、区全部获取出来
	LocationString string `json:"location_string,omitempty"`
	// 自我介绍
	Introduction string `json:"introduction,omitempty"`
	// 头像
	Avatar string `json:"avatar"`
	// 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
	// IP 注册IP
	IP string `json:"ip,omitempty"`

	// 是否身份认证
	IDCert bool `json:"id_cert"`

	// 是否工作认证
	WorkCert bool `json:"work_cert"`

	// 是否机构用户
	IsOrg bool `json:"is_org"`

	// 是否VIP
	IsVIP bool `json:"is_vip"`

	Role string `json:"role"`

	// 状态
	Stat *AccountStatInfo `json:"stat"`

	// 设置
	Settings *SettingInfo `json:"settings"`

	// 注册时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}

type AccountStatInfo struct {

	// 关注数
	FollowingCount int `json:"following_count"`

	// 粉丝数
	FansCount int `json:"fans_count"`

	// 话题数
	TopicCount int `json:"topic_count"`

	// 文章数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`
}

type SettingInfo struct {
	// 动态-赞
	ActivityLike bool `json:"activity_like"`
	// 动态-评论
	ActivityComment bool `json:"activity_comment"`
	// 动态-关注话题
	ActivityFollowTopic bool `json:"activity_follow_topic"`
	// 动态-关注成员
	ActivityFollowMember bool `json:"activity_follow_member"`
	// 通知-赞
	NotifyLike bool `json:"notify_like"`
	// 通知-评论
	NotifyComment bool `json:"notify_comment"`
	// 通知-新粉丝
	NotifyNewFans bool `json:"notify_new_fans"`
	// 通知-新成员
	NotifyNewMember bool `json:"notify_new_member"`
	// Language 语言
	Language string `json:"language"`
}
