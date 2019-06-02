package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgRelatedTopic struct {
	// 关联话题ID
	TopicID int64 `json:"topic_id,string"`

	// 顺序
	Seq int `json:"seq"`

	// 类型
	// normal
	// strong
	Type string `json:"type"`
}

func (p *ArgRelatedTopic) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required.Error(`请传入关联话题ID`)),
		validation.Field(&p.Type, validation.Required.Error(`请传入关联话题类型`),
			validation.In(TopicRelationStrong, TopicRelationNormal).Error("话题类型"),
		),
	)
}

type ArgBatchSaveRelatedTopics struct {
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 关联话题
	RelatedTopics []*ArgRelatedTopic `json:"related_topics"`
}

func (p *ArgBatchSaveRelatedTopics) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required.Error(`请传入话题ID`)),
		validation.Field(&p.RelatedTopics, validation.Required.Error(`请传入关联话题`)),
	)
}
