package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgNewTopicVersion struct {
	// 话题ID
	TopicID int64 `json:"topic_id,string", swaggertype:"string"`
	// 版本名称
	Name string `json:"name"`
}

func (p *ArgNewTopicVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Name,
			validation.Required,
			validation.RuneLength(0, 250),
		),
	)
}

type TopicVersionItem struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 顺序
	Seq int `json:"seq"`

	// 版本名称
	Name string `json:"name"`
}

func (p *TopicVersionItem) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Name, validation.Required),
	)

}

type ArgSaveTopicVersions struct {
	TopicID int64               `json:"topic_id,string" swaggertype:"string"`
	Items   []*TopicVersionItem `json:"items"`
}

func (p *ArgSaveTopicVersions) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Items),
		validation.Field(&p.TopicID, validation.Required),
	)

}
