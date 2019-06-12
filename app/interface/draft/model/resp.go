package model

type DraftCategoryResp struct {
	// ID
	ID int64 `db:"id" json:"id,string" swaggertype:"string"`
	// 分类名
	Name string `db:"name" json:"name"` // Name 分类名
	// 颜色ID
	ColorID int64 `db:"color_id" json:"color_id,string" swaggertype:"string"` // ColorID 颜色ID
	// 颜色值
	Color string `db:"color" json:"color"`

	AccountID int64 `db:"account_id" json:"-"` // AccountID 账户ID
}

type DraftResp struct {
	ID        int64              `db:"id" json:"id,string" swaggertype:"string"`
	Title     string             `db:"title" json:"title"`                  // Title 颜色名
	Content   *string            `db:"content" json:"content,omitempty"`    // Content 内容
	Text      *string            `db:"text" json:"-"`                       // Text 内容纯文本
	AccountID int64              `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	Category  *DraftCategoryResp `json:"category"`                          // 分类
	CreatedAt int64              `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64              `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
