package model

type TopicCatalogResp struct {
	ID        int64  `db:"id" json:"id,string"`                   // ID ID
	TopicID   int64  `db:"topic_id" json:"topic_id,string"`       // TopicID 分类ID
	Name      string `db:"name" json:"name"`                      // Name 名称
	Seq       int    `db:"seq" json:"seq"`                        // Seq 顺序
	Type      string `db:"type" json:"type"`                      // Type 类型
	RefID     *int64 `db:"ref_id" json:"ref_id,string,omitempty"` // RefID 引用ID
	ParentID  int64  `db:"parent_id" json:"parent_id,string"`     // ParentID 父级ID, 一级分类的父ID为 0
	CreatedAt int64  `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}
