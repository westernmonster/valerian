package models

import (
	"regexp"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// swagger:model
type EmailLoginReq struct {
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
	// Email 邮件地址
	Email string `json:"email"`
	// Password 密码，服务端不保存密码的明文值，请在提交前进行 MD5 哈希
	Password string `json:"password"`
}

func (p *EmailLoginReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Email,
			validation.Required.Error(`请输入邮件地址`),
			is.Email.Error("邮件地址格式不正确"),
		),
		validation.Field(&p.Password,
			validation.Required.Error(`"password" is required`),
			validation.RuneLength(32, 32).Error(`the length of "password" is incorrect`)),
		validation.Field(&p.Source,
			validation.Required.Error(`请输入手机号或邮件地址`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error("来源不在允许范围内")),
	)
}

// swagger:model
type MobileLoginReq struct {
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
	// Mobile 手机号码
	Mobile string `json:"mobile"`
	// Prefix 电话号码前缀，例如86
	Prefix string `json:"prefix"`
	// Password 密码，服务端不保存密码的明文值，请在提交前进行 MD5 哈希
	Password string `json:"password"`
}

func (p *MobileLoginReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Mobile,
			validation.Required.Error(`请输入手机号码`),
			ValidateMobile(p.Prefix),
		),
		validation.Field(&p.Prefix,
			validation.Required.Error(`请选择国家区号`),
		),
		validation.Field(&p.Password,
			validation.Required.Error(`"password" is required`),
			validation.RuneLength(32, 32).Error(`the length of "password" is incorrect`)),
		validation.Field(&p.Source,
			validation.Required.Error(`请输入手机号或邮件地址`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error("来源不在允许范围内")),
	)
}

// swagger:model
type LoginResult struct {
	// 用户角色，用于客户端权限管理
	Role string `json:"role"`
	// Token JWT Token， 请在 HTTP 请求头中添加
	// 例子： Authorization: Bearer  TJVA95OrM7E20RMHrHDcEfxjoYZgeFONFh7HgQ
	Token string `json:"token"`
}

func ValidateMobile(prefix string) *ValidateIdentityRule {
	return &ValidateIdentityRule{
		Prefix: prefix,
	}
}

type ValidateMobileRule struct {
	IdentityType int
	Prefix       string
}

func (p *ValidateMobileRule) Validate(v interface{}) error {
	mobile := v.(string)

	chinaRegex := regexp.MustCompile(ChinaMobileRegex)
	otherRegex := regexp.MustCompile(OtherMobileRegex)

	if p.Prefix == "86" {
		if !chinaRegex.MatchString(mobile) {
			return berr.Errorf("手机号码不正确")
		}
	} else { // China
		if !otherRegex.MatchString(mobile) {
			return berr.Errorf("手机号码不正确")
		}
	} // Other Country

	return nil
}
