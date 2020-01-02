package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ArgAddDiscuss struct {
	// 话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 分类ID
	CategoryID int64 `json:"category_id,string" swaggertype:"string"`
	// 标题，可以不传
	Title *string `json:"title,omitempty"`
	// 内容
	Content string `json:"content"`

	// 附件
	Files []*AddDiscussionFile `json:"files"`
}

func (p *ArgAddDiscuss) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Title, validation.RuneLength(0, 120)),
		validation.Field(&p.Content, validation.Required),
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Files),
	)
}

type ArgUpdateDiscuss struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题，可以不传
	Title *string `json:"title,omitempty"`

	// 内容
	Content string `json:"content"`
}

func (p *ArgUpdateDiscuss) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Content, validation.Required),
	)
}
