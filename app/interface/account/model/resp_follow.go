package model

type MemberItem struct {
	ID int64 `json:"id,string" swaggertype:"string" format:"int64" db:"id"`
	// 自我介绍
	Introduction string `json:"introduction,omitempty" db:"introduction"`
	// 头像
	Avatar string `json:"avatar" db:"avatar"`

	// 关注数
	FollowingCount int32 `json:"following_count" db:"-"`

	// 粉丝数
	FansCount int32 `json:"fans_count"  db:"-"`

	// 用户名
	UserName string `json:"user_name" db:"user_name"`

	// 性别 1为男， 2为女
	Gender int32 `json:"gender,omitempty" db:"gender"`

	// 是否身份认证
	IDCert bool `json:"id_cert" db:"id_cert"`

	// 是否工作认证
	WorkCert bool `json:"work_cert" db:"work_cert"`

	// 是否机构用户
	IsOrg bool `json:"is_org" db:"is_org"`

	// 是否VIP
	IsVIP bool `json:"is_vip" db:"is_vip"`
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
