package model

// 登录返回结果
// swagger:model
type LoginResp struct {
	// 用户ID
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
	// 用户角色，用于客户端权限管理
	Role string `json:"role"`
	// Access Token， 请在 HTTP 请求头中添加
	// 例子： Authorization: Bearer  TJVA95OrM7E20RMHrHDcEfxjoYZgeFONFh7HgQ
	AccessToken string `json:"access_token"`
	// 有效期 秒为单位
	ExpiresIn int `json:"expires_in"`
	// Token 类型，默认为 Bearer
	TokenType string `json:"token_type"`
	// Scope 暂不使用
	Scope string `json:"scope"`
	// Refresh Token 暂不使用
	RefreshToken string `json:"refresh_token,omitempty"`

	Profile *Profile `json:"profile"`
}

type Profile struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 手机前缀
	Prefix string `json:"prefix"`
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

	// 状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	IDCertStatus int `json:"id_cert_status"`

	// 是否工作认证
	WorkCert bool `json:"work_cert"`

	// 工作认证状态
	// -1 未认证
	// 0 审核中
	// 1 通过审核
	// 2 审核失败
	WorkCertStatus int `json:"work_cert_status"`

	// 是否机构用户
	IsOrg bool `json:"is_org"`

	// 是否VIP
	IsVIP bool `json:"is_vip"`

	Role string `json:"role"`
	// 注册时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`

	// 状态
	Stat *ProfileStat `json:"stat"`

	// 设置
	Settings *SettingResp `json:"settings"`
}

type ProfileStat struct {

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
