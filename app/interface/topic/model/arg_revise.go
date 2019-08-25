package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgAddRevise struct {
	// 文章 ID
	ArticleID int64 `json:"article_id,string" swaggertype:"string"`

	// 内容
	Content string `json:"content"`

	// 附件
	Files []*AddReviseFile `json:"files"`
}

func (p *ArgAddRevise) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ArticleID, validation.Required),
		validation.Field(&p.Content, validation.Required),
		validation.Field(&p.Files),
	)
}

type AddReviseFile struct {
	FileName string `json:"file_name"`          // FileName 文件名
	FileURL  string `json:"file_url,omitempty"` // FileURL 文件地址
	Seq      int    `json:"seq"`                // Seq 文件顺序
}

func (p *AddReviseFile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileName, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.FileURL, validation.Required, is.URL),
	)
}
