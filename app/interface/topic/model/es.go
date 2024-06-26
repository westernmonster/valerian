package model

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
)

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
	CreatedAt *int64 `json:"created_at,omitempty"`
	// UpdatedAt 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`
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
	Gender *int32 `json:"gender,omitempty"`

	// 出生年
	BirthYear *int32 `json:"birth_year,omitempty"`

	// 出生月
	BirthMonth *int32 `json:"birth_month,omitempty"`

	// 出生日
	BirthDay *int32 `json:"birth_day,omitempty"`

	// 地区
	Location *int64 `json:"location,omitempty" swaggertype:"string"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`

	// 头像
	Avatar *string `json:"avatar,omitempty"`

	// 注册来源
	Source *int32 `json:"source,omitempty"`

	// 是否身份认证
	IDCert *bool `json:"id_cert,omitempty"`

	// 是否工作认证
	WorkCert *bool `json:"work_cert,omitempty"`

	// 是否机构用户
	IsOrg *bool `json:"is_org,omitempty"`

	// 是否VIP用户
	IsVIP *bool `json:"is_vip,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,omitempty" swaggertype:"string"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"  swaggertype:"string"`
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

type TopicSearchResult struct {
	// 会员数据
	Data []*ESTopic `json:"data"`
	// 分页
	Paging *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
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

type Page struct {
	// 页码
	Pn int32 `json:"num"`
	// 页大小
	Ps int32 `json:"size"`
	// 统计数量
	Total int64 `json:"total"`
}

type ESPage struct {
	// 页码
	Pn int32 `json:"num"`
	// 页大小
	Ps int32 `json:"size"`
	// 统计数量
	Total int64 `json:"total"`
}

type SearchResult struct {
	Order  string            `json:"order"`
	Sort   string            `json:"sort"`
	Result []json.RawMessage `json:"data"`
	Page   *ESPage           `json:"page"`
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
	Pn int32 `json:"pn" form:"pn"`
	// 每页大小
	Ps int32 `json:"ps" form:"ps"`
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
