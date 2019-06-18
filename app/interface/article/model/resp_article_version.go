package model

type ArticleVersionResp struct {
	// 集合ID
	ArticleSetID int64 `db:"article_set_id" json:"article_set_id,string"  swaggertype:"string"`
	// 文章ID
	ArticleID int64 `db:"article_id" json:"article_id,string" swaggertype:"string"`

	// 顺序
	Seq int `db:"seq" json:"seq"`

	// 版本名称
	VersionName string `db:"version_name" json:"version_name"`

	// 文章名
	ArticleName string `db:"article_name" json:"article_name"`
}
