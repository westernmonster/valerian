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

	// 所关联话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`

	// 是否主话题
	Primary bool `json:"primary"`

	// 类型
	// view  // 只允许查看
	// edit // 允许所有成员编辑
	Permission string `json:"permission"`
}

func (p *ArgAddArticleRelation) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ArticleID, validation.Required),
		validation.Field(&p.ParentID),
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Permission, validation.Required, validation.In(AuthPermissionEdit, AuthPermissionView)),
	)
}

type AddArticleRelation struct {
	// 类目分类ID 如果根目录则传0
	ParentID int64 `json:"parent_id,string" swaggertype:"string"`

	// 所关联话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`

	// 是否主话题
	Primary bool `json:"primary"`

	// 类型
	// view  // 只允许查看
	// edit // 允许所有成员编辑
	Permission string `json:"permission"`
}

func (p *AddArticleRelation) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Permission, validation.Required, validation.In(AuthPermissionEdit, AuthPermissionView)),
	)
}

type ArgUpdateArticleRelation struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 是否主话题
	Primary bool `json:"primary"`

	// 类型
	// view  // 只允许查看
	// edit // 允许所有成员编辑
	Permission string `json:"permission"`
}

func (p *ArgUpdateArticleRelation) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Permission, validation.Required, validation.In(AuthPermissionEdit, AuthPermissionView)),
	)
}
