package usecase

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strconv"
	"strings"
	"time"

	"valerian/library/database/sqalx"

	"github.com/dgrijalva/jwt-go"
	"github.com/ztrue/tracerr"

	"valerian/infrastructure"
	"valerian/infrastructure/berr"
	"valerian/infrastructure/biz"
	"valerian/infrastructure/ecode"
	"valerian/infrastructure/gid"
	"valerian/models"
	"valerian/modules/repo"
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
		err = berr.Errorf("未找到该手机号")
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
		err = berr.Errorf("未找到该手机号")
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, infrastructure.TokenClaims{
		AccountID: accountID,
		Role:      role,

		StandardClaims: jwt.StandardClaims{
			Audience:  clientID,
			ExpiresAt: expiredAt,
			IssuedAt:  time.Now().Unix(),
			Issuer:    models.Issuer, // Issuer
		},
	})

	accesstoken, err = token.SignedString([]byte(models.JWTSignKey))
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

func (p *OauthUsecase) generateRefreshToken(clientID string) (refreshtoken string, err error) {
	// generate JWT refresh token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"aud": clientID,
	})

	refreshtoken, err = token.SignedString([]byte(models.JWTSignKey))
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *OauthUsecase) GetTokenInfo(tokenStr string) (claims *infrastructure.TokenClaims, err error) {
	claims = new(infrastructure.TokenClaims)
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, tracerr.Errorf("invalid token")
		}
		return []byte(models.JWTSignKey), nil
	})

	if err != nil {
		return
	}

	if !token.Valid {
		return nil, ecode.NoLogin
	}

	if c, ok := token.Claims.(*infrastructure.TokenClaims); ok {
		if c.Valid() == nil {
			claims = c
			return
		}
	}
	return nil, ecode.NoLogin

}

func convertKey(key string) *rsa.PublicKey {
	certPEM := key
	certPEM = strings.Replace(certPEM, "\\n", "\n", -1)
	certPEM = strings.Replace(certPEM, "\"", "", -1)
	block, _ := pem.Decode([]byte(certPEM))
	cert, _ := x509.ParseCertificate(block.Bytes)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	return rsaPublicKey
}
