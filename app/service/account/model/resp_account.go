package model

type BaseInfo struct {
	ID           int64   `json:"id,string" swaggertype:"string"` //  ID
	UserName     string  `json:"user_name"`                      //  用户名
	Gender       *int    `json:"gender,omitempty"`               //  性别
	Introduction *string `json:"introduction,omitempty"`         //  个人简介
	Avatar       string  `json:"avatar"`                         //  头像
	IDCert       bool    `json:"id_cert"`                        //  是否身份认证
	WorkCert     bool    `json:"work_cert"`                      //  是否工作认证
	IsOrg        bool    `json:"is_org"`                         //  是否机构用户
	IsVIP        bool    `json:"is_vip"`                         //  是否VIP用户
}

type ProfileInfo struct {
	ID           int64   `json:"id,string" swaggertype:"string"` //  ID
	UserName     string  `json:"user_name"`                      //  用户名
	Gender       *int    `json:"gender,omitempty"`               //  性别
	Introduction *string `json:"introduction,omitempty"`         //  个人简介
	// 所在地区值
	Location *int64 `json:"location,string,omitempty"`
	// 所在地区名，地区是层级结构，这里将国家、州/省、市、区全部获取出来
	LocationString *string `json:"location_string,omitempty"`
	Avatar         string  `json:"avatar"`    //  头像
	IDCert         bool    `json:"id_cert"`   //  是否身份认证
	WorkCert       bool    `json:"work_cert"` //  是否工作认证
	IsOrg          bool    `json:"is_org"`    //  是否机构用户
	IsVIP          bool    `json:"is_vip"`    //  是否VIP用户
	Role           string  `json:"role"`
	// 注册时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}
