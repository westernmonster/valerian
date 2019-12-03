package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgVerifyFeedback struct {
	FeedbackID   int64 `json:"feedback_id,string" swaggertype:"string"`
	VerifyStatus int32 `json:"verify_status,string" swaggertype:"string"`
	VerifyDesc   int64 `json:"verify_desc,string" swaggertype:"string"`
}

func (p *ArgVerifyFeedback) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FeedbackID, validation.Required),
		validation.Field(&p.VerifyStatus, validation.Required),
		validation.Field(&p.VerifyDesc, validation.Required),
	)
}
