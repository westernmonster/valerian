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
	AppID    string   `form:"appid" params:"appid"`
	KW       string   `form:"kw" params:"kw"`
	KwFields []string `form:"kw_fields,split" params:"kw_fields"`
	Order    []string `form:"order,split" params:"order"`
	Sort     []string `form:"sort,split" params:"sort" default:"desc"`
	Pn       int      `form:"pn" params:"pn;Range(1,5000)" default:"1"`
	Ps       int      `form:"ps" params:"ps;Range(1,10000)" default:"50"`
	Debug    bool     `form:"debug" params:"debug"`
	Source   []string
}
