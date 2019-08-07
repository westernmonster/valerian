package model

import "valerian/library/database/sqlx/types"

type Article struct {
	ID           int64         `db:"id" json:"id,string"`                 // ID ID
	Title        string        `db:"title" json:"title"`                  // Title 标题
	Avatar       *string       `db:"avatar" json:"avatar,omitempty"`      // Avatar 封面图
	Introduction string        `db:"introduction" json:"introduction"`    // Introduction 话题简介
	Private      types.BitBool `db:"private" json:"private"`              // Private 是否私有
	CreatedBy    int64         `db:"created_by" json:"created_by,string"` // CreatedBy 创建人
	Deleted      types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type ArticleVersion struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	Name      string        `db:"name" json:"name"`                    // Name 版本名
	Locale    string        `db:"locale" json:"locale"`                // Locale 语言编码
	ArticleID int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Content   string        `db:"content" json:"content"`              // Content 内容
	Seq       int           `db:"seq" json:"seq"`                      // Seq 版本顺序
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type ArticleFile struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	FileName  string        `db:"file_name" json:"file_name"`          // FileName 文件名
	FileURL   string        `db:"file_url" json:"file_url,omitempty"`  // FileURL 文件地址
	Seq       int           `db:"seq" json:"seq"`                      // Seq 文件顺序
	ArticleID int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type ArticleHistory struct {
	ID               int64         `db:"id" json:"id,string"`                                 // ID ID
	ArticleVersionID int64         `db:"article_version_id" json:"article_version_id,string"` // ArticleVersionID 文章版本ID
	UpdatedBy        int64         `db:"updated_by" json:"updated_by,string"`                 // UpdatedBy 更新人
	Content          string        `db:"content" json:"content"`                              // Content 内容
	ContentText      string        `db:"content_text" json:"content_text"`                    // ContentText 内容
	Seq              int           `db:"seq" json:"seq"`                                      // Seq 顺序
	Diff             string        `db:"diff" json:"diff"`                                    // Diff 更改内容
	ChangeID         string        `db:"change_id" json:"change_id,omitempty"`                // ChangeID 变更描述
	Description      string        `db:"description" json:"description"`                      // Description 修订说明
	Deleted          types.BitBool `db:"deleted" json:"deleted"`                              // Deleted 是否删除
	CreatedAt        int64         `db:"created_at" json:"created_at"`                        // CreatedAt 创建时间
	UpdatedAt        int64         `db:"updated_at" json:"updated_at"`                        // UpdatedAt 更新时间
}
