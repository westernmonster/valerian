package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgCreateConfig struct {
	AppName string `json:"app_name"`
	Env     string `json:"env"`
	Zone    string `json:"zone"`
	Name    string `json:"name"`
	State   int32  `json:"state"`
	From    int64  `json:"from,string" swaggertype:"string"`
	// 配置文件内容
	Comment  string `json:"comment"`
	Mark     string `json:"mark"`
	TreeID   int    `json:"tree_id"`
	SkipLint bool   `json:"skiplint"`
}

func (p *ArgCreateConfig) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AppName, validation.Required),
		validation.Field(&p.Env, validation.Required),
		validation.Field(&p.Zone, validation.Required),
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.State, validation.Required),
		validation.Field(&p.Comment, validation.Required),
		validation.Field(&p.Mark, validation.Required),
		validation.Field(&p.TreeID, validation.Required),
	)
}
