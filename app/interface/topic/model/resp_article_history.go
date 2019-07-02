package model

type ArticleHistoryResp struct {
	ID               int64             `json:"id,string" swaggertype:"string"`                 // ID ID
	ArticleVersionID int64             `json:"article_version_id,string" swaggertype:"string"` // ArticleVersionID 文章版本ID
	UpdatedBy        int64             `json:"updated_by,string"`                              // UpdatedBy 更新人
	Updator          *BasicAccountResp `json:"updator,omitempty"`                              // Updator 更新人
	Content          *string           `json:"content,omitempty"`                              // Content 内容
	ContentText      *string           `json:"content_text,omitempty"`                         // ContentText 内容
	Diff             *string           `json:"diff,omitempty"`                                 // Diff 更改内容
	ChangeID         string            `json:"change_id"`                                      // ChangeID
	Description      string            `json:"change_desc"`                                    // Description 修订说明
	Seq              int               `json:"seq"`                                            // Seq 版本顺序
	UpdatedAt        int64             `json:"updated_at"`                                     // UpdatedAt 创建时间
	CreatedAt        int64             `json:"created_at"`                                     // CreatedAt 创建时间
}
