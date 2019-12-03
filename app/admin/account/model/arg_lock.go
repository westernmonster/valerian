package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgAccountLock struct {
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
}

func (p *ArgAccountLock) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}

type ArgAccountUnlock struct {
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
}

func (p *ArgAccountUnlock) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}
