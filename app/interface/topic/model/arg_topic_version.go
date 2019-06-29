package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgNewTopicVersion struct {
	TopicID     int64  `json:"topic_id,string", swaggertype:"string"`
	VersionName string `json:"version_name"`
}

func (p *ArgNewTopicVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.VersionName,
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
	VersionName string `json:"version_name"`
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
