package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgFollow struct {
	Aid      int64  `json:"aid,string" swaggertype:"string"`
	Fid      int64  `json:"fid,string" swaggertype:"string"`
	Action   int    `json:"action"`
	RemoteIP string `json:"remote_ip"`
}

func (p *ArgFollow) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Fid, validation.Required),
	)
}
