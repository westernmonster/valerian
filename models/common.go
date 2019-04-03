package models

type CountryCode struct {
	ID     int64  `json:"id,string"` // ID ID
	EnName string `json:"en_name"`   // EnName 国家英文名
	CnName string `json:"cn_name"`   // CnName 国家中文名
	Code   string `json:"code"`      // Code 编码
	Prefix string `json:"prefix"`    // Prefix 前缀
}
