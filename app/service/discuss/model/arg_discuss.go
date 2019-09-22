package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgAddDiscuss struct {
	// 话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 分类ID
	CategoryID int64 `json:"category_id,string" swaggertype:"string"`
	// 标题，可以不传
	Title *string `json:"title,omitempty"`
	// 内容
	Content string `json:"content"`

	// 附件
	Files []*AddDiscussFile `json:"files"`
}

func (p *ArgAddDiscuss) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Title, validation.RuneLength(0, 120)),
		validation.Field(&p.Content, validation.Required),
		validation.Field(&p.CategoryID, validation.Required),
		validation.Field(&p.TopicID, validation.Required),
	)
}

type AddDiscussFile struct {
	FileName string `json:"file_name"`          // FileName 文件名
	FileURL  string `json:"file_url,omitempty"` // FileURL 文件地址
	Seq      int    `json:"seq"`                // Seq 文件顺序
}

func (p *AddDiscussFile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileName, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.FileURL, validation.Required, is.URL),
	)
}

type ArgUpdateDiscuss struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题，可以不传
	Title *string `json:"title,omitempty"`

	// 内容
	Content string `json:"content"`

	// 附件
	Files []*AddDiscussFile `json:"files"`
}

func (p *ArgUpdateDiscuss) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Content, validation.Required),
		validation.Field(&p.Files),
	)
}

type ArgDelDiscuss struct {
	ID int64 `json:"id,string" swaggertype:"string"`
}

func (p *ArgDelDiscuss) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
	)
}
