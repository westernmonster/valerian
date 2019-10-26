package model

import "valerian/library/database/sqlx/types"

type AccessToken struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	ClientID  string        `db:"client_id" json:"client_id"`          // ClientID Client ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID Account ID
	Token     string        `db:"token" json:"token"`                  // Token Token
	ExpiresAt int64         `db:"expires_at" json:"expires_at"`        // ExpiresAt 过期时间
	Scope     string        `db:"scope" json:"scope"`                  // Scope Scope
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type AuthorizationCode struct {
	ID          int64         `db:"id" json:"id,string"`                        // ID ID
	ClientID    string        `db:"client_id" json:"client_id"`                 // ClientID Client ID
	AccountID   int64         `db:"account_id" json:"account_id,string"`        // AccountID Account ID
	Code        string        `db:"code" json:"code"`                           // Code Code
	RedirectURI string        `db:"redirect_uri" json:"redirect_uri,omitempty"` // RedirectURI Redirect URI
	ExpiresAt   int64         `db:"expires_at" json:"expires_at"`               // ExpiresAt 过期时间
	Scope       string        `db:"scope" json:"scope"`                         // Scope Scope
	Deleted     types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type Client struct {
	ID           int64         `db:"id" json:"id,string"`                        // ID ID
	ClientID     string        `db:"client_id" json:"client_id"`                 // ClientID Client ID
	ClientSecret string        `db:"client_secret" json:"client_secret"`         // ClientSecret Client Secret
	Extra        string        `db:"extra" json:"extra"`                         // Extra Extra
	RedirectURI  string        `db:"redirect_uri" json:"redirect_uri,omitempty"` // RedirectURI Redirect URI
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type RefreshToken struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	ClientID  string        `db:"client_id" json:"client_id"`          // ClientID Client ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID Account ID
	Token     string        `db:"token" json:"token"`                  // Token Token
	ExpiresAt int64         `db:"expires_at" json:"expires_at"`        // ExpiresAt 过期时间
	Scope     string        `db:"scope" json:"scope"`                  // Scope Scope
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type Role struct {
	ID        string        `db:"id" json:"id"`                 // ID Role ID
	Name      string        `db:"name" json:"name"`             // Name Role Name
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type Scope struct {
	ID          int64         `db:"id" json:"id,string"`            // ID ID
	Scope       string        `db:"scope" json:"scope"`             // Scope Scope
	Description string        `db:"description" json:"description"` // Description Description
	IsDefault   int           `db:"is_default" json:"is_default"`   // IsDefault 是否默认, 0 否，1 是
	Deleted     types.BitBool `db:"deleted" json:"deleted"`         // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`   // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`   // UpdatedAt 更新时间
}

type Account struct {
	ID           int64         `db:"id" json:"id,string"`              // ID ID
	Mobile       string        `db:"mobile" json:"mobile"`             // Mobile 手机
	UserName     string        `db:"user_name" json:"user_name"`       // UserName 用户名
	Email        string        `db:"email" json:"email"`               // Email 邮件地址
	Password     string        `db:"password" json:"password"`         // Password 密码hash
	Role         string        `db:"role" json:"role"`                 // Role 角色
	Salt         string        `db:"salt" json:"salt"`                 // Salt 盐
	Gender       int32         `db:"gender" json:"gender"`             // Gender 性别
	BirthYear    int32         `db:"birth_year" json:"birth_year"`     // BirthYear 出生年
	BirthMonth   int32         `db:"birth_month" json:"birth_month"`   // BirthMonth 出生月
	BirthDay     int32         `db:"birth_day" json:"birth_day"`       // BirthDay 出生日
	Location     int64         `db:"location" json:"location,string"`  // Location 地区
	Introduction string        `db:"introduction" json:"introduction"` // Introduction 自我介绍
	Avatar       string        `db:"avatar" json:"avatar"`             // Avatar 头像
	Source       int32         `db:"source" json:"source"`             // Source 注册来源
	IP           int64         `db:"ip" json:"ip,string"`              // IP 注册IP
	IDCert       types.BitBool `db:"id_cert" json:"id_cert"`           // IDCert 是否身份认证
	WorkCert     types.BitBool `db:"work_cert" json:"work_cert"`       // WorkCert 是否工作认证
	IsOrg        types.BitBool `db:"is_org" json:"is_org"`             // IsOrg 是否机构用户
	IsVip        types.BitBool `db:"is_vip" json:"is_vip"`             // IsVip 是否VIP用户
	Deleted      types.BitBool `db:"deleted" json:"deleted"`           // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`     // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`     // UpdatedAt 更新时间
	Prefix       string        `db:"prefix" json:"prefix"`             // Prefix 手机前缀
}
