package model

type TopicCatalogResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`
	//  名称
	// required: true
	Name string `json:"name"`

	// 顺序
	// required: true
	Seq int `json:"seq"`

	// 类型
	// required: true
	Type string `json:"type"`

	// 引用ID
	RefID *int64 `json:"ref_id,string,omitempty" swaggertype:"string"`

	// 父级ID
	ParentID int64 `json:"parent_id,string" swaggertype:"string"`
}
