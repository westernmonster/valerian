package models

import (
	"regexp"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateIdentity(identityType int, prefix string) *ValidateIdentityRule {
	return &ValidateIdentityRule{
		IdentityType: identityType,
		Prefix:       prefix,
	}
}

type ValidateIdentityRule struct {
	IdentityType int
	Prefix       string
}

func (p *ValidateIdentityRule) Validate(v interface{}) error {
	identity := v.(string)

	if p.IdentityType == IdentityEmail {
		if !govalidator.IsEmail(identity) {
			return berr.Errorf("邮件地址不正确")
		}
	} else {
		chinaRegex := regexp.MustCompile(ChinaMobileRegex)
		otherRegex := regexp.MustCompile(OtherMobileRegex)

		if p.Prefix == "86" {
			if !chinaRegex.MatchString(identity) {
				return berr.Errorf("手机号码不正确")
			}
		} else { // China
			if !otherRegex.MatchString(identity) {
				return berr.Errorf("手机号码不正确")
			}
		} // Other Country
	}

	return nil
}

// 忘记密码请求
// swagger:model
type ForgetPasswordReq struct {
	Identity string `json:"identity"`
	Valcode  string `json:"valcode"`
	Prefix   string `json:"prefix"`
	// 标识类型, 1手机, 2邮件
	IdentityType int `json:"identity_type"`
}

func (p *ForgetPasswordReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Identity,
			validation.Required.Error(`请输入手机号或邮件地址`),
			ValidateIdentity(p.IdentityType, p.Prefix)),
		validation.Field(&p.Valcode,
			validation.Required.Error(`请输入验证码`),
			validation.RuneLength(6, 6).Error(`验证码必须为6位数字`),
			is.Digit.Error("验证码必须为6位数字")),
		validation.Field(&p.IdentityType,
			validation.Required.Error(`请输入类型`),
			validation.In(IdentityEmail, IdentityMobile).Error("登录标识类型不正确")),
	)
}

// 重设密码请求
// swagger:model
type ResetPasswordReq struct {
	Password  string `json:"password"`
	SessionID string `json:"session_id"`
}

func (p *ResetPasswordReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Password,
			validation.Required.Error(`请输入密码`),
			validation.RuneLength(32, 32).Error(`密码格式不正确`)),
		validation.Field(&p.SessionID,
			validation.Required.Error(`请输入 Session ID`)),
	)
}

// 更改密码请求
// swagger:model
type ChangePasswordReq struct {
	Password string `json:"password"`
}

func (p *ChangePasswordReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Password,
			validation.Required.Error(`请输入密码`),
			validation.RuneLength(32, 32).Error(`密码格式不正确`)),
	)
}
