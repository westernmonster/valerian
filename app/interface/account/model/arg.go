package model

import (
	"regexp"
	"valerian/library/ecode"

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
			return ecode.InvalidEmail
		}
	} else {
		chinaRegex := regexp.MustCompile(ChinaMobileRegex)
		otherRegex := regexp.MustCompile(OtherMobileRegex)

		if p.Prefix == "86" {
			if !chinaRegex.MatchString(identity) {
				return ecode.InvalidMobile
			}
		} else { // China
			if !otherRegex.MatchString(identity) {
				return ecode.InvalidMobile
			}
		} // Other Country
	}

	return nil
}

// 忘记密码请求
// swagger:model
type ArgForgetPassword struct {
	Identity string `json:"identity"`
	Valcode  string `json:"valcode"`
	Prefix   string `json:"prefix"`
	// 标识类型, 1手机, 2邮件
	IdentityType int `json:"identity_type"`
}

func (p *ArgForgetPassword) Validate() error {
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
type ArgResetPassword struct {
	Password  string `json:"password"`
	SessionID string `json:"session_id"`
}

func (p *ArgResetPassword) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Password,
			validation.Required.Error(`请输入密码`)),
		validation.Field(&p.SessionID,
			validation.Required.Error(`请输入 Session ID`)),
	)
}

// 更改密码请求
// swagger:model
type ArgChangePassword struct {
	Password string `json:"password"`
}

func (p *ArgChangePassword) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Password, validation.Required),
	)
}

type ArgUpdateProfile struct {
	// 用户头像
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Avatar *string `json:"avatar,omitempty"`

	// 用户性别， 1 为男，2 为女
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Gender *int `json:"gener,omitempty"`

	// 地区
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Location *int64 `json:"location,string,omitempty" swaggertype:"string"`

	// 出生年
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	BirthYear *int `db:"birth_year" json:"birth_year,omitempty"`

	// 出生月
	BirthMonth *int `db:"birth_month" json:"birth_month,omitempty"`

	// 出生日
	BirthDay *int `db:"birth_day" json:"birth_day,omitempty"`

	// 个性签名
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Introduction *string `json:"introduction,omitempty"`

	// 密码
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Password *string `json:"password,omitempty"`
}

type ArgFollow struct {
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
}

func (p *ArgFollow) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}
