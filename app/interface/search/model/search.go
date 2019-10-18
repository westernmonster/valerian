package model

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TopicSearchResult struct {
	// 会员数据
	Data []*ESTopic `json:"data"`
	// 分页
	Paging *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type ArticleSearchResult struct {
	// 会员数据
	Data []*ESArticle `json:"data"`
	// 分页
	Paging *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type DiscussionSearchResult struct {
	// 会员数据
	Data []*ESDiscussion `json:"data"`
	// 分页
	Paging *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type AllSearchParams struct {
	// 搜索关键词
	KW string `json:"kw"`
}

type AllSearchResult struct {
	Topics      []*ESTopic `json:"topics"`
	TopicsCount int64      `json:"topics_count"`

	Articles      []*ESArticle `json:"articles"`
	ArticlesCount int64        `json:"articles_count"`

	Accounts      []*ESAccount `json:"accounts"`
	AccountsCount int64        `json:"accounts_count"`

	Discussions      []*ESDiscussion `json:"discussions"`
	DiscussionsCount int64           `json:"discussions_count"`
}

type AccountSearchResult struct {
	// 会员数据
	Data []*ESAccount `json:"data"`
	// 分页
	Paging *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type Paging struct {
	// 统计数量
	Total int64 `json:"total,omitempty"`
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}

type Page struct {
	// 页码
	Pn int `json:"num"`
	// 页大小
	Ps int `json:"size"`
	// 统计数量
	Total int64 `json:"total"`
}

type SearchResult struct {
	Order  string            `json:"order"`
	Sort   string            `json:"sort"`
	Result []json.RawMessage `json:"data"`
	Page   *Page             `json:"page"`
	Debug  string            `json:"debug"`
}

type BasicSearchParams struct {
	// 搜索关键词
	KW string `json:"kw" form:"kw"`
	// 搜索关键词所用的字段
	KwFields []string `json:"kw_fields" form:"kw_fields"`
	// 排序的顺序
	// desc, asc
	Order []string `json:"order" form:"order"`
	// 排序的字段
	Sort []string `json:"sort" form:"sort"`
	// 页码
	Pn int `json:"pn" form:"pn"`
	// 每页大小
	Ps int `json:"ps" form:"ps"`
	// 是否输出Debug信息
	Debug bool `json:"debug" form:"debug"`
	// 输出的字段
	Source []string `json:"source" form:"source"`
}

func (p *BasicSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Pn, validation.Required),
		validation.Field(&p.Ps, validation.Required),
	)
}

type ArticleSearchParams struct {
	*BasicSearchParams
}

func (p *ArticleSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Pn, validation.Required),
		validation.Field(&p.Ps, validation.Required),
	)
}

type AccountSearchParams struct {
	*BasicSearchParams
}

func (p *AccountSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Pn, validation.Required),
		validation.Field(&p.Ps, validation.Required),
	)
}

type TopicSearchParams struct {
	*BasicSearchParams
	// Query string `json:"query"`
}

func (p *TopicSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Pn, validation.Required),
		validation.Field(&p.Ps, validation.Required),
	)
}

type DiscussionSearchParams struct {
	*BasicSearchParams
	// Query string `json:"query"`
}

func (p *DiscussionSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Pn, validation.Required),
		validation.Field(&p.Ps, validation.Required),
	)
}
