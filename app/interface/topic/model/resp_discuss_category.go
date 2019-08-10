package model

type DiscussCategoryResp struct {
	ID      int64  `json:"id,string" swaggertype:"string"`       // ID ID
	TopicID int64  `json:"topic_id,string" swaggertype:"string"` // ID ID
	Name    string `json:"name"`
	Seq     int    `json:"seq"`
}
