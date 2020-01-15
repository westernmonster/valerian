package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgLike struct {
	// 目标ID
	TargetID int64 `json:"target_id,string" swaggertype:"string"`

	// 目标类型
	//  article, discussion, revise
	TargetType string `json:"target_type"`
}

func (p *ArgLike) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TargetID, validation.Required),
		validation.Field(&p.TargetType, validation.Required, validation.In(TargetTypeRevise, TargetTypeDiscussion, TargetTypeComment, TargetTypeArticle)),
	)
}

type ArgCancelLike struct {
	// 目标ID
	TargetID int64 `json:"target_id,string" swaggertype:"string"`

	// 目标类型
	// topic, article, discussion, revise
	TargetType string `json:"target_type"`
}

func (p *ArgCancelLike) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TargetID, validation.Required),
		validation.Field(&p.TargetType, validation.Required, validation.In(TargetTypeRevise, TargetTypeDiscussion, TargetTypeComment, TargetTypeArticle)),
	)
}

type ArgDislike struct {
	// 目标ID
	TargetID int64 `json:"target_id,string" swaggertype:"string"`

	// 目标类型
	// topic, article, discussion, revise
	TargetType string `json:"target_type"`
}

func (p *ArgDislike) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TargetID, validation.Required),
		validation.Field(&p.TargetType, validation.Required, validation.In(TargetTypeRevise, TargetTypeDiscussion, TargetTypeComment, TargetTypeArticle)),
	)
}

type ArgCancelDislike struct {
	// 目标ID
	TargetID int64 `json:"target_id,string" swaggertype:"string"`

	// 目标类型
	// topic, article, discussion, revise
	TargetType string `json:"target_type"`
}

func (p *ArgCancelDislike) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TargetID, validation.Required),
		validation.Field(&p.TargetType, validation.Required, validation.In(TargetTypeRevise, TargetTypeDiscussion, TargetTypeComment, TargetTypeArticle)),
	)
}
