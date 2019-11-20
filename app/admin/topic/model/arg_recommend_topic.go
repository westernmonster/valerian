package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgAddRecommendTopic struct {
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
}

func (p *ArgAddRecommendTopic) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
	)
}

type ArgDelete struct {
	ID int64 `json:"id,string,omitempty" swaggertype:"string"`
}

func (p *ArgDelete) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
	)
}
