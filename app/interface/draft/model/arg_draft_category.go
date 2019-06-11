package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ArgAddDraftCategory struct {
	Name string `json:"name"`

	Color string `json:"color"`
}

func (p *ArgAddDraftCategory) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required.Error(`请输入颜色名`)),
		validation.Field(&p.Color, validation.Required.Error(`请输入颜色值`),
			validation.Match(regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)).Error(`颜色值格式不正确`)),
	)
}

type ArgUpdateDraftCategory struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Color string `json:"color"`
}

func (p *ArgUpdateDraftCategory) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required.Error(`请输入颜色ID`)),
		validation.Field(&p.Name, validation.Required.Error(`请输入颜色名`)),
		validation.Field(&p.Color, validation.Required.Error(`请输入颜色值`),
			validation.Match(regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)).Error(`颜色值格式不正确`)),
	)
}
