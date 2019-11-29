package model

type MemberTopicResp struct {
	Items  []*TargetTopic `json:"items"`
	Paging *Paging        `json:"paging"`
}
