package model

import (
	"regexp"
	"valerian/library/ecode"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// 邮件注册请求
type ArgEmail struct {
	// 邮件地址
	Email string `json:"email"`
	// 验证码 6位数字
	Valcode string `json:"valcode"`
	// 密码
	Password string `json:"password"`
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int32 `json:"source"`
	// ClientID OAUTH2 Client ID
	ClientID string `json:"client_id"`
}

func (p *ArgEmail) Validate() error {
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
		validation.Field(&p.ClientID,
			validation.Required.Error(`"client_id" is required`)),
	)
}

// 手机注册请求
type ArgMobile struct {
	// 手机号码
	Mobile string `json:"mobile"`
	// Prefix 电话号码前缀，例如86
	Prefix string `json:"prefix"`
	// 验证码 6位数字
	Valcode string `json:"valcode"`
	// 密码
	Password string `json:"password"`
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int32 `json:"source"`

	// ClientID OAUTH2 Client ID
	ClientID string `json:"client_id"`
}

func (p *ArgMobile) Validate() error {
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
		validation.Field(&p.ClientID,
			validation.Required.Error(`"client_id" is required`)),
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
