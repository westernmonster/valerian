package model

type IDCertResp struct {
	//  状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	Status int `json:"status"`
	// 失败原因
	AuditConclusions *string `json:"audit_conclusions,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty"`
	// 证件号
	IdentificationNumber *string `json:"identification_number,omitempty"`
	// 证件类型, identityCard代表身份证
	IDCardType *string `json:"id_card_type,omitempty"`
	// 证件有效期起始日期
	IDCardStartDate *string `json:"id_card_start_date,omitempty"`
	// 证件有效期截止日期
	IDCardExpiry *string `json:"id_card_expiry,omitempty"`
	// 性别
	Sex *string `json:"sex,omitempty"`
	// 证件照正面图片
	IDCardFrontPic *string `json:"id_card_front_pic,omitempty"`
	// 证件照背面图片
	IDCardBackPic *string `json:"id_card_back_pic,omitempty"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`
}
