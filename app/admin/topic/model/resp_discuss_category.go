package model

type DiscussCategoryResp struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 名称
	Name string `json:"name"`
	// 顺序
	Seq int32 `json:"seq"`
}
