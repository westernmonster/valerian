package models

import (
	"valerian/infrastructure/berr"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TopicCatalog struct {
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

type TopicLevel1Catalog struct {
	ID *int64 `json:"id,string" swaggertype:"string"`
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

	Children []*TopicLevel2Catalog `json:"children"`
}

func (p *TopicLevel1Catalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`),
			validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet).Error("类型不正确")),
		validation.Field(&p.Children, ValidateLevel1Children(p.Type)),
		validation.Field(&p.RefID, ValidateRefID(p.Type)),
	)
}

func ValidateRefID(rtype string) *ValidateTypeRule {
	return &ValidateTypeRule{
		Type: rtype,
	}
}

type ValidateTypeRule struct {
	Type string
}

func (p *ValidateTypeRule) Validate(v interface{}) error {
	refID := v.(*int64)
	switch p.Type {
	case TopicCatalogTaxonomy:
		if refID != nil {
			return berr.Errorf(`分类无需设置 ref_id`)
		}
		break
	case TopicCatalogArticle:
		if refID == nil {
			return berr.Errorf(`请输入ref_id`)
		}
		break
	case TopicCatalogTestSet:
		if refID == nil {
			return berr.Errorf(`请输入ref_id`)
		}
		break
	}

	return nil
}

func ValidateLevel1Children(rtype string) *ValidateLevel1ChildrenRule {
	return &ValidateLevel1ChildrenRule{
		Type: rtype,
	}
}

type ValidateLevel1ChildrenRule struct {
	Type string
}

func (p *ValidateLevel1ChildrenRule) Validate(v interface{}) error {
	children := v.([]*TopicLevel2Catalog)
	switch p.Type {
	case TopicCatalogTaxonomy:
		break
	case TopicCatalogArticle:
		if children != nil && len(children) > 0 {
			return berr.Errorf(`文章不能有下级`)
		}
		break
	case TopicCatalogTestSet:
		if children != nil && len(children) > 0 {
			return berr.Errorf(`题库不能有下级`)
		}
		break
	}

	return nil
}

func ValidateLevel2Children(rtype string) *ValidateLevel2ChildrenRule {
	return &ValidateLevel2ChildrenRule{
		Type: rtype,
	}
}

type ValidateLevel2ChildrenRule struct {
	Type string
}

func (p *ValidateLevel2ChildrenRule) Validate(v interface{}) error {
	children := v.([]*TopicChildCatalog)
	switch p.Type {
	case TopicCatalogTaxonomy:
		break
	case TopicCatalogArticle:
		if children != nil && len(children) > 0 {
			return berr.Errorf(`文章不能有下级`)
		}
		break
	case TopicCatalogTestSet:
		if children != nil && len(children) > 0 {
			return berr.Errorf(`题库不能有下级`)
		}
		break
	}

	return nil
}

// func checkLevel2Duplicate(value interface{}) error {
// 	children, _ := value.([]*TopicLevel2Catalog)
// }

type TopicLevel2Catalog struct {
	ID *int64 `json:"id,string" swaggertype:"string"`
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

	Children []*TopicChildCatalog `json:"children"`
}

func (p *TopicLevel2Catalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`),
			validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet).Error("类型不正确")),
		validation.Field(&p.Children, ValidateLevel2Children(p.Type)),
		validation.Field(&p.RefID, ValidateRefID(p.Type)),
	)
}

type TopicChildCatalog struct {
	ID *int64 `json:"id,string" swaggertype:"string"`
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
}

func (p *TopicChildCatalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`),
			validation.In(TopicCatalogArticle, TopicCatalogTestSet).Error("类型不正确")),
		validation.Field(&p.RefID, ValidateRefID(p.Type)),
	)
}

// 批量更新话题分类请求
type ArgSaveTopicCatalog struct {
	// 话题分类
	// required: true
	Items []*ArgTopicCatalog `json:"items"`
	// 父级ID
	// required: true
	ParentID int64 `json:"parent_id,string,omitempty" swaggertype:"string"`
}

func (p *ArgSaveTopicCatalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Items,
			validation.Required.Error(`请添加话题分类`)),
	)
}

type ArgTopicCatalog struct {
	ID *int64 `json:"id,string" swaggertype:"string"`
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
}

func (p *ArgTopicCatalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`),
			validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet).Error("类型不正确")),
		validation.Field(&p.RefID, ValidateRefID(p.Type)),
	)
}
