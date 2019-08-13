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
	Result []json.RawMessage `json:"result"`
	Page   *Page             `json:"page"`
	Debug  string            `json:"debug"`
}

type BasicSearchParams struct {
	KW       string   `json:"kw"`
	KwFields []string `json:"kw_fields"`
	Order    []string `json:"order"`
	Sort     []string `json:"sort"`
	Pn       int      `json:"pn"`
	Ps       int      `json:"ps"`
	Debug    bool     `json:"debug"`
	Source   []string `json:"source"`
}

type AccountSearchParams struct {
	Bsp   *BasicSearchParams
	Query string `json:"query"`
}
