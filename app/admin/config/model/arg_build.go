package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgCreateBuild struct {
	AppName string `json:"app_name"`
	Env     string `json:"env"`
	Zone    string `json:"zone"`
	Name    string `json:"name"`
	TagID   int64  `json:"tag_id,string" swaggertype:"string"`
	TreeID  int    `json:"tree_id"`
}

func (p *ArgCreateBuild) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AppName, validation.Required),
		validation.Field(&p.Env, validation.Required),
		validation.Field(&p.Zone, validation.Required),
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.TagID, validation.Required),
		validation.Field(&p.TreeID, validation.Required),
	)
}
