package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Account 账户
type Account struct {
	ID           int64   `db:"id" json:"id,string"`                        // ID ID
	Mobile       string  `db:"mobile" json:"mobile"`                       // Mobile 手机
	Email        string  `db:"email" json:"email"`                         // Email 邮件地址
	Gender       *int    `db:"gender" json:"gender,omitempty"`             // Gender 性别
	BirthYear    *int    `db:"birth_year" json:"birth_year,omitempty"`     // BirthYear 出生年
	BirthMonth   *int    `db:"birth_month" json:"birth_month,omitempty"`   // BirthMonth 出生月
	BirthDay     *int    `db:"birth_day" json:"birth_day,omitempty"`       // BirthDay 出生日
	Introduction *string `db:"introduction" json:"introduction,omitempty"` // Introduction 自我介绍
	Avatar       string  `db:"avatar" json:"avatar"`                       // Avatar 头像
	Source       int     `db:"source" json:"source"`                       // Source 注册来源
	IP           int     `db:"ip" json:"ip"`                               // IP 注册IP
	CreatedAt    int64   `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64   `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type LoginReq struct {
	// Identity 登录标识，可以传入邮件或手机号，请在提交前进行验证
	Identity string `json:"identity" validate:"required,identity"`
	// Password 密码，服务端不保存密码的明文值，请在提交前进行 MD5 哈希
	Password string `json:"password" validate:"required"`
}

type LoginResult struct {
	// 用户角色，用于客户端权限管理
	Role string `json:"role"`
	// Token JWT Token， 请在 HTTP 请求头中添加
	// 例子： Authorization: Bearer  TJVA95OrM7E20RMHrHDcEfxjoYZgeFONFh7HgQ
	Token string `json:"token"`
}

func (p *LoginReq) Validate() error {
	return validation.ValidateStruct(
		p,
		// validation.Field(&p.Email, validation.Required.Error(`"email" is required`), is.Email.Error(`the format of "email" is incorrect`)),
		validation.Field(&p.Password, validation.Required.Error(`"password" is required`), validation.RuneLength(32, 32).Error(`the length of "password" is incorrect`)),
	)
}

type SendValcodeReq struct {
	// Type 验证码类型，1为注册验证码，2为
	Type     int    `json:"type" validate:"required"`
	Identity string `json:"identity" validate:"required"`
}

type ForgetPasswordReq struct {
	Identity string `json:"identity" validate:"required"`
	Valcode  string `json:"valcode" validate:"required"`
}

type ResetPasswordReq struct {
	Password  string `json:"password" validate:"required"`
	SessionID string `json:"session_id" validate:"required"`
}

type RegisterReq struct {
	Identity string `json:"identity" validate:"required"`
	Valcode  string `json:"valcode" validate:"required"`
}
