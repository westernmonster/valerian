package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgCreateApp struct {
	AppName string `json:"app_name"`
	TreeID  int32  `json:"tree_id"`
}

func (p *ArgCreateApp) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AppName, validation.Required),
		validation.Field(&p.TreeID, validation.Required),
	)
}
