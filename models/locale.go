package models

type Locale struct {
	// ID ID
	ID int64 `db:"id" json:"id,string"`
	// Locale 语言编码
	Locale string `db:"locale" json:"locale"`
	// Name 语言名称
	Name string `db:"name" json:"name"`
}
