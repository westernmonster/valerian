package model

import "valerian/library/database/sqlx/types"

type MemberItem struct {
	ID int64 `json:"id,string" swaggertype:"string" format:"int64" db:"id"`
	// 自我介绍
	Introduction string `json:"introduction,omitempty" db:"introduction"`
	// 头像
	Avatar string `json:"avatar" db:"avatar"`

	// 关注数
	FollowingCount int `json:"following_count" db:"-"`

	// 粉丝数
	FansCount int `json:"fans_count"  db:"-"`

	// 用户名
	UserName string `json:"user_name" db:"user_name"`

	// 性别 1为男， 2为女
	Gender int `json:"gender,omitempty" db:"gender"`

	// 是否身份认证
	IDCert types.BitBool `json:"id_cert" db:"id_cert"`

	// 是否工作认证
	WorkCert types.BitBool `json:"work_cert" db:"work_cert"`

	// 是否机构用户
	IsOrg types.BitBool `json:"is_org" db:"is_org"`

	// 是否VIP
	IsVIP types.BitBool `json:"is_vip" db:"is_vip"`
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
