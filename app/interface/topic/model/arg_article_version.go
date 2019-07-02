package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ArgNewArticleVersion struct {
	ArticleID int64  `json:"article_id,string", swaggertype:"string"`
	Name      string `json:"name"`
}

func (p *ArgNewArticleVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ArticleID, validation.Required),
		validation.Field(&p.Name, validation.Required, validation.RuneLength(0, 250)),
	)
}

type ArticleVersionItem struct {
	// 文章版本ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 顺序
	Seq int `json:"seq"`

	// 版本名称
	Name string `json:"name"`
}

func (p *ArticleVersionItem) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Name, validation.Required, validation.RuneLength(0, 250)),
	)
}

type ArgSaveArticleVersions struct {
	ArticleID int64                 `json:"article_id,string" swaggertype:"string"`
	Items     []*ArticleVersionItem `json:"items"`
}

func (p *ArgSaveArticleVersions) Validate() error {
	fmt.Println(1111111)
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ArticleID, validation.Required),
		validation.Field(&p.Items),
	)
}
