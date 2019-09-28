package model

import (
	"regexp"
	"valerian/library/ecode"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// 邮件验证码请求
type ArgEmailValcode struct {
	// 邮件地址
	Email string `json:"email"`

	// 验证码类型, 1为注册验证码, 2为重置密码验证码
	CodeType int `json:"code_type"`
}

func (p *ArgEmailValcode) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Email, validation.Required.Error(`请输入邮件地址`), is.Email.Error("邮件地址格式不正确")),
		validation.Field(&p.CodeType,
			validation.Required.Error(`请输入验证码类型`),
			validation.In(ValcodeRegister, ValcodeForgetPassword).Error("验证码类型不在允许范围内")),
	)
}

// 短信验证码请求
type ArgMobileValcode struct {
	// 手机号码
	Mobile string `json:"mobile"`

	// Prefix 电话号码前缀，例如86
	Prefix string `json:"prefix"`

	// 验证码类型, 1为注册验证码, 2为重置密码验证码, 3为登录验证码
	CodeType int `json:"code_type"`
}

func (p *ArgMobileValcode) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Mobile, validation.Required.Error(`请输入手机号`),
			ValidateMobile(p.Prefix),
		),
		validation.Field(&p.CodeType,
			validation.Required.Error(`请输入验证码类型`),
			validation.In(ValcodeLogin, ValcodeRegister, ValcodeForgetPassword, ValcodeDirect).Error("验证码类型不在允许范围内")),
	)
}

func ValidateMobile(prefix string) *ValidateMobileRule {
	return &ValidateMobileRule{
		Prefix: prefix,
	}
}

type ValidateMobileRule struct {
	Prefix string
}

func (p *ValidateMobileRule) Validate(v interface{}) error {
	mobile := v.(string)

	chinaRegex := regexp.MustCompile(ChinaMobileRegex)
	otherRegex := regexp.MustCompile(OtherMobileRegex)

	if p.Prefix == "86" {
		if !chinaRegex.MatchString(mobile) {
			return ecode.InvalidMobile
		}
	} else { // China
		if !otherRegex.MatchString(mobile) {
			return ecode.InvalidMobile
		}
	} // Other Country

	return nil
}
