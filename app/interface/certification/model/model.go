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

type Account struct {
	ID           int64         `db:"id" json:"id,string"`                        // ID ID
	Mobile       string        `db:"mobile" json:"mobile"`                       // Mobile 手机
	Email        string        `db:"email" json:"email"`                         // Email 邮件地址
	UserName     string        `db:"user_name" json:"user_name"`                 // UserName 用户名
	Password     string        `db:"password" json:"password"`                   // Password 密码hash
	Role         string        `db:"role" json:"role"`                           // Role 角色
	Salt         string        `db:"salt" json:"salt"`                           // Salt 盐
	Gender       *int          `db:"gender" json:"gender,omitempty"`             // Gender 性别
	BirthYear    *int          `db:"birth_year" json:"birth_year,omitempty"`     // BirthYear 出生年
	BirthMonth   *int          `db:"birth_month" json:"birth_month,omitempty"`   // BirthMonth 出生月
	BirthDay     *int          `db:"birth_day" json:"birth_day,omitempty"`       // BirthDay 出生日
	Location     *int64        `db:"location" json:"location,omitempty,string"`  // Location 地区
	Introduction *string       `db:"introduction" json:"introduction,omitempty"` // Introduction 自我介绍
	Avatar       string        `db:"avatar" json:"avatar"`                       // Avatar 头像
	Source       int           `db:"source" json:"source"`                       // Source 注册来源
	IP           int64         `db:"ip" json:"ip,string"`                        // IP 注册IP
	IDCert       types.BitBool `db:"id_cert" json:"id_cert"`                     // IDCert 是否身份认证
	WorkCert     types.BitBool `db:"work_cert" json:"work_cert"`                 // WorkCert 是否工作认证
	IsOrg        types.BitBool `db:"is_org" json:"is_org"`                       // IsOrg 是否机构用户
	IsVIP        types.BitBool `db:"is_vip" json:"is_vip"`                       // IsVIP 是否VIP用户
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}
