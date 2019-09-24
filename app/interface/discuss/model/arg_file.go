package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UpdateDiscussionFile struct {
	ID       *int64 `json:"id,string,omitempty" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url"`  // FileURL 文件地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}

func (p *UpdateDiscussionFile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileName, validation.Required),
		validation.Field(&p.FileURL, validation.Required, is.URL),
	)
}

type ArgSaveDiscussionFiles struct {
	DiscussionID int64                   `json:"article_id,string" swaggertype:"string"`
	Items        []*UpdateDiscussionFile `json:"items"`
}

func (p *ArgSaveDiscussionFiles) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.DiscussionID, validation.Required),
		validation.Field(&p.Items),
	)
}

type AddDiscussionFile struct {
	FileName string `json:"file_name"`          // FileName 文件名
	FileURL  string `json:"file_url,omitempty"` // FileURL 文件地址
	Seq      int    `json:"seq"`                // Seq 文件顺序
}

func (p *AddDiscussionFile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileName, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.FileURL, validation.Required, is.URL),
	)
}
