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
}
