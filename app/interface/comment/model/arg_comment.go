package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ArgAddComment struct {
	// 类型
	// revise 补充
	// article 文章
	// discussion 话题讨论
	// comment 评论
	TargetType string `json:"type"`

	// 内容
	Content string `json:"content"`
	// 资源ID
	// 表示话题、文章、讨论的ID
	TargetID int64 `json:"resource_id,string" swaggertype:"string"`
}

func (p *ArgAddComment) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TargetType, validation.Required, validation.In(TargetTypeRevise, TargetTypeDiscussion, TargetTypeArticle, TargetTypeComment)),
		validation.Field(&p.Content, validation.Required, validation.RuneLength(0, 1000)),
		validation.Field(&p.TargetID, validation.Required),
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
