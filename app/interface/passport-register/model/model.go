package model

import "valerian/library/database/sqlx/types"

type AccessToken struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	ClientID  string        `db:"client_id" json:"client_id"`          // ClientID Client ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID Account ID
	Token     string        `db:"token" json:"token"`                  // Token Token
	ExpiresAt int64         `db:"expires_at" json:"expires_at,string"` // ExpiresAt 过期时间
	Scope     string        `db:"scope" json:"scope"`                  // Scope Scope
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type RefreshToken struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	ClientID  string        `db:"client_id" json:"client_id"`          // ClientID Client ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID Account ID
	Token     string        `db:"token" json:"token"`                  // Token Token
	ExpiresAt int64         `db:"expires_at" json:"expires_at,string"` // ExpiresAt 过期时间
	Scope     string        `db:"scope" json:"scope"`                  // Scope Scope
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type AccountRole struct {
	ID        string        `db:"id" json:"id"`                 // ID 角色ID
	Name      string        `db:"name" json:"name"`             // Name 角色名
	Perms     string        `db:"perms" json:"perms"`           // Perms 权限
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type Client struct {
	ID           int64         `db:"id" json:"id,string"`                        // ID ID
	ClientID     string        `db:"client_id" json:"client_id"`                 // ClientID Client ID
	ClientSecret string        `db:"client_secret" json:"client_secret"`         // ClientSecret Client Secret
	Extra        string        `db:"extra" json:"extra"`                         // Extra Extra
	RedirectURI  *string       `db:"redirect_uri" json:"redirect_uri,omitempty"` // RedirectURI Redirect URI
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type Account struct {
	ID           int64         `db:"id" json:"id,string"`                        // ID ID
	Mobile       string        `db:"mobile" json:"mobile"`                       // Mobile 手机
	Email        string        `db:"email" json:"email"`                         // Email 邮件地址
	UserName     string        `db:"user_name" json:"user_name"`                 // UserName 用户名
	Password     string        `db:"password" json:"password"`                   // Password 密码hash
	Role         string        `db:"role" json:"role"`                           // Role 角色
	Salt         string        `db:"salt" json:"salt"`                           // Salt 盐
	Gender       *int          `db:"gender" json:"gender,omitempty"`             // Gender 性别
	BirthYear    *int          `db:"birth_year" json:"birth_year,omitempty"`     // BirthYear 出生年
	BirthMonth   *int          `db:"birth_month" json:"birth_month,omitempty"`   // BirthMonth 出生月
	BirthDay     *int          `db:"birth_day" json:"birth_day,omitempty"`       // BirthDay 出生日
	Location     *int64        `db:"location" json:"location,omitempty,string"`  // Location 地区
	Introduction *string       `db:"introduction" json:"introduction,omitempty"` // Introduction 自我介绍
	Avatar       string        `db:"avatar" json:"avatar"`                       // Avatar 头像
	Source       int           `db:"source" json:"source"`                       // Source 注册来源
	IP           int64         `db:"ip" json:"ip,string"`                        // IP 注册IP
	IDCert       types.BitBool `db:"id_cert" json:"id_cert"`                     // IDCert 是否身份认证
	WorkCert     types.BitBool `db:"work_cert" json:"work_cert"`                 // WorkCert 是否工作认证
	IsOrg        types.BitBool `db:"is_org" json:"is_org"`                       // IsOrg 是否机构用户
	IsVIP        types.BitBool `db:"is_vip" json:"is_vip"`                       // IsVIP 是否VIP用户
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type Area struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Name      string        `db:"name" json:"name"`             // Name 名称
	Code      string        `db:"code" json:"code"`             // Code 编码
	Type      string        `db:"type" json:"type"`             // Type 编码
	Parent    int64         `db:"parent" json:"parent,string"`  // Parent 父级ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type CountryCode struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Name      string        `db:"name" json:"name"`             // Name 国家名
	Emoji     string        `db:"emoji" json:"emoji"`           // Emoji 国旗
	CnName    string        `db:"cn_name" json:"cn_name"`       // CnName 国家中文名
	Code      string        `db:"code" json:"code"`             // Code 编码
	Prefix    string        `db:"prefix" json:"prefix"`         // Prefix 前缀
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type AccountStat struct {
	AccountID       int64 `db:"account_id" json:"account_id,string"`      // AccountID 用户ID
	Following       int   `db:"following" json:"following"`               // Following 关注数
	Fans            int   `db:"fans" json:"fans"`                         // Fans 粉丝数
	ArticleCount    int   `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int   `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	TopicCount      int   `db:"topic_count" json:"topic_count"`           // TopicCount 讨论数
	Black           int   `db:"black" json:"black"`                       // Black 黑名单数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type MessageStat struct {
	AccountID   int64 `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	UnreadCount int   `db:"unread_count" json:"unread_count"`    // UnreadCount 未读数
	CreatedAt   int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
