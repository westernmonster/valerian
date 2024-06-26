package model

import (
	"valerian/library/ecode"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CatalogArticleItem struct {
	ID int64 `json:"id,string" swaggertype:"string"`
	// 标题
	Title string `json:"title"`
	// 内容
	Excerpt string `json:"excerpt"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 补充个数
	ReviseCount int `json:"revise_count"`
	// 评论数
	CommentCount int `json:"comment_count"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`
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

	Children []*TopicLevel2Catalog `json:"children,omitempty"`

	// 文章
	Article *CatalogArticleItem `json:"article,omitempty"`
}

func (p *TopicLevel1Catalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`),
			validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet, TopicCatalogTopic).Error("类型不正确")),
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
	case TopicCatalogTopic:
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

	Children []*TopicChildCatalog `json:"children,omitempty"`

	// 文章
	Article *CatalogArticleItem `json:"article,omitempty"`
}

func (p *TopicLevel2Catalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`),
			validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet, TopicCatalogTopic, TopicCatalogTopic).Error("类型不正确")),
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

	// 文章
	Article *CatalogArticleItem `json:"article,omitempty"`
}

func (p *TopicChildCatalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`), validation.In(TopicCatalogArticle, TopicCatalogTestSet, TopicCatalogTopic).Error("类型不正确")),
		validation.Field(&p.RefID, ValidateRefID(p.Type)),
	)
}

// 批量更新话题分类请求
type ArgSaveTopicCatalog struct {
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
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
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Items),
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
		validation.Field(&p.Name, validation.Required, validation.RuneLength(0, 100)),
		validation.Field(&p.Type, validation.Required, validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet, TopicCatalogTopic)),
		validation.Field(&p.RefID, ValidateRefID(p.Type)),
	)
}
