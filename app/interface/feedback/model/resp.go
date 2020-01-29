package model

type FeedbackTypeResp struct {
	ID   int    `json:"id,string" swaggertype:"string"` // ID ID
	Type string `json:"type"`                           // Type 目标类型
	Name string `json:"name"`                           // Name 备注
}

type FeedbackListResp struct {
	Items  []*FeedbackItem `json:"items"`
	Paging *Paging         `json:"paging"`
}

type FeedbackItem struct {
	ID           int64  `json:"id,string"`         // ID ID
	TargetID     int64  `json:"target_id,string"`  // TargetID 目标ID
	TargetType   int32  `json:"target_type"`       // TargetType 目标类型
	TargetDesc   string `json:"target_desc"`       // TargetDesc 目标描述（用于通知）
	FeedbackType int32  `json:"feedback_type"`     // FeedbackType 举报类型
	FeedbackDesc string `json:"feedback_desc"`     // FeedbackDesc 备注
	CreatedBy    int64  `json:"created_by,string"` // CreatedBy 举报人
	CreatedAt    int64  `json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64  `json:"updated_at"`        // UpdatedAt 更新时间
	VerifyStatus int32  `json:"verify_status"`     // CheckStatus 审核状态
	VerifyDesc   string `json:"verify_desc"`       // CheckDesc 审核结果/原因
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
