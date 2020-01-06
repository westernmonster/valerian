package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgCreateCommonConfig struct {
	Team     string `json:"team" `
	Env      string `json:"env" `
	Zone     string `json:"zone" `
	Name     string `json:"name" `
	State    int32  `json:"state" `
	Comment  string `json:"comment" `
	Mark     string `json:"mark" `
	SkipLint bool   `json:"skiplint"`
}

func (p *ArgCreateCommonConfig) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Team, validation.Required),
		validation.Field(&p.Env, validation.Required),
		validation.Field(&p.Zone, validation.Required),
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.State, validation.Required),
		validation.Field(&p.Comment, validation.Required),
		validation.Field(&p.Mark, validation.Required),
	)
}

type ArgUpdateCommonConfig struct {
	ID        int64  `json:"id,string"`
	State     int32  `json:"state"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	Mark      string `json:"mark"`
	UpdatedAt int64  `json:"updated_at"`
	SkipLint  bool   `json:"skiplint"`
}

func (p *ArgUpdateCommonConfig) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.State, validation.Required, validation.In(ConfigStateInProgress, ConfigStateEnd)),
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Comment, validation.Required),
		validation.Field(&p.Mark, validation.Required),
		validation.Field(&p.UpdatedAt, validation.Required),
	)
}
