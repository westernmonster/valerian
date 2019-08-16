package model

import (
	"encoding/json"
	"valerian/library/database/sqlx/types"

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

type AccountSearchParams struct {
	Bsp     *BasicSearchParams
	TopicID *int64 `json:"topic_id,string,omitempty" swaggertype:"string"`
}

func (p *AccountSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Bsp),
	)
}

type TopicSearchParams struct {
	Bsp     *BasicSearchParams
	TopicID int64 `json:"topic_id,string,omitempty" swaggertype:"string"`
	// Query string `json:"query"`
}

func (p *TopicSearchParams) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required),
		validation.Field(&p.Bsp),
	)
}

type ESAccount struct {
	// ID
	ID *int64 `json:"id,string,omitempty" swaggertype:"string"`

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
	IDCert *types.BitBool `json:"id_cert,omitempty"`

	// 是否工作认证
	WorkCert *types.BitBool `json:"work_cert,omitempty"`

	// 是否机构用户
	IsOrg *types.BitBool `json:"is_org,omitempty"`

	// 是否VIP用户
	IsVIP *types.BitBool `json:"is_vip,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,string,omitempty" swaggertype:"string"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,string,omitempty"  swaggertype:"string"`

	// 是否话题成员
	IsTopicMember bool `json:"is_topic_member"`
}

type AccountSearchResult struct {
	// 排序顺序
	Order string `json:"order"`
	// 排序的字段
	Sort string `json:"sort"`
	// 会员数据
	Data []*ESAccount `json:"data"`
	// 分页
	Page *Page `json:"page"`
	// 调试
	Debug string `json:"debug"`
}

type ESTopicMember struct {
	// ID ID
	ID *int64 `json:"id,string,omitempty"  swaggertype:"string"`
	// 用户名
	UserName *string `json:"user_name,omitempty"`
	// 手机
	Mobile *string `json:"mobile,omitempty"`

	// 邮件地址
	Email *string `json:"email,omitempty"`

	// 头像
	Avatar *string `json:"avatar,omitempty"`
}

type ESTopic struct {
	// ID ID
	ID *int64 `json:"id,string,omitempty"  swaggertype:"string"`
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
	CreatedBy *ESTopicMember `json:"created_by,omitempty"`
	// CreatedAt 创建时间
	CreatedAt *int64 `json:"created_at,string"  swaggertype:"string"`
	// UpdatedAt 更新时间
	UpdatedAt *int64 `json:"updated_at,string"  swaggertype:"string"`

	// 是否已经授权
	IsAuthed bool `json:"is_authed"`
}

type TopicSearchResult struct {
	// 排序顺序
	Order string `json:"order"`
	// 排序的字段
	Sort string `json:"sort"`
	// 会员数据
	Data []*ESTopic `json:"data"`
	// 分页
	Page *Page `json:"page"`
	// 调试
	Debug string `json:"debug"`
}
