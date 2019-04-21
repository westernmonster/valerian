package modules

import (
	"crypto/rsa"
	"sync"
	"time"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/osin"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type authenticator struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

type AuthenticatorValidator interface {
	GenerateAccessToken(data *osin.AccessData, generaterefresh bool) (string, string, error)
	RetrieveJWK() []byte
}

var once sync.Once

// Instance represent a single authenticator instance.
var Instance *authenticator

func NewAuthenticatorValidator() AuthenticatorValidator {
	once.Do(func() {
		Instance = new(authenticator)
		Instance.PrivateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
		Instance.PublicKey, _ = jwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)
	})

	return Instance
}

func (a *authenticator) RetrieveJWK() []byte {
	return publicJWK
}

// generate JWT access token
func (a *authenticator) GenerateAccessToken(data *osin.AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
	logrus.Info(data)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":    time.Now().Unix() + 86400, // Expiration Time
		"iat":    time.Now().Unix(),         // Issued at
		"iss":    "flywk.com",               // Issuer
		"aud":    data.Client.GetId(),       // Audience
		"sub":    data.AccessData.UserData.(infrastructure.CustomUserData).AccountID,
		"scopes": "api:everything",
	})

	token.Header["kid"] = "5fc5f790-da55-4b11-b3fd-6a00d8a93de9"
	accesstoken, err = token.SignedString(a.PrivateKey)
	if err != nil {
		return "", "", err
	}

	if !generaterefresh {
		return
	}

	// generate JWT refresh token
	token = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud": data.Client.GetId(),
	})

	refreshtoken, err = token.SignedString(a.PrivateKey)
	if err != nil {
		return "", "", err
	}
	return
}

var (
	privateKeyPEM = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

	publicKeyPEM = []byte(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAzWRDbbcXXI8Pq6n6/0QQ
EO64Ai9aB36VkHTWLIZARCGCB+IgYmnjCrA5pkqCajEaIzP39ATitS2qLa2fv1st
4hDc1UUoUXBNJxAIdgWYg/dPtlpyxDk8CiUaNSHvnK3idykOfa0dHNygMeRouaYw
VbLEdms3opdwyz7zp9IfPunOliUPs2+Lzx87eoI1elq8+JUbDLR2YXE5a9zKHCg0
/m1eq3ydU00m8rl5geT55r8HNUK1jb+DKiJT0Dxv8MNj0HdzCfFYnlQCg2IccsuJ
+kGCltKNXgGoWfdMd1JTTqKUkqyL9GgVbTH2RaiL4FwONHhoetFB5Jg5myLZawDx
XLGPeiIgMcNCsl/MfhDMOCr4kFtNbst3ilF68rnTYNx4es54lvp3laQMPlklR5qj
e9zDdv4SOgdROOgK5Gl1i0P0PmlvPukSVVpBpewOzZunNSMZbiTUtVyp54/ynYeP
sVw5rPxfbFbfO/yT1ar5ETMYhgUHUP9e1C4CmQvkGa0wMtURpJwViDHtMPZCUL9n
pUrQ1A/rvjaNHDEbhKEpMvzn7t8UHEL2YRncZBL20kDZiH1J8EIRnUr7ZNwfU7wC
ej/SxE012xq/e05EzLA6MNwVUzjqUYEQ/FSdmkS2xlRc8PWiNsltUjGZjEApoGHY
jud2sLn8jR/sm1d1kKUntkcCAwEAAQ==
-----END PUBLIC KEY-----`)
	publicJWK = []byte(`{ kty: 'RSA', use: 'sig', kid: '5fc5f790-da55-4b11-b3fd-6a00d8a93de9', e: 'AQAB', n: 'AM1kQ223F1yPD6up-v9EEBDuuAIvWgd-lZB01iyGQEQhggfiIGJp4wqwOaZKgmoxGiMz9_QE4rUtqi2tn79bLeIQ3NVFKFFwTScQCHYFmIP3T7ZacsQ5PAolGjUh75yt4ncpDn2tHRzcoDHkaLmmMFWyxHZrN6KXcMs-86fSHz7pzpYlD7Nvi88fO3qCNXpavPiVGwy0dmFxOWvcyhwoNP5tXqt8nVNNJvK5eYHk-ea_BzVCtY2_gyoiU9A8b_DDY9B3cwnxWJ5UAoNiHHLLifpBgpbSjV4BqFn3THdSU06ilJKsi_RoFW0x9kWoi-BcDjR4aHrRQeSYOZsi2WsA8Vyxj3oiIDHDQrJfzH4QzDgq-JBbTW7Ld4pRevK502DceHrOeJb6d5WkDD5ZJUeao3vcw3b-EjoHUTjoCuRpdYtD9D5pbz7pElVaQaXsDs2bpzUjGW4k1LVcqeeP8p2Hj7FcOaz8X2xW3zv8k9Wq-REzGIYFB1D_XtQuApkL5BmtMDLVEaScFYgx7TD2QlC_Z6VK0NQP6742jRwxG4ShKTL85-7fFBxC9mEZ3GQS9tJA2Yh9SfBCEZ1K-2TcH1O8Ano_0sRNNdsav3tORMywOjDcFVM46lGBEPxUnZpEtsZUXPD1ojbJbVIxmYxAKaBh2I7ndrC5_I0f7JtXdZClJ7ZH'}`)
)
