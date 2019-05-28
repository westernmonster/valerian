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

type AuthorizationCode struct {
	ID          int64         `db:"id" json:"id,string"`                        // ID ID
	ClientID    string        `db:"client_id" json:"client_id"`                 // ClientID Client ID
	AccountID   int64         `db:"account_id" json:"account_id,string"`        // AccountID Account ID
	Code        string        `db:"code" json:"code"`                           // Code Code
	RedirectURI *string       `db:"redirect_uri" json:"redirect_uri,omitempty"` // RedirectURI Redirect URI
	ExpiresAt   int64         `db:"expires_at" json:"expires_at,string"`        // ExpiresAt 过期时间
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
	RedirectURI  *string       `db:"redirect_uri" json:"redirect_uri,omitempty"` // RedirectURI Redirect URI
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
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
