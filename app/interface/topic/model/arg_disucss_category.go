package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgSaveDiscussCategories struct {
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 话题分类
	// required: true
	Items []*ArgDisucssCategory `json:"items"`
}

func (p *ArgSaveDiscussCategories) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Items),
	)
}

type ArgDisucssCategory struct {
	ID *int64 `json:"id,string" swaggertype:"string"`
	//  名称
	// required: true
	Name string `json:"name"`

	// 顺序
	// required: true
	Seq int `json:"seq"`
}

func (p *ArgDisucssCategory) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required, validation.RuneLength(0, 100)),
	)
}
