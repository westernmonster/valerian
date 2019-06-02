package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgID struct {
	// 话题ID
	ID int64 `json:"id,string" swaggertype:"string"`
}

func (p *ArgID) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required.Error(`请输入ID`)),
	)
}
