package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgAddArticle struct {
	// 标题
	// required: true
	Title string `json:"title"`
	// 内容
	// required: true
	Content string `json:"content"`
	//  封面图
	Cover *string `json:"cover,omitempty"`
	// 简介
	Introduction string `json:"introduction"` // Introduction 话题简介

	// 是否私有
	// required: true
	Private bool `json:"private"`

	// 文章语言
	// required: true
	Locale string `json:"locale"`

	// 文章集合ID
	ArticleSetID *int64 `json:"article_set_id,string,omitempty"  swaggertype:"string"`

	// 版本名称
	// required: true
	VersionName string `json:"version_name"`

	// 附件
	Files []*AddArticleFile `json:"files"`

	// 关联话题
	Relations []*AddArticleRelation `json:"relations"`
}

func (p *ArgAddArticle) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Title, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.Content, validation.Required),
		validation.Field(&p.Locale, validation.Required),
		validation.Field(&p.VersionName, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.Files),
		validation.Field(&p.Relations),
	)
}

type AddArticleFile struct {
	FileName string `json:"file_name"`          // FileName 文件名
	FileURL  string `json:"file_url,omitempty"` // FileURL 文件地址
	Seq      int    `json:"seq"`                // Seq 文件顺序
}

func (p *AddArticleFile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileName, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.FileURL, validation.Required, is.URL),
	)
}

type AddArticleRelation struct {
	// 类目分类ID 如果根目录则传0
	ParentID int64 `json:"parent_id,string" swaggertype:"string"`

	// 所关联话题 ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`

	// 是否主话题
	Primary bool `json:"primary"`
}

func (p *AddArticleRelation) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
	)
}

type ArgUpdateArticle struct {
	// 文章ID
	ID int64 `json:"id,string"  swaggertype:"string"`

	// 标题
	Title *string `json:"title,omitempty"`
	// 内容
	Content *string `json:"content,omitempty"`

	// 文章语言
	Locale *string `json:"locale"`

	//  封面图
	Cover *string `json:"cover,omitempty"`
	// 简介
	Introduction *string `json:"introduction,omitempty"` // Introduction 话题简介
	// 是否私有
	Private *bool `json:"private,omitempty"`

	// 版本名称
	VersionName *string `json:"version_name"`
}

func (p *ArgUpdateArticle) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Title, validation.NilOrNotEmpty, validation.RuneLength(0, 250)),
		validation.Field(&p.Introduction, validation.NilOrNotEmpty, validation.RuneLength(0, 500)),
		validation.Field(&p.Content, validation.NilOrNotEmpty),
		validation.Field(&p.Locale, validation.NilOrNotEmpty),
		validation.Field(&p.Cover, validation.NilOrNotEmpty, is.URL),
		validation.Field(&p.VersionName, validation.NilOrNotEmpty, validation.RuneLength(0, 250)),
	)
}
