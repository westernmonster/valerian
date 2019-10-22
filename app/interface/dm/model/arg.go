package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgMarkRead struct {
	ID int64 `json:"id,string,omitempty" swaggertype:"string"`
}

func (p *ArgMarkRead) Validate() (err error) {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
	)
}
