package models

// 更新用户资料请求
type UpdateProfileReq struct {
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

// 用户资料
type ProfileResp struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 手机
	Mobile string `json:"mobile" format:"mobile"`
	// 邮件地址
	Email string `json:"email" format:"email"`
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
	// 注册时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}

type AccountSetting struct {
	Locale string `json:"locale"`
}

type AccountCertification struct {
}

type AccountProfile struct {
	// 关注用户数
	Followed int `json:"followed"`
	// 收藏数
	Collection int `json:"collection"`
	// 文章数
	Articles int `json:"articles"`
}
