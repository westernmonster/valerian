package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgCreateTag struct {
	AppName   string `json:"app_name" `
	Env       string `json:"env" `
	Zone      string `json:"zone" `
	ConfigIDs string `json:"config_ids" `
	Mark      string `json:"mark" `
	TreeID    int    `json:"tree_id" `
}

func (p *ArgCreateTag) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AppName, validation.Required),
		validation.Field(&p.Env, validation.Required),
		validation.Field(&p.Zone, validation.Required),
		validation.Field(&p.ConfigIDs, validation.Required),
		validation.Field(&p.Mark, validation.Required),
		validation.Field(&p.TreeID, validation.Required),
	)
}

type ArgLastTags struct {
	AppName string `json:"app_name" `
	Zone    string `json:"zone" `
	Env     string `json:"env" `
	Build   string `json:"build" `
	TreeID  int    `json:"tree_id" `
}

type ArgTagsByBuild struct {
	AppName string `json:"app_name" `
	Zone    string `json:"zone" `
	Env     string `json:"env" `
	Build   string `json:"build" `
	Pn      int    `json:"pn" default:"1" `
	Ps      int    `json:"ps" default:"20" `
	TreeID  int    `json:"tree_id" `
}

type ArgTag struct {
	TagID int64 `json:"tag_id" `
}

type ArgUpdateTag struct {
	AppName   string `json:"app_name" `
	Env       string `json:"env" `
	Zone      string `json:"zone" `
	ConfigIDs string `json:"config_ids"`
	Mark      string `json:"mark"`
	Build     string `json:"build"`
	TreeID    int    `json:"tree_id"`
	Force     int    `json:"force"`
}

type ArgUpdateTagID struct {
	AppName string `json:"app_name" `
	Zone    string `json:"zone" `
	Env     string `json:"env" `
	Build   string `json:"build" `
	TreeID  int    `json:"tree_id" `
	TagID   int64  `json:"tag_id,string" swaggertype:"string" `
}

type ArgTagConfigDiff struct {
	TagID   int64  `json:"tag_id,string"  swaggertype:"string" `
	Name    string `json:"name" `
	TreeID  int    `json:"tree_id" `
	AppID   int64  `json:"app_id,string"  swaggertype:"string" `
	BuildID int64  `json:"build_id,string"  swaggertype:"string" `
}
