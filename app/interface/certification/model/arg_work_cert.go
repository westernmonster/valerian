package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgWorkCert struct {
	// 工作证
	WorkPic string `json:"work_pic"`
	// 其他证明
	OtherPic string `json:"other_pic"`
	// 公司
	Company string `json:"company"`
	// 部门
	Department string `json:"department"`
	// 职位
	Position string `json:""position`

	// 工作证有效期
	// 过期时间
	// Unix时间戳
	ExpiresAt int64 `json:"expires_at"`
}

func (p *ArgWorkCert) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.WorkPic, validation.Required, is.URL),
		validation.Field(&p.OtherPic, is.URL),
		validation.Field(&p.Company, validation.Required),
		validation.Field(&p.Department, validation.Required),
		validation.Field(&p.Position, validation.Required),
	)
}
