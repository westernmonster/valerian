package model

import (
	"regexp"
	"valerian/library/ecode"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation"
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

func ValidateIdentity(identityType int32, prefix string) *ValidateIdentityRule {
	return &ValidateIdentityRule{
		IdentityType: identityType,
		Prefix:       prefix,
	}
}

type ValidateIdentityRule struct {
	IdentityType int32
	Prefix       string
}

func (p *ValidateIdentityRule) Validate(v interface{}) error {
	identity := v.(string)

	if p.IdentityType == IdentityEmail {
		if !govalidator.IsEmail(identity) {
			return ecode.InvalidEmail
		}
	} else {
		chinaRegex := regexp.MustCompile(ChinaMobileRegex)
		otherRegex := regexp.MustCompile(OtherMobileRegex)

		if p.Prefix == "86" {
			if !chinaRegex.MatchString(identity) {
				return ecode.InvalidMobile
			}
		} else { // China
			if !otherRegex.MatchString(identity) {
				return ecode.InvalidMobile
			}
		} // Other Country
	}

	return nil
}
