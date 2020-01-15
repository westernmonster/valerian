package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ArgAddFeedback struct {
	//  目标ID
	//  会员、话题、文章、讨论、补充的ID
	//  如果是反馈，传-1
	TargetID int64 `json:"target_id,string"`
	//  目标类型
	//  1 反馈
	//  2 成员
	//  3 话题
	//  4 文章
	//  5 讨论
	//  6 补充
	//  7 评论
	TargetType int32 `json:"target_type"`
	//  举报类型
	//  通过 /list/feedback_types 获取
	Type int32 `json:"type"`
	// 备注
	Desc string `json:"desc"`
}

func (p *ArgAddFeedback) Validate() (err error) {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Type, validation.Required),
		validation.Field(&p.TargetID, validation.Required),
		validation.Field(&p.TargetType, validation.In(TargetTypeFeedback, TargetTypeMember, TargetTypeTopic, TargetTypeArticle, TargetTypeDiscuss, TargetTypeRevise, TargetTypeComment)),
	)
}

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
		validation.Field(&p.VerifyStatus, validation.Required, validation.In(FeedbackStatusCommited, FeedbackStatusApproved, FeedbackStatusRejected)),
	)
}
