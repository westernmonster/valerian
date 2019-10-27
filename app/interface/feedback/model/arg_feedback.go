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
