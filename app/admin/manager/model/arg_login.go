package model

import validation "github.com/go-ozzo/ozzo-validation"

// 邮件登录请求
type ArgAdminLogin struct {
	// 用户名登录
	UserName string `json:"user_name"`
	// Password 密码，长度至少为6位
	Password string `json:"password"`
}

func (p *ArgAdminLogin) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.UserName, validation.Required),
		validation.Field(&p.Password, validation.Required, validation.RuneLength(6, 50)),
	)
}
