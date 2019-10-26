package model

type BaseInfo struct {
	ID           int64  `json:"id,string" swaggertype:"string"` //  ID
	UserName     string `json:"user_name"`                      //  用户名
	Gender       int32  `json:"gender,omitempty"`               //  性别
	Introduction string `json:"introduction,omitempty"`         //  个人简介
	Avatar       string `json:"avatar"`                         //  头像
	IDCert       bool   `json:"id_cert"`                        //  是否身份认证
	WorkCert     bool   `json:"work_cert"`                      //  是否工作认证
	IsOrg        bool   `json:"is_org"`                         //  是否机构用户
	IsVIP        bool   `json:"is_vip"`                         //  是否VIP用户
}

type SelfProfile struct {
	ID int64 `json:"id,string" swaggertype:"string"` //  ID

	// 手机前缀
	Prefix string `json:"prefix"`
	// 手机
	Mobile string `json:"mobile" format:"mobile"`
	// 邮件地址
	Email string `json:"email" format:"email"`
	//  用户名
	UserName string `json:"user_name"`
	//  性别
	Gender int32 `json:"gender,omitempty"`
	// 出生年
	BirthYear int32 `json:"birth_year,omitempty"`
	// 出生月
	BirthMonth int32 `json:"birth_month,omitempty"`
	// 出生日
	BirthDay int32 `json:"birth_day,omitempty"`
	//  个人简介
	Introduction string `json:"introduction,omitempty"`
	//  头像
	Avatar string `json:"avatar"`

	// 来源，1:Web, 2:iOS; 3:Android
	Source int32 `json:"source"`

	// 所在地区值
	Location int64 `json:"location,string,omitempty"`
	// 所在地区名，地区是层级结构，这里将国家、州/省、市、区全部获取出来
	LocationString string `json:"location_string,omitempty"`
	//  是否身份认证
	IDCert bool `json:"id_cert"`

	// 状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	IDCertStatus int32 `json:"id_cert_status"`

	//  是否工作认证
	WorkCert bool `json:"work_cert"`

	// IP 注册IP
	IP string `json:"ip,omitempty"`

	// 工作认证状态
	// -1 未认证
	// 0 审核中
	// 1 通过审核
	// 2 审核失败
	WorkCertStatus int32 `json:"work_cert_status"`
	//  是否机构用户
	IsOrg bool `json:"is_org"`
	//  是否VIP用户
	IsVIP bool `json:"is_vip"`
	// 角色
	Role string `json:"role"`
	// 注册时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}

type ProfileInfo struct {
	ID           int64  `json:"id,string" swaggertype:"string"` //  ID
	UserName     string `json:"user_name"`                      //  用户名
	Gender       int32  `json:"gender,omitempty"`               //  性别
	Introduction string `json:"introduction,omitempty"`         //  个人简介

	// 所在地区值
	Location int64 `json:"location,string,omitempty"`
	// 所在地区名，地区是层级结构，这里将国家、州/省、市、区全部获取出来
	LocationString string `json:"location_string,omitempty"`
	Avatar         string `json:"avatar"`    //  头像
	IDCert         bool   `json:"id_cert"`   //  是否身份认证
	WorkCert       bool   `json:"work_cert"` //  是否工作认证
	IsOrg          bool   `json:"is_org"`    //  是否机构用户
	IsVIP          bool   `json:"is_vip"`    //  是否VIP用户
	Role           string `json:"role"`
	// 注册时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}
