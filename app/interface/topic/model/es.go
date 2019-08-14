package model

import "encoding/json"

type ES struct {
	Addr string
}

type Page struct {
	Pn    int   `json:"num"`
	Ps    int   `json:"size"`
	Total int64 `json:"total"`
}

type SearchResult struct {
	Order  string            `json:"order"`
	Sort   string            `json:"sort"`
	Result []json.RawMessage `json:"data"`
	Page   *Page             `json:"page"`
	Debug  string            `json:"debug"`
}

type BasicSearchParams struct {
	// 搜索关键词
	KW string `json:"kw"`
	// 搜索关键词所用的字段
	KwFields []string `json:"kw_fields"`
	// 排序的顺序
	// desc, asc
	Order []string `json:"order"`
	// 排序的字段
	Sort []string `json:"sort"`
	// 页码
	Pn int `json:"pn"`
	// 每页大小
	Ps int `json:"ps"`
	// 是否输出Debug信息
	Debug bool `json:"debug"`
	// 输出的字段
	Source []string `json:"source"`
}

type AccountSearchParams struct {
	Bsp *BasicSearchParams
	// Query string `json:"query"`
}
