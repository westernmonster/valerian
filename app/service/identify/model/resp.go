package model

type TokenResp struct {
	// 用户ID
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
	// Access Token， 请在 HTTP 请求头中添加
	// 例子： Authorization: Bearer  TJVA95OrM7E20RMHrHDcEfxjoYZgeFONFh7HgQ
	AccessToken string `json:"access_token"`
	// 有效期 秒为单位
	ExpiresIn int `json:"expires_in"`
	// Token 类型，默认为 Bearer
	TokenType string `json:"token_type"`
	// Refresh Token 暂不使用
	RefreshToken string `json:"refresh_token,omitempty"`
}
