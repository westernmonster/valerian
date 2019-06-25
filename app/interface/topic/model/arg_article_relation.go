package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgSetPrimaryArticleRelation struct {
	// ID 文章在话题目录中的ID
	ID int64 `json:"id,string,omitempty" swaggertype:"string"`

	ArticleID int64 `json:"article_id,string" swaggertype:"string"`
}

func (p *ArgSetPrimaryArticleRelation) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ArticleID, validation.Required),
		validation.Field(&p.ID, validation.Required),
	)
}

type ArgDelArticleRelation struct {
	// ID 文章在话题目录中的ID
	ID int64 `json:"id,string,omitempty" swaggertype:"string"`

	ArticleID int64 `json:"article_id,string" swaggertype:"string"`
}

func (p *ArgDelArticleRelation) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ArticleID, validation.Required),
		validation.Field(&p.ID, validation.Required),
	)
}

type ArgAddArticleRelation struct {
	ArticleID int64 `json:"article_id,string" swaggertype:"string"`

	// 类目分类ID 如果根目录则传0
	ParentID int64 `json:"parent_id,string" swaggertype:"string"`

	// 所关联话题 ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 是否主话题
	Primary bool `json:"primary"`
}

func (p *ArgAddArticleRelation) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ArticleID, validation.Required),
		validation.Field(&p.ParentID, validation.Required),
		validation.Field(&p.TopicID, validation.Required),
	)
}
