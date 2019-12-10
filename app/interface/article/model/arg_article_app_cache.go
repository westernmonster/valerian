package model

type ArgArticleAppCache struct {
	// required: true
	Items []*AppCacheReqItem `json:"items"`
}

type AppCacheReqItem struct{
	// 文章 ID
	// required: true
	ID int64 `json:"id,string" swaggertype:"string"`
	// 更新时间戳
	UpdatedAt string `json:"updated_at"`
}

type ArgReviseAppCache struct {
	// required: true
	Items []*AppCacheReqItem `json:"items"`
}