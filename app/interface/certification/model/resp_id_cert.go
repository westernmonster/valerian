package model

type IDCertResp struct {
	Status               int     `json:"status"`                          // Status 状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	AuditConclusions     *string `json:"audit_conclusions,omitempty"`     // AuditConclusions 失败原因
	Name                 *string `json:"name,omitempty"`                  // Name 姓名
	IdentificationNumber *string `json:"identification_number,omitempty"` // IdentificationNumber 证件号
	IDCardType           *string `json:"id_card_type,omitempty"`          // IDCardType 证件类型, identityCard代表身份证
	IDCardStartDate      *string `json:"id_card_start_date,omitempty"`    // IDCardStartDate 证件有效期起始日期
	IDCardExpiry         *string `json:"id_card_expiry,omitempty"`        // IDCardExpiry 证件有效期截止日期
	Sex                  *string `json:"sex,omitempty"`                   // Sex 性别
	IDCardFrontPic       *string `json:"id_card_front_pic,omitempty"`     // IDCardFrontPic 证件照正面图片
	IDCardBackPic        *string `json:"id_card_back_pic,omitempty"`      // IDCardBackPic 证件照背面图片
	CreatedAt            int64   `json:"created_at"`                      // CreatedAt 创建时间
	UpdatedAt            int64   `json:"updated_at"`                      // UpdatedAt 更新时间
}
