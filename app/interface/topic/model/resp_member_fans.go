package model

import "valerian/library/database/sqlx/types"

type FollowItem struct {
	// 账户ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64" db:"id"`
	// 自我介绍
	Introduction *string `json:"introduction,omitempty" db:"introduction"`
	// 头像
	Avatar string `json:"avatar" db:"avatar"`

	// 用户名
	UserName string `json:"user_name" db:"user_name"`

	// 性别 1为男， 2为女
	Gender *int `json:"gender,omitempty" db:"gender"`

	// 是否身份认证
	IDCert types.BitBool `json:"id_cert" db:"id_cert"`

	// 是否工作认证
	WorkCert types.BitBool `json:"work_cert" db:"work_cert"`

	// 是否机构用户
	IsOrg types.BitBool `json:"is_org" db:"is_org"`

	// 是否VIP
	IsVIP types.BitBool `json:"is_vip" db:"is_vip"`

	// 关注数
	FollowCount int `json:"follow_count" db:"-"`

	// 粉丝数
	FansCount int `json:"fans_count"  db:"-"`

	// 是否当前粉丝
	IsMember bool `json:"is_member" db:"-"`

	// 是否邀请
	Invited bool `json:"invited" db:"-"`
}

type TopicMemberFansResp struct {
	Items  []*FollowItem `json:"items"`
	Paging *Paging       `json:"paging"`
}
