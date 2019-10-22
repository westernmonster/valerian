package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgProcessInvite struct {
	// 请求的ID
	ID int64 `json:"id,string" swaggertype:"string"`

	Result bool `json:"result"`
}

func (p *ArgProcessInvite) Validate() (err error) {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
	)
}
