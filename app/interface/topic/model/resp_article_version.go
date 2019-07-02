package model

type ArticleVersionResp struct {
	// 版本ID
	ID int64 `db:"id" json:"id,string"  swaggertype:"string"`

	// 文章ID
	ArticleID int64 `db:"article_id" json:"article_id,string" swaggertype:"string"`

	// 顺序
	Seq int `db:"seq" json:"seq"`

	// 版本名称
	Name string `db:"name" json:"name"`

	// 文章名
	Content string `db:"content" json:"content"`

	Histories []*ArticleHistoryResp `db:"-", json:"histories,omitempty"`
}
