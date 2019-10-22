package model

type LoginResp struct {
	// 用户ID
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
	// 用户角色，用于客户端权限管理
	Role string `json:"role"`

	Profile *Profile `json:"profile"`
}

// 用户资料
type Profile struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 邮件地址
	Email string `json:"email" format:"email"`

	// 用户名
	UserName string `json:"user_name" format:"user_name"`

	Role string `json:"role"`
	// 注册时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}
