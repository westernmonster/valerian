package model

import "valerian/library/database/sqlx/types"

type Article struct {
	ID           int64         `db:"id" json:"id,string"`                         // ID ID
	Title        string        `db:"title" json:"title"`                          // Title 标题
	Content      string        `db:"content" json:"content"`                      // Content 内容
	Cover        *string       `db:"cover" json:"cover,omitempty"`                // Cover 封面图
	Introduction string        `db:"introduction" json:"introduction"`            // Introduction 话题简介
	Private      types.BitBool `db:"private" json:"private"`                      // Private 是否私有
	ArticleSetID int64         `db:"article_set_id" json:"article_set_id,string"` // ArticleSetID 文章集合ID
	VersionName  string        `db:"version_name" json:"version_name"`            // VersionName 版本名称
	Seq          int           `db:"seq" json:"seq"`                              // Seq 版本顺序
	CreatedBy    int64         `db:"created_by" json:"created_by,string"`         // CreatedBy 创建人
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                      // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`                // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`                // UpdatedAt 更新时间
}

type ArticleSet struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type ArticleFile struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	FileName  string        `db:"file_name" json:"file_name"`          // FileName 文件名
	FileURL   *string       `db:"file_url" json:"file_url,omitempty"`  // FileURL 文件地址
	Seq       int           `db:"seq" json:"seq"`                      // Seq 文件顺序
	ArticleID int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
