package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgLinkInfo struct {
	// 外部链接
	Link string `json:"link"`
}

func (p *ArgLinkInfo) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Link, validation.Required, is.URL),
	)
}
