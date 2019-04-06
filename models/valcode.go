package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RequestEmailValcodeReq struct {
	// 邮件地址
	Email string `json:"email"`

	// 验证码类型, 1为注册验证码, 2为重置密码验证码
	CodeType int `json:"code_type"`
}

func (p *RequestEmailValcodeReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Email, validation.Required.Error(`请输入邮件地址`), is.Email.Error("邮件地址格式不正确")),
		validation.Field(&p.CodeType,
			validation.Required.Error(`请输入验证码类型`),
			validation.In(ValcodeRegister, ValcodeForgetPassword).Error("验证码类型不在允许范围内")),
	)
}

type RequestMobileValcodeReq struct {
	// 手机号码
	Mobile string `json:"mobile"`

	// Prefix 电话号码前缀，例如86
	Prefix string `json:"prefix"`

	// 验证码类型, 1为注册验证码, 2为重置密码验证码
	CodeType int `json:"code_type"`
}

func (p *RequestMobileValcodeReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Mobile, validation.Required.Error(`请输入手机号或邮件地址`),
			ValidateMobile(p.Prefix),
		),
		validation.Field(&p.CodeType,
			validation.Required.Error(`请输入验证码类型`),
			validation.In(ValcodeRegister, ValcodeForgetPassword).Error("验证码类型不在允许范围内")),
	)
}
