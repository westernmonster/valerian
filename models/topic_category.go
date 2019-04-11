package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type TopicCategory struct {
	ID        int64  `db:"id" json:"id,string"`                 // ID ID
	TopicID   int64  `db:"topic_id" json:"topic_id,string"`     // TopicID 分类ID
	Name      string `db:"name" json:"name"`                    // Name 分类名
	ParentID  int64  `db:"parent_id" json:"parent_id,string"`   // ParentID 父级ID, 一级分类的父ID为 0
	CreatedBy int64  `db:"created_by" json:"created_by,string"` // CreatedBy 创建人
	Seq       int    `db:"seq" json:"seq"`                      // Seq 顺序
	CreatedAt int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

// 创建话题分类请求
// swagger:model
type CreateTopicCategoryReq struct {
	// 话题ID
	// required: true
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	//  分类名
	// required: true
	Name string `json:"name"`
	// 顺序
	// required: true
	Seq int `json:"seq"`
	// ParentID 父级ID, 一级分类的父ID为 0
	// required: true
	ParentID int64 `json:"parent_id,string" swaggertype:"string"`
}

func (p *CreateTopicCategoryReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID,
			validation.Required.Error(`请输入话题ID`),
		),
		validation.Field(&p.Name,
			validation.Required.Error(`请输入分类名`),
			validation.RuneLength(0, 100).Error(`分类名最大长度为100个字符`)),
	)
}

// 更新话题分类请求
// swagger:model
type UpdateTopicCategoryReq struct {
	//  分类名
	// required: true
	Name string `json:"name"`
	// 顺序
	// required: true
	Seq int `json:"seq"`
	// ParentID 父级ID, 一级分类的父ID为 0
	// required: true
	ParentID int64 `json:"parent_id,string"  swaggertype:"string"`
}

func (p *UpdateTopicCategoryReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入分类名`),
			validation.RuneLength(0, 100).Error(`分类名最大长度为100个字符`)),
	)
}

type TopicCategoryChildItem struct {
	ID *int64 `json:"id,string" swaggertype:"string"`
	//  分类名
	// required: true
	Name string `json:"name"`
	// 顺序
	// required: true
	Seq int `json:"seq"`
}

func (p *TopicCategoryChildItem) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入分类名`),
			validation.RuneLength(0, 100).Error(`分类名最大长度为100个字符`)),
	)
}

type TopicCategoryParentItem struct {
	ID *int64 `json:"id,string" swaggertype:"string"`
	//  分类名
	// required: true
	Name string `json:"name"`
	// 顺序
	// required: true
	Seq int `json:"seq"`

	Children []*TopicCategoryChildItem `json:"children"`
}

func (p *TopicCategoryParentItem) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入分类名`),
			validation.RuneLength(0, 100).Error(`分类名最大长度为100个字符`)),
	)
}

type TopicCategoriesResp struct {
	// 话题分类
	Items []*TopicCategoryParentItem `json:"items"`
	// 话题ID
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
}

// 批量更新话题分类请求
// swagger:model
type SaveTopicCategoriesReq struct {
	// 话题分类
	// required: true
	Items []*TopicCategoryParentItem `json:"items"`
	// 话题ID
	// required: true
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
}

func (p *SaveTopicCategoriesReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID,
			validation.Required.Error(`请输入话题ID`),
		),
		validation.Field(&p.Items,
			validation.Required.Error(`请添加话题分类`)),
	)
}
