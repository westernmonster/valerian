package repo

type OauthAuthorizationCode struct {
	ID          int64   `db:"id" json:"id,string"`                        // ID ID
	ClientID    string  `db:"client_id" json:"client_id"`                 // ClientID Client ID
	AccountID   int64   `db:"account_id" json:"account_id,string"`        // AccountID Account ID
	Code        string  `db:"code" json:"code"`                           // Code Code
	RedirectURI *string `db:"redirect_uri" json:"redirect_uri,omitempty"` // RedirectURI Redirect URI
	ExpiresAt   int64   `db:"expires_at" json:"expires_at,string"`        // ExpiresAt 过期时间
	Scope       string  `db:"scope" json:"scope"`                         // Scope Scope
	Deleted     int     `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt   int64   `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt   int64   `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}
