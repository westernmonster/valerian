package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgRenewToken struct {
	RefreshToken string `json:"refresh_token"`
	ClientID     string `json:"client_id"`
}

func (p *ArgRenewToken) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ClientID,
			validation.Required.Error(`"client_id" is required`),
		),
		validation.Field(&p.RefreshToken,
			validation.Required.Error(`"refresh_token" is required`),
		),
	)

}

type ArgAuthorize struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

func (p *ArgAuthorize) Validate() error {
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

type ArgGrantTypeAuthorizationCode struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
}

type ArgGrantTypePassword struct {
	GrantType string `json:"grant_type"`
	ClientID  string `json:"client_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Scope     string `json:"scope"`
}

type ArgGrantTypeClientCredentials struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type ArgGrantTypeRefreshToken struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

type ArgLogout struct {
	ClientID string `json:"client_id"`
}

func (p *ArgLogout) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ClientID, validation.Required),
	)

}

type ArgCloseAccount struct {
	// 验证码 6位数字
	Valcode string `json:"valcode"`
}

func (p *ArgCloseAccount) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Valcode, validation.Required, validation.RuneLength(6, 6), is.Digit),
	)
}
