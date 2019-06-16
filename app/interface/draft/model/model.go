package model

import "valerian/library/database/sqlx/types"

type Color struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	Name      string        `db:"name" json:"name"`                    // Name 颜色名
	Color     string        `db:"color" json:"color,omitempty"`        // Color 颜色
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type DraftCategory struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	Name      string        `db:"name" json:"name"`                    // Name 分类名
	ColorID   int64         `db:"color_id" json:"color_id,string"`     // ColorID 颜色ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type Draft struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	Title      string        `db:"title" json:"title"`                    // Title 颜色名
	Content    *string       `db:"content" json:"content,omitempty"`      // Content 内容
	Text       *string       `db:"text" json:"text,omitempty"`            // Text 内容纯文本
	AccountID  int64         `db:"account_id" json:"account_id,string"`   // AccountID 账户ID
	CategoryID *int64        `db:"category_id" json:"category_id,string"` // CategoryID 分类
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}
