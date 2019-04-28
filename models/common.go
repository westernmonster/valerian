package models

// CountryCode 国际电话区号
// swagger:model
type CountryCode struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	//  国家名
	Name string `json:"name"`
	//  国旗
	Emoji string `json:"emoji"`
	// 国家中文名
	CnName string `json:"cn_name"`
	// 编码
	Code string `json:"code"`
	// 前缀/电话区号
	// example: 86
	Prefix string `json:"prefix"`
}
