package model

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ES struct {
	Addr string
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
	KW string `json:"kw"`
	// 搜索关键词所用的字段
	KwFields []string `json:"kw_fields"`
	// 排序的顺序
	// desc, asc
	Order []string `json:"order"`
	// 排序的字段
	Sort []string `json:"sort"`
	// 页码
	Pn int `json:"pn"`
	// 每页大小
	Ps int `json:"ps"`
	// 是否输出Debug信息
	Debug bool `json:"debug"`
	// 输出的字段
	Source []string `json:"source"`
}

func (p *BasicSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Pn, validation.Required),
		validation.Field(&p.Ps, validation.Required),
	)
}

type ArticleSearchParams struct {
	Bsp *BasicSearchParams
}

func (p *ArticleSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Bsp),
	)
}

type AccountSearchParams struct {
	Bsp *BasicSearchParams
}

func (p *AccountSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Bsp),
	)
}

type TopicSearchParams struct {
	Bsp *BasicSearchParams
	// Query string `json:"query"`
}

func (p *TopicSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Bsp),
	)
}

type DiscussSearchParams struct {
	Bsp *BasicSearchParams
	// Query string `json:"query"`
}

type ESAccount struct {
	// ID
	ID int64 `json:"id,string,omitempty" swaggertype:"string"`

	// 手机
	Mobile *string `json:"mobile,omitempty"`

	// 邮件地址
	Email *string `json:"email,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 角色
	Role string `json:"role,omitempty"`

	// 性别
	Gender *int `json:"gender,omitempty"`

	// 出生年
	BirthYear *int `json:"birth_year,omitempty"`

	// 出生月
	BirthMonth *int `json:"birth_month,omitempty"`

	// 出生日
	BirthDay *int `json:"birth_day,omitempty"`

	// 地区
	Location *int64 `json:"location,omitempty,string" swaggertype:"string"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`

	// 头像
	Avatar *string `json:"avatar,omitempty"`

	// 注册来源
	Source *int `json:"source,omitempty"`

	// 是否身份认证
	IDCert *bool `json:"id_cert,omitempty"`

	// 是否工作认证
	WorkCert *bool `json:"work_cert,omitempty"`

	// 是否机构用户
	IsOrg *bool `json:"is_org,omitempty"`

	// 是否VIP用户
	IsVIP *bool `json:"is_vip,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,string,omitempty" swaggertype:"string"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,string,omitempty"  swaggertype:"string"`
}

type AccountSearchResult struct {
	// 会员数据
	Data []*ESAccount `json:"data"`
	// 分页
	Page *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type ESCreator struct {
	// ID ID
	ID int64 `json:"id,string,omitempty"  swaggertype:"string"`
	// 用户名
	UserName *string `json:"user_name,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty"`

	Introduction *string `json:"introduction,omitempty"`
}

type ESTopic struct {
	// ID ID
	ID int64 `json:"id,string,omitempty"  swaggertype:"string"`
	// Name 话题名
	Name *string `json:"name,omitempty" `
	// Avatar 话题头像
	Avatar *string `json:"avatar,omitempty"`
	// Bg 背景图
	Bg *string `json:"bg,omitempty"`
	// Introduction 话题简介
	Introduction *string `json:"introduction,omitempty"`
	// AllowDiscuss 允许讨论
	AllowDiscuss *bool `json:"allow_discuss,omitempty"`
	// AllowChat 开启群聊
	AllowChat *bool `json:"allow_chat,omitempty"`
	// IsPrivate 是否私密
	IsPrivate *bool `json:"is_private,omitempty"`
	// ViewPermission 查看权限
	ViewPermission *string `json:"view_permission,omitempty"`
	// EditPermission 编辑权限
	EditPermission *string `json:"edit_permission,omitempty"`
	// JoinPermission 加入权限
	JoinPermission *string `json:"join_permission,omitempty"`
	// CatalogViewType 分类视图
	CatalogViewType *string `json:"catalog_view_type,omitempty"`
	// CreatedBy 创建人
	Creator *ESCreator `json:"creator,omitempty"`
	// CreatedAt 创建时间
	CreatedAt *int64 `json:"created_at,string"  swaggertype:"string"`
	// UpdatedAt 更新时间
	UpdatedAt *int64 `json:"updated_at,string"  swaggertype:"string"`
}

type TopicSearchResult struct {
	// 会员数据
	Data []*ESTopic `json:"data"`
	// 分页
	Page *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type ESArticle struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title *string `json:"title"`

	// 内容
	Content *string `json:"content"`

	// 内容
	ContentText *string `json:"content_text"`

	//  禁止补充
	DisableRevise *bool `json:"disable_revise"`

	//  禁止评论
	DisableComment *bool `json:"disable_comment"`

	Creator *ESCreator `json:"creator"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,string,omitempty" swaggertype:"string"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,string,omitempty"  swaggertype:"string"`
}

type ArticleSearchResult struct {
	// 会员数据
	Data []*ESArticle `json:"data"`
	// 分页
	Page *Page `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type ESDiscussionTopic struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// Name 话题名
	Name *string `json:"name,omitempty" `

	// 头像
	Avatar *string `json:"avatar,omitempty"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`
}

type ESDiscussion struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title *string `json:"title"`

	// 内容
	Content *string `json:"content"`

	// 内容
	ContentText *string `json:"content_text"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,string,omitempty" swaggertype:"string"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,string,omitempty"  swaggertype:"string"`

	Creator *ESCreator `json:"creator"`

	Topic *ESDiscussionTopic `json:"topic"`

	Category *ESDiscussionCategory `json:"category,omitempty"`
}

type ESDiscussionCategory struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// Name
	Name *string `json:"name,omitempty" `

	Seq *int `json:"seq"`
}

type DiscussSearchResult struct {
	// 会员数据
	Data []*ESDiscussion `json:"data"`
	// 分页
	Page *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type AllSearchParams struct {
	// 搜索关键词
	KW string `json:"kw"`
}

type AllSearchResult struct {
	Topics      []*ESTopic `json:"topics"`
	TopicsCount int        `json:"topics_count"`

	Articles      []*ESArticle `json:"articles"`
	ArticlesCount int          `json:"articles_count"`

	Accounts      []*ESAccount `json:"accounts"`
	AccountsCount int          `json:"accounts_count"`
}
