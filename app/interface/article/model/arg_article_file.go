package model

type UpdateArticleFile struct {
	ID       *int64 `json:"id,string,omitempty" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url`   // FileURL 文件地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}

type ArgSaveArticleFiles struct {
	ArticleID int64                `json:"article_id,string" swaggertype:"string"`
	Items     []*UpdateArticleFile `json:"items"`
}
