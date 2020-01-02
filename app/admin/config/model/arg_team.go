package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgCreateTeam struct {
	Name string `json:"name"`
	Env  string `json:"env"`
	Zone string `json:"zone"`
}

func (p *ArgCreateTeam) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Env, validation.Required),
		validation.Field(&p.Zone, validation.Required),
	)
}
