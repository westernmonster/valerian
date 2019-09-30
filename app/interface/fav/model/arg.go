package model

type ArgAddFav struct {
	// 目标ID
	TargetID int64 `json:"target_id,string" swaggertype:"string"`

	// 目标类型
	// topic, article, discussion, revise
	TargetType string `json:"target_type"`
}

type ArgDelFav struct {
	// 目标ID
	TargetID int64 `json:"target_id,string" swaggertype:"string"`

	// 目标类型
	// topic, article, discussion, revise
	TargetType string `json:"target_type"`
}
