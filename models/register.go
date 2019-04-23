package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// 邮件注册请求
// swagger:model
type EmailRegisterReq struct {
	// 邮件地址
	Email string `json:"email"`
	// 验证码 6位数字
	Valcode string `json:"valcode"`
	// 密码
	Password string `json:"password"`
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
}

func (p *EmailRegisterReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Email, validation.Required.Error(`请输入手机号或邮件地址`),
			is.Email.Error("邮件格式不正确"),
		),
		validation.Field(&p.Password,
			validation.Required.Error(`"password" is required`),
			validation.RuneLength(6, 50).Error(`"password" length must in 6-50 characters`)),
		validation.Field(&p.Valcode,
			validation.Required.Error(`请输入验证码`),
			validation.RuneLength(6, 6).Error(`验证码必须为6位数字`),
			is.Digit.Error("验证码必须为6位数字")),
		validation.Field(&p.Source,
			validation.Required.Error(`请输入来源`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error("来源不在允许范围内")),
	)
}

// 手机注册请求
// swagger:model
type MobileRegisterReq struct {
	// 手机号码
	Mobile string `json:"mobile"`
	// Prefix 电话号码前缀，例如86
	Prefix string `json:"prefix"`
	// 验证码 6位数字
	Valcode string `json:"valcode"`
	// 密码
	Password string `json:"password"`
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
}

func (p *MobileRegisterReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Mobile,
			validation.Required.Error(`请输入手机号码`),
			ValidateMobile(p.Prefix),
		),
		validation.Field(&p.Password,
			validation.Required.Error(`"password" is required`),
			validation.RuneLength(6, 50).Error(`"password" length must in 6-50 characters`)),
		validation.Field(&p.Valcode,
			validation.Required.Error(`请输入验证码`),
			validation.RuneLength(6, 6).Error(`验证码必须为6位数字`),
			is.Digit.Error("验证码必须为6位数字")),
		validation.Field(&p.Source,
			validation.Required.Error(`请输入来源`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error("来源不在允许范围内")),
	)
}
