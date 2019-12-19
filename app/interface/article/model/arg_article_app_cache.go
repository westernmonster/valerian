package model

type ArgArticleAppCache struct {
	// required: true
	Items   []*AppCacheReqItem `json:"items"`
	// 参照文章获取 "目前支持：files,relations,histories,meta" (Revise 忽略这个参数）
	Include string             `json:"include"`
}

type AppCacheReqItem struct {
	// 文章 ID
	// required: true
	ID int64 `json:"id,string" swaggertype:"string"`
	// 更新时间戳
	UpdatedAt int64 `json:"updated_at"`
}

type ArgReviseAppCache struct {
	// required: true
	Items []*AppCacheReqItem `json:"items"`
}
