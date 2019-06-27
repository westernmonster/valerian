package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgNewTopicVersion struct {
	FromTopicID int64  `json:"from_topic_id,string", swaggertype:"string"`
	VersionName string `json:"version_name"`
}

func (p *ArgNewTopicVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FromTopicID, validation.Required.Error(`请输入来源话题`)),
		validation.Field(&p.VersionName,
			validation.Required.Error(`请输入版本名`),
			validation.RuneLength(0, 250).Error(`版本名最大长度为250个字符`),
		),
	)
}

type ArgMergeTopicVersion struct {
	FromTopicSetID int64 `json:"from_topic_set_id,string", swaggertype:"string"`
	ToTopicSetID   int64 `json:"to_topic_set_id,string", swaggertype:"string"`
}

func (p *ArgMergeTopicVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FromTopicSetID, validation.Required.Error(`请输入来源话题集合ID`)),
		validation.Field(&p.ToTopicSetID, validation.Required.Error(`请输入合并话题集合ID`)),
	)
}

type TopicVersionItem struct {
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`

	// 顺序
	Seq int `json:"seq"`

	// 版本名称
	VersionName string `json:"version_name"`
}
type ArgSaveTopicVersions struct {
	TopicSetID int64               `json:"topic_set_id,string" swaggertype:"string"`
	Items      []*TopicVersionItem `json:"items"`
}

func (p *ArgSaveTopicVersions) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Items),
		validation.Field(&p.TopicSetID, validation.Required),
	)

}
