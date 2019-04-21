package infrastructure

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const TokenSigningKey = "flywk$*^hn"

type CustomUserData struct {
	AccountID int64  `json:"account_id,string"`
	Role      string `json:"role"`
	jwt.StandardClaims
}
