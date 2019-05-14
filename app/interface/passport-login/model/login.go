package model

import (
	"regexp"
	"valerian/infrastructure/berr"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// 邮件登录请求
type ArgEmailLogin struct {
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
	// Email 邮件地址
	Email string `json:"email"`
	// Password 密码，长度至少为6位
	Password string `json:"password"`
	// ClientID OAUTH2 Client ID
	ClientID string `json:"client_id"`
}

func (p *ArgEmailLogin) Validate() error {
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
type ArgDigitLogin struct {
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

func (p *ArgDigitLogin) Validate() error {
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
type ArgMobileLogin struct {
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

func (p *ArgMobileLogin) Validate() error {
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

// 登录返回结果
type LoginResp struct {
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

	Profile *Profile `json:"profile"`
}

// 用户资料
type Profile struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 手机
	Mobile string `json:"mobile" format:"mobile"`
	// 邮件地址
	Email string `json:"email" format:"email"`

	// 用户名
	UserName string `json:"user_name" format:"user_name"`
	// 性别 1为男， 2为女
	Gender *int `json:"gender,omitempty"`
	// 出生年
	BirthYear *int `json:"birth_year,omitempty"`
	// 出生月
	BirthMonth *int `json:"birth_month,omitempty"`
	// 出生日
	BirthDay *int `json:"birth_day,omitempty"`
	// 所在地区值
	Location *int64 `json:"location,string,omitempty"`
	// 所在地区名，地区是层级结构，这里将国家、州/省、市、区全部获取出来
	LocationString *string `json:"location_string,omitempty"`
	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`
	// 头像
	Avatar string `json:"avatar"`
	// 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
	// IP 注册IP
	IP *string `json:"ip,omitempty"`

	// 是否身份认证
	IDCert bool `json:"id_cert"`

	// 是否工作认证
	WorkCert bool `json:"work_cert"`

	// 是否机构用户
	IsOrg bool `json:"is_org"`

	// 是否VIP
	IsVIP bool `json:"is_vip"`
	// 注册时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}
