package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgAddFav struct {
	// 目标ID
	TargetID int64 `json:"target_id,string" swaggertype:"string"`

	// 目标类型
	// topic, article, discussion, revise
	TargetType string `json:"target_type"`
}

func (p *ArgAddFav) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TargetID, validation.Required),
		validation.Field(&p.TargetType, validation.Required, validation.In(TargetTypeTopic, TargetTypeRevise, TargetTypeDiscussion, TargetTypeArticle)),
	)
}

type ArgDelFav struct {
	// 目标ID
	TargetID int64 `json:"target_id,string" swaggertype:"string"`

	// 目标类型
	// topic, article, discussion, revise
	TargetType string `json:"target_type"`
}

func (p *ArgDelFav) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TargetID, validation.Required),
		validation.Field(&p.TargetType, validation.Required, validation.In(TargetTypeTopic, TargetTypeRevise, TargetTypeDiscussion, TargetTypeArticle)),
	)
}
