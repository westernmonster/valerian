package model

import "valerian/library/database/sqlx/types"

type AccountArticleAttr struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	ArticleID int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Read      types.BitBool `db:"read" json:"read"`                    // Read 是否阅读
	Like      types.BitBool `db:"like" json:"like"`                    // Like 是否点赞
	Fav       types.BitBool `db:"fav" json:"fav"`                      // Fav 是否收藏
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type Article struct {
	ID             int64         `db:"id" json:"id,string"`                    // ID ID
	Title          string        `db:"title" json:"title"`                     // Title 标题
	Content        string        `db:"content" json:"content"`                 // Content 内容
	ContentText    string        `db:"content_text" json:"content_text"`       // ContentText
	DisableRevise  types.BitBool `db:"disable_revise" json:"disable_revise"`   // DisableRevise 禁止补充
	DisableComment types.BitBool `db:"disable_comment" json:"disable_comment"` // DisableComment 禁止评论
	CreatedBy      int64         `db:"created_by" json:"created_by,string"`    // CreatedBy 创建人
	Deleted        types.BitBool `db:"deleted" json:"deleted"`                 // Deleted 是否删除
	CreatedAt      int64         `db:"created_at" json:"created_at"`           // CreatedAt 创建时间
	UpdatedAt      int64         `db:"updated_at" json:"updated_at"`           // UpdatedAt 更新时间
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
	ID          int64         `db:"id" json:"id,string"`                 // ID ID
	ArticleID   int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Content     string        `db:"content" json:"content"`              // Content 内容
	ContentText string        `db:"content_text" json:"content_text"`    // ContentText 内容
	Seq         int           `db:"seq" json:"seq"`                      // Seq 顺序
	Diff        string        `db:"diff" json:"diff"`                    // Diff 更改内容
	UpdatedBy   int64         `db:"updated_by" json:"updated_by,string"` // UpdatedBy 更新人
	ChangeDesc  string        `db:"change_desc" json:"change_desc"`      // ChangeDesc 修订说明
	Deleted     types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type TopicCatalog struct {
	ID         int64         `db:"id" json:"id,string"`                    // ID ID
	Name       string        `db:"name" json:"name"`                       // Name 名称
	Seq        int           `db:"seq" json:"seq"`                         // Seq 顺序
	Type       string        `db:"type" json:"type"`                       // Type 类型
	ParentID   int64         `db:"parent_id" json:"parent_id,string"`      // ParentID 父ID
	RefID      *int64        `db:"ref_id" json:"ref_id,omitempty,string"`  // RefID 引用ID
	TopicID    int64         `db:"topic_id" json:"topic_id,string"`        // TopicID 话题ID
	IsPrimary  types.BitBool `db:"is_primary" json:"is_primary"`           // IsPrimary
	Permission *string       `db:"permission" json:"permission,omitempty"` // Permission
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                 // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`           // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`           // UpdatedAt 更新时间
}

type ImageURL struct {
	ID         int64         `db:"id" json:"id,string"`               // ID ID
	TargetID   int64         `db:"target_id" json:"target_id,string"` // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`    // TargetType 目标类型
	URL        string        `db:"url" json:"url"`                    // URL 图片路径
	Deleted    types.BitBool `db:"deleted" json:"deleted"`            // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`      // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`      // UpdatedAt 更新时间
}

type Revise struct {
	ID          int64         `db:"id" json:"id,string"`                 // ID ID
	ArticleID   int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Title       string        `db:"title" json:"title"`                  // Title 标题
	Content     string        `db:"content" json:"content"`              // Content 内容
	ContentText string        `db:"content_text" json:"content_text"`    // ContentText 内容纯文本
	CreatedBy   int64         `db:"created_by" json:"created_by,string"` // CreatedBy 创建人
	Deleted     types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type ReviseFile struct {
	ID        int64         `db:"id" json:"id,string"`                // ID ID
	FileName  string        `db:"file_name" json:"file_name"`         // FileName 文件名
	FileURL   *string       `db:"file_url" json:"file_url,omitempty"` // FileURL 文件地址
	Seq       int           `db:"seq" json:"seq"`                     // Seq 文件顺序
	ReviseID  int64         `db:"revise_id" json:"revise_id,string"`  // ReviseID 文章ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`             // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`       // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`       // UpdatedAt 更新时间
}
