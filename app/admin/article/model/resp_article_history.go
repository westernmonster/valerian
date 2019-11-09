package model

type ArticleHistoryResp struct {
	ID         int64    `json:"id,string" swaggertype:"string"`         // ID ID
	ArticleID  int64    `json:"article_id,string" swaggertype:"string"` // ArticleID 文章ID
	Updator    *Creator `json:"updator,omitempty"`                      // Updator 更新人
	Diff       string   `json:"diff,omitempty"`                         // Diff 更改内容
	ChangeDesc string   `json:"change_desc"`                            // ChangeDesc 修订说明
	Seq        int      `json:"seq"`                                    // Seq 版本顺序
	UpdatedAt  int64    `json:"updated_at"`                             // UpdatedAt 创建时间
	CreatedAt  int64    `json:"created_at"`                             // CreatedAt 创建时间
}

type ArticleHistoryItem struct {
	ID         int64    `json:"id,string" swaggertype:"string"`         // ID ID
	ArticleID  int64    `json:"article_id,string" swaggertype:"string"` // ArticleID 文章ID
	Updator    *Creator `json:"updator,omitempty"`                      // Updator 更新人
	ChangeDesc string   `json:"change_desc"`                            // ChangeDesc 修订说明
	Seq        int      `json:"seq"`                                    // Seq 版本顺序
	UpdatedAt  int64    `json:"updated_at"`                             // UpdatedAt 创建时间
	CreatedAt  int64    `json:"created_at"`                             // CreatedAt 创建时间
}

type ArticleHistoryListResp struct {
	Items  []*ArticleHistoryItem `json:"items"`
	Paging *Paging               `json:"paging"`
}
