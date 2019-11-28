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
	Content string `json:"content"`

	// 禁止补充
	DisableRevise bool `json:"disable_revise"`

	// 禁止评论
	DisableComment bool `json:"disable_comment"`

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
		validation.Field(&p.Files),
	)
}

type AddArticleFile struct {
	FileName string `json:"file_name"`          // FileName 文件名
	FileURL  string `json:"file_url,omitempty"` // FileURL 文件地址
	FileType string `json:"file_type"`          // FileType 文件类型
	Seq      int    `json:"seq"`                // Seq 文件顺序
}

func (p *AddArticleFile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileName, validation.Required, validation.RuneLength(0, 250)),
		validation.Field(&p.FileURL, validation.Required, is.URL),
		validation.Field(&p.FileType, validation.Required, validation.In(FileTypeWord, FileTypePPT, FileTypeExcel, FileTypePDF)),
	)
}

type ArgUpdateArticle struct {
	// 文章ID
	ID int64 `json:"id,string"  swaggertype:"string"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 内容
	Content string `json:"content"`

	// 禁止补充
	DisableRevise *bool `json:"disable_revise"`

	// 禁止评论
	DisableComment *bool `json:"disable_comment"`

	// 更新说明
	ChangeDesc string `json:"change_desc"`
}

func (p *ArgUpdateArticle) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.Content, validation.Required),
	)
}
