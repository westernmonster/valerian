package model

type BasicAccountResp struct {
	// 头像
	Avatar string `db:"avatar" json:"avatar"`
	// 用户名
	UserName string `db:"user_name" json:"user_name"`
	// 账户ID
	AccountID int64 `db:"account_id" json:"account_id,string" swaggertype:"string"`
}
