package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgWorkCert struct {
	// AccountID 账户ID
	AccountID   int64  `json:"account_id,string" swaggertype:"string"`
	// 审核原因
	AuditResult string `json:"audit_result"`
	// 通过 or 拒绝
	Approve     bool   `json:"approve"`
}

func (p *ArgWorkCert) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}
