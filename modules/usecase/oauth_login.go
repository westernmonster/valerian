package usecase

import (
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/berr"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/infrastructure/gid"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

// EmailLogin 登录
func (p *OauthUsecase) EmailLogin(ctx *biz.BizContext, req *models.EmailLoginReq, ip string) (loginResult *models.LoginResult, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	client, exist, err := p.OauthClientRepository.GetByCondition(tx, map[string]string{
		"client_id": req.ClientID,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		err = berr.Errorf("未找到该client")
		return
	}

	user, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
		"email": req.Email,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}

	if !exist {
		err = berr.Errorf("邮件地址不正确")
		return
	}

	passwordHash, err := hashPassword(req.Password, user.Salt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !strings.EqualFold(user.Password, passwordHash) {
		err = berr.Errorf("密码不正确")
		return
	}

	token, err := p.grantAccessToken(tx, client, user, models.ExpiresIn, "")
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	loginResult = &models.LoginResult{
		AccountID:    user.ID,
		Role:         user.Role,
		AccessToken:  token.Token,
		ExpiresIn:    models.ExpiresIn,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: "",
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

}

// MobileLogin 登录
func (p *OauthUsecase) MobileLogin(ctx *biz.BizContext, req *models.MobileLoginReq, ip string) (loginResult *models.LoginResult, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	client, exist, err := p.OauthClientRepository.GetByCondition(tx, map[string]string{
		"client_id": req.ClientID,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		err = berr.Errorf("未找到该client")
		return
	}

	user, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
		"mobile": req.Prefix + req.Mobile,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}

	if !exist {
		err = berr.Errorf("为找到该手机号")
		return
	}

	passwordHash, err := hashPassword(req.Password, user.Salt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !strings.EqualFold(user.Password, passwordHash) {
		err = berr.Errorf("密码不正确")
		return
	}

	token, err := p.grantAccessToken(tx, client, user, models.ExpiresIn, "")
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	loginResult = &models.LoginResult{
		AccountID:    user.ID,
		Role:         user.Role,
		AccessToken:  token.Token,
		ExpiresIn:    models.ExpiresIn,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: "",
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// DigitLogin 验证码登录
func (p *OauthUsecase) DigitLogin(ctx *biz.BizContext, req *models.DigitLoginReq, ip string) (loginResult *models.LoginResult, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	client, exist, err := p.OauthClientRepository.GetByCondition(tx, map[string]string{
		"client_id": req.ClientID,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		err = berr.Errorf("未找到该client")
		return
	}

	user, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
		"mobile": req.Prefix + req.Mobile,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}

	if !exist {
		err = berr.Errorf("为找到该手机号")
		return
	}

	// Valcode
	mobile := req.Prefix + req.Mobile
	correct, valcodeItem, errValcode := p.ValcodeRepository.IsCodeCorrect(tx, mobile, models.ValcodeLogin, req.Valcode)
	if errValcode != nil {
		err = tracerr.Wrap(errValcode)
		return
	}
	if !correct {
		err = berr.Errorf("验证码不正确或已经使用")
		return
	}
	valcodeItem.Used = 1

	err = p.ValcodeRepository.Update(tx, valcodeItem)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	token, err := p.grantAccessToken(tx, client, user, models.ExpiresIn, "")
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	loginResult = &models.LoginResult{
		AccountID:    user.ID,
		Role:         user.Role,
		AccessToken:  token.Token,
		ExpiresIn:    models.ExpiresIn,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: "",
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *OauthUsecase) grantAccessToken(node sqalx.Node, client *repo.OauthClient, user *repo.Account, expiresIn int, scope string) (token *repo.OauthAccessToken, err error) {
	tx, err := node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	// Delete expired access tokens
	cond := map[string]string{
		"client_id":  client.ClientID,
		"expires_at": strconv.FormatInt(time.Now().Unix(), 10),
		"user_id":    strconv.FormatInt(user.ID, 10),
	}

	err = p.OauthAccessTokenRepository.DeleteByCondition(tx, cond)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	// create new access token
	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	expiredAt := time.Now().UTC().Add(time.Duration(expiresIn) * time.Second).Unix()
	token = &repo.OauthAccessToken{
		ID:        id,
		ClientID:  client.ClientID,
		ExpiresAt: expiredAt,
		Scope:     scope,
		AccountID: user.ID,
	}

	jwtToken, err := p.generateAccessToken(client.ClientID, user.ID, user.Role, expiredAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	token.Token = jwtToken

	err = p.OauthAccessTokenRepository.Insert(tx, token)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if e := tx.Commit(); e != nil {
		err = tracerr.Wrap(e)
		return
	}
	return
}

// generate JWT access token
func (p *OauthUsecase) generateAccessToken(clientID string, accountID int64, role string, expiredAt int64) (accesstoken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, infrastructure.TokenClaims{
		AccountID: accountID,
		Role:      role,

		StandardClaims: jwt.StandardClaims{
			Audience:  clientID,
			ExpiresAt: expiredAt,
			IssuedAt:  time.Now().Unix(),
			Issuer:    models.Issuer, // Issuer
		},
	})

	token.Header["kid"] = "5fc5f790-da55-4b11-b3fd-6a00d8a93de9"
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	accesstoken, err = token.SignedString(key)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

func (p *OauthUsecase) generateRefreshToken(clientID string) (refreshtoken string, err error) {
	// generate JWT refresh token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud": clientID,
	})

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	refreshtoken, err = token.SignedString(key)
	if err != nil {
		err = tracerr.Wrap(err)
		return
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
