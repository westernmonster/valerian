package model

type FeedbackTypeResp struct {
	ID   int    `json:"id,string" swaggertype:"string"` // ID ID
	Type string `json:"type"`                           // Type 目标类型
	Name string `json:"name"`                           // Name 备注
}
