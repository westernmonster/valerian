package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgNewVersion struct {
	FromTopicID int64  `json:"from_topic_id,string", swaggertype:"string"`
	VersionName string `json:"version_name"`
}

func (p *ArgNewVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FromTopicID, validation.Required.Error(`请输入来源话题`)),
		validation.Field(&p.VersionName,
			validation.Required.Error(`请输入版本名`),
			validation.RuneLength(0, 250).Error(`版本名最大长度为250个字符`),
		),
	)
}
