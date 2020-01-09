package model

type CommonConfigItem struct {
	// ID
	ID        int64  `json:"id,string"`
	Name      string `json:"name"`       // Name 名称
	Comment   string `json:"comment"`    // Comment 配置文件
	State     int32  `json:"state"`      // State 状态
	Mark      string `json:"mark"`       // Mark 备注
	Operator  string `json:"operator"`   // Operator 操作者
	CreatedAt int64  `json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64  `json:"updated_at"` // UpdatedAt 更新时间
}

type CommonConfigListResp struct {
	Total    int32               `json:"total"`
	PageSize int32               `json:"page_size"`
	Page     int32               `json:"page"`
	Items    []*CommonConfigItem `json:"items"`
}
