package model

import (
	"regexp"
	"valerian/library/ecode"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateIdentity(identityType int32, prefix string) *ValidateIdentityRule {
	return &ValidateIdentityRule{
		IdentityType: identityType,
		Prefix:       prefix,
	}
}

type ValidateIdentityRule struct {
	IdentityType int32
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
	IdentityType int32 `json:"identity_type"`
}

func (p *ArgForgetPassword) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Identity, validation.Required, ValidateIdentity(p.IdentityType, p.Prefix)),
		validation.Field(&p.Valcode, validation.Required, validation.RuneLength(6, 6), is.Digit),
		validation.Field(&p.IdentityType, validation.Required, validation.In(IdentityEmail, IdentityMobile)),
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
		validation.Field(&p.Password, validation.Required),
		validation.Field(&p.SessionID, validation.Required),
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

	// 更新用户名
	UserName *string `json:"user_name,omitempty"`

	// 用户性别， 1 为男，2 为女
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Gender *int32 `json:"gender,omitempty"`

	// 地区
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Location *int64 `json:"location,string,omitempty" swaggertype:"string"`

	// 出生年
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	BirthYear *int32 `db:"birth_year" json:"birth_year,omitempty"`

	// 出生月
	BirthMonth *int32 `db:"birth_month" json:"birth_month,omitempty"`

	// 出生日
	BirthDay *int32 `db:"birth_day" json:"birth_day,omitempty"`

	// 个性签名
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Introduction *string `json:"introduction,omitempty"`

	// 密码
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Password *string `json:"password,omitempty"`
}
