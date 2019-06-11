package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ArgAddDraftCategory struct {
	Name    string `json:"name"`
	ColorID int64  `json:"color_id,string" swaggertype:"string"`
}

func (p *ArgAddDraftCategory) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required.Error(`请输入分类名`)),
		validation.Field(&p.ColorID, validation.Required.Error(`请输入颜色`)),
	)
}

type ArgUpdateDraftCategory struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	ColorID int64 `json:"color_id,string" swaggertype:"string"`
}

func (p *ArgUpdateDraftCategory) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required.Error(`请输入分类ID`)),
		validation.Field(&p.Name, validation.Required.Error(`请输入分类名`)),
		validation.Field(&p.ColorID, validation.Required.Error(`请输入颜色`)),
	)
}
