package model

import "valerian/library/database/sqlx/types"

type IDCertification struct {
	ID                   int64         `db:"id" json:"id,string"`                                          // ID ID
	AccountID            int64         `db:"account_id" json:"account_id,string"`                          // AccountID 账户ID
	Status               int           `db:"status" json:"status"`                                         // Status 状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	AuditConclusions     *string       `db:"audit_conclusions" json:"audit_conclusions,omitempty"`         // AuditConclusions 失败原因
	Name                 *string       `db:"name" json:"name,omitempty"`                                   // Name 姓名
	IdentificationNumber *string       `db:"identification_number" json:"identification_number,omitempty"` // IdentificationNumber 证件号
	IDCardType           *string       `db:"id_card_type" json:"id_card_type,omitempty"`                   // IDCardType 证件类型, identityCard代表身份证
	IDCardStartDate      *string       `db:"id_card_start_date" json:"id_card_start_date,omitempty"`       // IDCardStartDate 证件有效期起始日期
	IDCardExpiry         *string       `db:"id_card_expiry" json:"id_card_expiry,omitempty"`               // IDCardExpiry 证件有效期截止日期
	Address              *string       `db:"address" json:"address,omitempty"`                             // Address 地址
	Sex                  *string       `db:"sex" json:"sex,omitempty"`                                     // Sex 性别
	IDCardFrontPic       *string       `db:"id_card_front_pic" json:"id_card_front_pic,omitempty"`         // IDCardFrontPic 证件照正面图片
	IDCardBackPic        *string       `db:"id_card_back_pic" json:"id_card_back_pic,omitempty"`           // IDCardBackPic 证件照背面图片
	FacePic              *string       `db:"face_pic" json:"face_pic,omitempty"`                           // FacePic 认证过程中拍摄的人像正面照图片
	EthnicGroup          *string       `db:"ethnic_group" json:"ethnic_group,omitempty"`                   // EthnicGroup 证件上的民族
	Deleted              types.BitBool `db:"deleted" json:"deleted"`                                       // Deleted 是否删除
	CreatedAt            int64         `db:"created_at" json:"created_at"`                                 // CreatedAt 创建时间
	UpdatedAt            int64         `db:"updated_at" json:"updated_at"`                                 // UpdatedAt 更新时间
}
