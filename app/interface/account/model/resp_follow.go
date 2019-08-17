package model

type MemberItem struct {
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`
	// 头像
	Avatar string `json:"avatar"`

	// 关注数
	FollowCount int `json:"follow_count"`

	// 粉丝数
	FansCount int `json:"fans_count"`

	// 用户名
	UserName string `json:"user_name"`

	// 性别 1为男， 2为女
	Gender *int `json:"gender,omitempty"`

	// 是否身份认证
	IDCert bool `json:"id_cert"`

	// 是否工作认证
	WorkCert bool `json:"work_cert"`

	// 是否机构用户
	IsOrg bool `json:"is_org"`

	// 是否VIP
	IsVIP bool `json:"is_vip"`
}

type MemberResp struct {
	Items  []*MemberItem `json:"items"`
	Paging *Paging       `json:"paging"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
