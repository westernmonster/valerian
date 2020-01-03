package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgAdminLockAccount struct {
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
}

func (p *ArgAdminLockAccount) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}

type ArgAdminUpdateProfile struct {
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
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

func (p *ArgAdminUpdateProfile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}
