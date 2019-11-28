package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgArticleFile struct {
	ID       *int64 `json:"id,string,omitempty" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url"`  // FileURL 文件地址
	FileType string `json:"file_type"` // FileType 文件类型
	Seq      int    `json:"seq"`       // Seq 文件顺序
}

func (p *ArgArticleFile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileName, validation.Required),
		validation.Field(&p.FileURL, validation.Required, is.URL),
		validation.Field(&p.FileType, validation.Required, validation.In(FileTypeWord, FileTypePPT, FileTypeExcel, FileTypePDF)),
	)
}

type ArgSaveArticleFiles struct {
	ArticleID int64             `json:"article_id,string" swaggertype:"string"`
	Items     []*ArgArticleFile `json:"items"`
}

func (p *ArgSaveArticleFiles) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ArticleID, validation.Required),
		validation.Field(&p.Items),
	)
}
