package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgVerifyFeedback struct {
	FeedbackID int64 `json:"feedback_id,string" swaggertype:"string"`
	//审核状态 0-未审核，1-审核通过，2 审核不通过
	VerifyStatus int32  `json:"verify_status"`
	VerifyDesc   string `json:"verify_desc"`
}

func (p *ArgVerifyFeedback) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FeedbackID, validation.Required),
	)
}
