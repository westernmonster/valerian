package models

import (
	"regexp"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// 邮件登录请求
// swagger:model
type EmailLoginReq struct {
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
	// Email 邮件地址
	Email string `json:"email"`
	// Password 密码，长度至少为6位
	Password string `json:"password"`
	// ClientID OAUTH2 Client ID
	ClientID string `json:"client_id"`
}

func (p *EmailLoginReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Email,
			validation.Required.Error(`"email" is required`),
			is.Email.Error(`"email" is not validate`),
		),
		validation.Field(&p.Password,
			validation.Required.Error(`"password" is required`),
			validation.RuneLength(6, 50).Error(`"password" length must in 6-50 characters`)),
		validation.Field(&p.Source,
			validation.Required.Error(`"source" is required`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error(`"source" is not validate`)),
		validation.Field(&p.ClientID,
			validation.Required.Error(`"client_id" is required`)),
	)
}

// 验证码登录
type DigitLoginReq struct {
	// 手机号码
	Mobile string `json:"mobile"`
	// Prefix 国际区号，例如86
	Prefix string `json:"prefix"`
	// 验证码 6位数字
	Valcode string `json:"valcode"`
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`

	// ClientID
	ClientID string `json:"client_id"`
}

func (p *DigitLoginReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Mobile,
			validation.Required.Error(`请输入手机号码`),
			ValidateMobile(p.Prefix),
		),
		validation.Field(&p.Prefix,
			validation.Required.Error(`请输入国际区号`),
		),
		validation.Field(&p.Valcode,
			validation.Required.Error(`请输入验证码`),
			validation.RuneLength(6, 6).Error(`验证码必须为6位数字`),
			is.Digit.Error("验证码必须为6位数字")),
		validation.Field(&p.Source,
			validation.Required.Error(`请输入来源`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error("来源不在允许范围内")),
		validation.Field(&p.ClientID, validation.Required.Error(`"client_id" is required`)),
	)
}

// 手机登录请求
// swagger:model
type MobileLoginReq struct {
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
	// Mobile 手机号码
	Mobile string `json:"mobile"`
	// Prefix 国际区号，例如86
	Prefix string `json:"prefix"`
	// Password 密码
	Password string `json:"password"`
	// ClientID
	ClientID string `json:"client_id"`
}

func (p *MobileLoginReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Mobile,
			validation.Required.Error(`"mobile" is required`),
			ValidateMobile(p.Prefix),
		),
		validation.Field(&p.Prefix,
			validation.Required.Error(`"prefix" is required`),
		),
		validation.Field(&p.Password,
			validation.Required.Error(`"password" password is required`),
			validation.RuneLength(6, 50).Error(`"password" length must in 6-50 characters`)),
		validation.Field(&p.Source,
			validation.Required.Error(`"Source" is required`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error(`"source" is not validate`)),
		validation.Field(&p.ClientID, validation.Required.Error(`"client_id" is required`)),
	)
}

// 登录返回结果
// swagger:model
type LoginResult struct {
	// 用户ID
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
	// 用户角色，用于客户端权限管理
	Role string `json:"role"`
	// Access Token， 请在 HTTP 请求头中添加
	// 例子： Authorization: Bearer  TJVA95OrM7E20RMHrHDcEfxjoYZgeFONFh7HgQ
	AccessToken string `json:"access_token"`
	// 有效期 秒为单位
	ExpiresIn int `json:"expires_in"`
	// Token 类型，默认为 Bearer
	TokenType string `json:"token_type"`
	// Scope 暂不使用
	Scope string `json:"scope"`
	// Refresh Token 暂不使用
	RefreshToken string `json:"refresh_token,omitempty"`

	Profile *ProfileResp `json:"profile"`
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
			return berr.Errorf(`"mobile" format is not validate`)
		}
	} else { // China
		if !otherRegex.MatchString(mobile) {
			return berr.Errorf(`"mobile" format is not validate`)
		}
	} // Other Country

	return nil
}
