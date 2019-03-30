package models

import (
	"regexp"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Account 账户
type Account struct {
	ID           int64   `db:"id" json:"id,string"`                        // ID ID
	Mobile       string  `db:"mobile" json:"mobile"`                       // Mobile 手机
	Email        string  `db:"email" json:"email"`                         // Email 邮件地址
	Gender       *int    `db:"gender" json:"gender,omitempty"`             // Gender 性别
	BirthYear    *int    `db:"birth_year" json:"birth_year,omitempty"`     // BirthYear 出生年
	BirthMonth   *int    `db:"birth_month" json:"birth_month,omitempty"`   // BirthMonth 出生月
	BirthDay     *int    `db:"birth_day" json:"birth_day,omitempty"`       // BirthDay 出生日
	Location     *int64  `db:"location" json:"location,string,omitempty"`  // Location 地区
	Introduction *string `db:"introduction" json:"introduction,omitempty"` // Introduction 自我介绍
	Avatar       string  `db:"avatar" json:"avatar"`                       // Avatar 头像
	Source       int     `db:"source" json:"source"`                       // Source 注册来源
	IP           int     `db:"ip" json:"ip"`                               // IP 注册IP
	CreatedAt    int64   `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64   `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type LoginReq struct {
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
	// Identity 登录标识，可以传入邮件或手机号，请在提交前进行验证
	Identity string `json:"identity"`
	// Password 密码，服务端不保存密码的明文值，请在提交前进行 MD5 哈希
	Password string `json:"password"`
}

type LoginResult struct {
	// 用户角色，用于客户端权限管理
	Role string `json:"role"`
	// Token JWT Token， 请在 HTTP 请求头中添加
	// 例子： Authorization: Bearer  TJVA95OrM7E20RMHrHDcEfxjoYZgeFONFh7HgQ
	Token string `json:"token"`
}

func (p *LoginReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Identity,
			validation.Required.Error(`请输入手机号或邮件地址`),
			validation.By(ValidateIdentity)),
		validation.Field(&p.Password,
			validation.Required.Error(`"password" is required`),
			validation.RuneLength(32, 32).Error(`the length of "password" is incorrect`)),
		validation.Field(&p.Source,
			validation.Required.Error(`请输入手机号或邮件地址`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error("来源不在允许范围内")),
	)
}

type ForgetPasswordReq struct {
	Identity string `json:"identity"`
	Valcode  string `json:"valcode"`
}

func (p *ForgetPasswordReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Identity, validation.Required.Error(`请输入手机号或邮件地址`), validation.By(ValidateIdentity)),
		validation.Field(&p.Valcode,
			validation.Required.Error(`请输入验证码`),
			validation.RuneLength(6, 6).Error(`验证码必须为6位数字`),
			is.Digit.Error("验证码必须为6位数字")),
	)
}

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

type RegisterReq struct {
	// 用户标识， 可以为邮件或手机号码
	Identity string `json:"identity"`
	// 验证码 6位数字
	Valcode string `json:"valcode"`
	// 密码 后端不保存明文密码，请于前端求得当前密码MD5哈希值后发送给后端
	Password string `json:"password"`
	// Source 来源，1:Web, 2:iOS; 3:Android
	Source int `json:"source"`
}

func (p *RegisterReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Identity, validation.Required.Error(`请输入手机号或邮件地址`), validation.By(ValidateIdentity)),
		validation.Field(&p.Password,
			validation.Required.Error(`请输入密码`),
			validation.RuneLength(32, 32).Error(`密码格式不正确`)),
		validation.Field(&p.Valcode,
			validation.Required.Error(`请输入验证码`),
			validation.RuneLength(6, 6).Error(`验证码必须为6位数字`),
			is.Digit.Error("验证码必须为6位数字")),
		validation.Field(&p.Source,
			validation.Required.Error(`请输入手机号或邮件地址`),
			validation.In(SourceAndroid, SourceiOS, SourceWeb).Error("来源不在允许范围内")),
	)
}

func ValidateIdentity(obj interface{}) error {
	identity := obj.(string)

	regex := regexp.MustCompile(MobileRegex)

	if !govalidator.IsEmail(identity) && !regex.MatchString(identity) {
		return berr.Errorf("手机号或者邮件地址不正确")
	}

	return nil
}

type RequestValcodeReq struct {
	// 用户标识, 可以为邮件或手机号码
	Identity string `json:"identity"`
	// 验证码类型, 1为注册验证码, 2为重置密码验证码
	CodeType int `json:"code_type"`
}

func (p *RequestValcodeReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Identity, validation.Required.Error(`请输入手机号或邮件地址`), validation.By(ValidateIdentity)),
		validation.Field(&p.CodeType,
			validation.Required.Error(`请输入验证码类型`),
			validation.In(ValcodeRegister, ValcodeForgetPassword).Error("验证码类型不在允许范围内")),
	)
}

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

type UpdateProfileReq struct {
	// 用户头像
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Avatar *string `json:"avatar,omitempty"`

	// 用户性别， 1 为男，2 为女
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Gender *int `json:"gener,omitempty"`

	// 地区
	// 如果无需更改该字段，在提交JSON请求中请不要包含该字段
	Location *int64 `json:"location,string,omitempty"`

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

type IdentityAuthReq struct {
	CertType   int
	CertNumber string
	RealName   string
	Company    string
	Department string
	Position   string
}
