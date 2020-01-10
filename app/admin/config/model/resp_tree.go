package model

type TreeItem struct {
	TreeID       int32  `json:"tree_id"`
	Name         string `json:"name"`
	PlatformID   int32  `json:"platform_id"`
	PlatformName string `json:"platform_name"`
	Mark         string `json:"mark"`
	CreatedAt    int64  `json:"created_at"` // CreatedAt 创建时间
	UpdatedAt    int64  `json:"updated_at"` // UpdatedAt 更新时间
}

type TreeListResp struct {
	Total    int32       `json:"total"`
	PageSize int32       `json:"page_size"`
	Page     int32       `json:"page"`
	Items    []*TreeItem `json:"items"`
}
