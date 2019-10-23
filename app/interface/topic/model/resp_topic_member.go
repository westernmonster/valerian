package model

type TopicMemberResp struct {
	// 账户ID
	AccountID int64 `db:"account_id" json:"account_id,string" swaggertype:"string"`
	// 成员角色
	// owner 所有者
	// user 普通用户
	// admin 管理员
	Role string `db:"role" json:"role"`

	// 头像
	Avatar string `db:"avatar" json:"avatar"`
	// 用户名
	UserName string `db:"user_name" json:"user_name"`
}

type TopicMembersPagedResp struct {
	Count    int                `json:"count"`
	PageSize int                `json:"page_size"`
	Data     []*TopicMemberResp `json:"data"`
}

type FollowItem struct {
	// 账户ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64" db:"id"`
	// 自我介绍
	Introduction string `json:"introduction,omitempty" db:"introduction"`
	// 头像
	Avatar string `json:"avatar" db:"avatar"`

	// 用户名
	UserName string `json:"user_name" db:"user_name"`

	// 性别 1为男， 2为女
	Gender int `json:"gender,omitempty" db:"gender"`

	// 是否身份认证
	IDCert bool `json:"id_cert" db:"id_cert"`

	// 是否工作认证
	WorkCert bool `json:"work_cert" db:"work_cert"`

	// 是否机构用户
	IsOrg bool `json:"is_org" db:"is_org"`

	// 是否VIP
	IsVIP bool `json:"is_vip" db:"is_vip"`

	// 关注数
	FollowingCount int `json:"following_count" db:"-"`

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

type Paging struct {
	// 统计数量
	Total int `json:"total,omitempty"`
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
