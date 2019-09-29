package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgFollow struct {
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
}

func (p *ArgFollow) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}

type ArgUnfollow struct {
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
}

func (p *ArgUnfollow) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required),
	)
}
