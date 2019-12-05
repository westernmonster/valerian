package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgVerifyFeedback struct {
	FeedbackID   int64  `json:"feedback_id,string" swaggertype:"string"`
	VerifyStatus int  `json:"verify_status"`
	VerifyDesc   string `json:"verify_desc"`
}

func (p *ArgVerifyFeedback) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FeedbackID, validation.Required),
		validation.Field(&p.VerifyStatus, validation.Required),
		validation.Field(&p.VerifyDesc, validation.Required),
	)
}
