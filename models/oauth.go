package models

import validation "github.com/go-ozzo/ozzo-validation"

type AuthorizeReq struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

func (p *AuthorizeReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ClientID,
			validation.Required.Error(`"client_id" is required`),
		),
		validation.Field(&p.RedirectURI,
			validation.Required.Error(`"redirect_uri" is required`),
		),
		validation.Field(&p.ResponseType,
			validation.Required.Error(`"response_type" is required`),
			validation.In(ResponseTypeCode, ResponseTypeToken).Error(`"response_type" is not allowed`)),
	)

}

type GrantTypeAuthorizationCodeReq struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
}

type GrantTypePasswordReq struct {
	GrantType string `json:"grant_type"`
	ClientID  string `json:"client_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Scope     string `json:"scope"`
}

type GrantTypeClientCredentialsReq struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GrantTypeRefreshTokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}