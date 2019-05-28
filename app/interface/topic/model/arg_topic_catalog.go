package model

import (
	"valerian/library/ecode"

	validation "github.com/go-ozzo/ozzo-validation"
)

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
			return ecode.ShouldNotSetRefID
		}
		break
	case TopicCatalogArticle:
		if refID == nil {
			return ecode.RefIDRequired
		}
		break
	case TopicCatalogTestSet:
		if refID == nil {
			return ecode.RefIDRequired
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
			return ecode.ChildrenIsNotAllowed
		}
		break
	case TopicCatalogTestSet:
		if children != nil && len(children) > 0 {
			return ecode.ChildrenIsNotAllowed
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
			return ecode.ChildrenIsNotAllowed
		}
		break
	case TopicCatalogTestSet:
		if children != nil && len(children) > 0 {
			return ecode.ChildrenIsNotAllowed
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
