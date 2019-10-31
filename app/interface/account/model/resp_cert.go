package model

type MemberCertInfo struct {
	// 申请人
	Member *MemberInfo `json:"member"`

	// 身份认证信息
	IDCert *IDCertificationInfo `json:"id_cert"`

	// 工作认证信息
	WorkCert *WorkCertificationInfo `json:"work_cert"`
}

type IDCertificationInfo struct {
	// 账户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 姓名
	Name string `json:"name,omitempty"`

	// 证件号
	IdentificationNumber string `json:"identification_number,omitempty"`

	// 证件类型, identityCard代表身份证
	IDCardType string `json:"id_card_type,omitempty"`

	// Status 状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	Status int32 `json:"status"`

	UpdatedAt int64 `json:"updated_at"`
}

type WorkCertificationInfo struct {
	// 账户ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 公司
	Company string `json:"company,omitempty"`

	// 部门
	Department string `json:"department,omitempty"`

	// 职位
	Position string `json:"position,omitempty"`

	// Status 状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	Status int32 `json:"status"`

	UpdatedAt int64 `json:"updated_at"`
}
