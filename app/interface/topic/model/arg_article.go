package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgAddArticle struct {
	// 标题
	// required: true
	Title string `json:"title"`
	//  封面图
	Avatar *string `json:"avatar,omitempty"`
	// 简介
	Introduction string `json:"introduction"` // Introduction 话题简介

	// 是否私有
	// required: true
	Private bool `json:"private"`

	Versions []*AddArticleVersion `json:"versions"`

	// 附件
	Files []*AddArticleFile `json:"files"`

	// 关联话题
	Relations []*AddArticleRelation `json:"relations"`
}

func (p *ArgAddArticle) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Title, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.Files),
		validation.Field(&p.Relations),
		validation.Field(&p.Versions),
		validation.Field(&p.Files),
	)
}

type AddArticleVersion struct {
	// 顺序
	Seq int `json:"seq"`

	// 版本名称
	Name string `json:"name"`

	// 文章语言
	// required: true
	Locale string `json:"locale"`

	// 内容
	// required: true
	Content string `json:"content"`
}

func (p *AddArticleVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.Locale, validation.Required),
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

	// 所关联话题版本 ID
	TopicVersionID int64 `json:"topic_version_id,string" swaggertype:"string"`

	// 是否主话题
	Primary bool `json:"primary"`
}

func (p *AddArticleRelation) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicVersionID, validation.Required),
	)
}

type ArgUpdateArticle struct {
	// 文章ID
	ID int64 `json:"id,string"  swaggertype:"string"`

	// 标题
	Title *string `json:"title,omitempty"`
	//  封面图
	Avatar *string `json:"avatar,omitempty"`

	Introduction *string `json:"introduction,omitempty"` // Introduction 话题简介

	// 是否私有
	Private *bool `json:"private,omitempty"`
}

func (p *ArgUpdateArticle) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Title, validation.NilOrNotEmpty),
		validation.Field(&p.Avatar, validation.NilOrNotEmpty),
		validation.Field(&p.Introduction, validation.NilOrNotEmpty),
	)
}

type ArgUpdateArticleVersion struct {
	// 文章版本ID
	ID int64 `json:"id,string"  swaggertype:"string"`

	// 内容
	Content *string `json:"content,omitempty"`

	// 文章语言
	Locale *string `json:"locale,omitempty"`

	// 版本名称
	Name *string `json:"name"`

	ChangeDesc string `json:"change_desc"`

	ChangeID string `json:"change_id"`
}

func (p *ArgUpdateArticleVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Content, validation.NilOrNotEmpty),
		validation.Field(&p.Name, validation.NilOrNotEmpty, validation.RuneLength(0, 250)),
		validation.Field(&p.Locale, validation.NilOrNotEmpty),
		validation.Field(&p.ChangeDesc, validation.Required),
		validation.Field(&p.ChangeID, validation.Required),
	)
}

type ArgReportArticle struct {
	// 文章ID
	ID int64 `json:"id,string"  swaggertype:"string"`
	// 举报类型
	Type int `json:"type"`
	// 补充信息
	Desc *string `json:"desc"`
}

func (p *ArgReportArticle) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Type, validation.Required),
		validation.Field(&p.Desc, validation.NilOrNotEmpty, validation.RuneLength(0, 500)),
	)
}
