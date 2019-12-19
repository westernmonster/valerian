package model

type FeedbackListResp struct {
	Items  []*Feedback `json:"items"`
	Paging *Paging     `json:"paging"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
