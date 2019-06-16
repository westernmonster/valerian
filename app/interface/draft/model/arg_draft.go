package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgAddDraft struct {
	Title      string  `json:"title"`
	Content    *string `json:"content"`
	CategoryID *int64  `json:"category_id,string,omitempty" swaggertype:"string"`
}

func (p *ArgAddDraft) Validate() (err error) {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Title, validation.Required.Error(`请输入标题`),
			validation.RuneLength(1, 250).Error(`标题最大长度为250个字符`)),
	)
}

type ArgUpdateDraft struct {
	ID         int64   `json:"id,string,omitempty" swaggertype:"string"`
	Title      string  `json:"title"`
	Content    *string `json:"content"`
	CategoryID *int64  `json:"category_id,string,omitempty" swaggertype:"string"`
}

func (p *ArgUpdateDraft) Validate() (err error) {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required.Error(`请输入ID`)),
		validation.Field(&p.Title, validation.Required.Error(`请输入标题`),
			validation.RuneLength(1, 250).Error(`标题最大长度为250个字符`)),
	)
}
