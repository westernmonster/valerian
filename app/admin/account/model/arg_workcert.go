package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgWorkCert struct {
	AccountID   int64  `json:"account_id,string" swaggertype:"string"`
	AuditResult string `json:"audit_result"`
	Approve     bool   `json:"approve"`
}

func (p *ArgWorkCert) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}
