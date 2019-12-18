package model

type ArgArticleAppCache struct {
	// required: true
	Items   []*AppCacheReqItem `json:"items"`
	Include string             `json:"include"`
}

type AppCacheReqItem struct {
	// 文章 ID
	// required: true
	ID int64 `json:"id,string" swaggertype:"string"`
	// 更新时间戳
	UpdatedAt int64 `json:"updated_at,string" swaggertype:"string"`
}

type ArgReviseAppCache struct {
	// required: true
	Items []*AppCacheReqItem `json:"items"`
}
