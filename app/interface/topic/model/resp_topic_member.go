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
