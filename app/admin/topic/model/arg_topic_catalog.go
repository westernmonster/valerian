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
	// 反对数
	DislikeCount int `json:"dislike_count"`
	// 补充个数
	ReviseCount int `json:"revise_count"`
	// 评论数
	CommentCount int `json:"comment_count"`

	Images []string `json:"images"`
	// 发布日期
	CreatedAt int64 `json:"created_at"`
}

type TopicRootCatalog struct {
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

	IsPrimary bool `json:"is_primary"`

	// 引用ID
	RefID int64 `json:"ref_id,string,omitempty" swaggertype:"string"`

	Children []*TopicParentCatalog `json:"children"`

	// 文章
	Article *TargetArticle `json:"article,omitempty"`
}

func (p *TopicRootCatalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`),
			validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet).Error("类型不正确")),
		validation.Field(&p.Children, ValidateRootChildren(p.Type)),
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
	refID := v.(int64)
	switch p.Type {
	case TopicCatalogTaxonomy:
		if refID != 0 {
			return ecode.ShouldNotSetRefID
		}
		break
	case TopicCatalogArticle:
		if refID == 0 {
			return ecode.RefIDRequired
		}
		break
	case TopicCatalogTestSet:
		if refID == 0 {
			return ecode.RefIDRequired
		}
		break
	}

	return nil
}

func ValidateRootChildren(rtype string) *ValidateRootChildrenRule {
	return &ValidateRootChildrenRule{
		Type: rtype,
	}
}

type ValidateRootChildrenRule struct {
	Type string
}

func (p *ValidateRootChildrenRule) Validate(v interface{}) error {
	children := v.([]*TopicParentCatalog)
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

func ValidateParentChildren(rtype string) *ValidateParentChildrenRule {
	return &ValidateParentChildrenRule{
		Type: rtype,
	}
}

type ValidateParentChildrenRule struct {
	Type string
}

func (p *ValidateParentChildrenRule) Validate(v interface{}) error {
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

// func checkParentDuplicate(value interface{}) error {
// 	children, _ := value.([]*TopicParentCatalog)
// }

type TopicParentCatalog struct {
	ID *int64 `json:"id,string" swaggertype:"string"`
	//  名称
	// required: true
	Name string `json:"name"`

	// 顺序
	// required: true
	Seq int `json:"seq"`

	IsPrimary bool `json:"is_primary"`

	// 类型
	// required: true
	Type string `json:"type"`

	// 引用ID
	RefID int64 `json:"ref_id,string,omitempty" swaggertype:"string"`

	Children []*TopicChildCatalog `json:"children"`

	Article *TargetArticle `json:"article,omitempty"`
}

func (p *TopicParentCatalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name,
			validation.Required.Error(`请输入名称`),
			validation.RuneLength(0, 100).Error(`名称最大长度为100个字符`)),
		validation.Field(&p.Type,
			validation.Required.Error(`请输入类型`),
			validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet).Error("类型不正确")),
		validation.Field(&p.Children, ValidateParentChildren(p.Type)),
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

	IsPrimary bool `json:"is_primary"`

	// 引用ID
	RefID int64 `json:"ref_id,string,omitempty" swaggertype:"string"`

	Article *TargetArticle `json:"article,omitempty"`
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
	Seq int32 `json:"seq"`

	// 类型
	// required: true
	Type string `json:"type"`

	// 引用ID
	RefID int64 `json:"ref_id,string,omitempty" swaggertype:"string"`
}

func (p *ArgTopicCatalog) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required, validation.RuneLength(0, 100)),
		validation.Field(&p.Type, validation.Required, validation.In(TopicCatalogTaxonomy, TopicCatalogArticle, TopicCatalogTestSet)),
		validation.Field(&p.RefID, ValidateRefID(p.Type)),
	)
}

type ArticleItem struct {
	TopicID   int64
	ArticleID int64
}
