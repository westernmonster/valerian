package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgNewVersion struct {
	FromArticleID int64  `json:"from_article_id,string", swaggertype:"string"`
	VersionName   string `json:"version_name"`
}

func (p *ArgNewVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FromArticleID, validation.Required.Error(`请输入来源文章`)),
		validation.Field(&p.VersionName,
			validation.Required.Error(`请输入版本名`),
			validation.RuneLength(0, 250).Error(`版本名最大长度为250个字符`),
		),
	)
}

type ArgMergeVersion struct {
	FromArticleSetID int64 `json:"from_article_set_id,string", swaggertype:"string"`
	ToArticleSetID   int64 `json:"to_article_set_id,string", swaggertype:"string"`
}

func (p *ArgMergeVersion) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FromArticleSetID, validation.Required.Error(`请输入来源文章集合ID`)),
		validation.Field(&p.ToArticleSetID, validation.Required.Error(`请输入合并文章集合ID`)),
	)
}

type VersionItem struct {
	ArticleID int64 `json:"article_id,string" swaggertype:"string"`

	// 顺序
	Seq int `json:"seq"`

	// 版本名称
	VersionName string `json:"version_name"`
}
type ArgSaveVersions struct {
	ArticleSetID int64          `json:"article_set_id,string" swaggertype:"string"`
	Items        []*VersionItem `json:"items"`
}
