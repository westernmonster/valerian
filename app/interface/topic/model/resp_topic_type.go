package model

type TopicTypeResp struct {
	ID   int    `db:"id" json:"id"`     // ID ID
	Name string `db:"name" json:"name"` // Name 话题类型
}
