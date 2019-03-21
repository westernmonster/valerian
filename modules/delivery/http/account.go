package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

type AccountCtrl struct {
	infrastructure.BaseCtrl

	AccountUsecase interface {
		GetByID(userID int64) (item *repo.Account, err error)
		DoLogin(req *models.LoginReq, ip string) (item *repo.Account, err error)
	}
}

func GetBearer(auth []string) (jwt string, ok bool) {
	for _, v := range auth {
		ret := strings.Split(v, " ")
		if len(ret) == 2 && ret[0] == "Bearer" {
			return ret[1], true
		}
	}
	return "", false
}

// Auth godoc
func (p *AccountCtrl) Auth(ctx *gin.Context) {
	tokenStr := ""

	if v, ok := GetBearer(ctx.Request.Header["Authorization"]); ok {
		tokenStr = v
	} else {
		cookieValue, err := ctx.Cookie("token")
		if err == nil {
			tokenStr = cookieValue
		}
	}

	if tokenStr == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"success": false,
			"message": "jwt token is not set",
		})
		return
	}

	claims := new(infrastructure.TokenClaims)

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(infrastructure.TokenSigningKey), nil
	})

	if err != nil {
		logrus.Error(fmt.Sprintf("Error parsing token: %v", err))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"success": false,
			"message": "服务器端发生错误",
		})
		return
	}

	if claims, ok := token.Claims.(*infrastructure.TokenClaims); ok && token.Valid {
		account, err := p.AccountUsecase.GetByID(claims.AccountID)
		if err != nil {
			logrus.Error(fmt.Sprintf("Error on fetch user info: %v", err))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"success": false,
				"message": "服务器端发生错误",
			})
			return
		}

		ctx.Set("AccountID", account.ID)
		ctx.Set("Role", claims.Role)
		ctx.Next()
	} else {
		logrus.Error(fmt.Sprintf("Error parsing token: %v", err))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"success": false,
			"message": "服务器端发生错误",
		})
		return
	}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags session
// @Accept json
// @Produce json
// @Param req body models.LoginReq true "用户登录"
// @Success 200 {object} models.LoginResult "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /session [post]
func (p *AccountCtrl) Login(ctx *gin.Context) {
	req := new(models.LoginReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "验证失败，请检查您的输入",
		})

		return
	}

	ip := ctx.ClientIP()
	user, err := p.AccountUsecase.DoLogin(req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	cookie, err := p.createCookie(user)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)
	p.SuccessResp(ctx, models.LoginResult{
		Token: cookie.Value,
		Role:  "user",
	})

	return
}

func (p *AccountCtrl) createCookie(user *repo.Account) (cookie *http.Cookie, err error) {
	age := 90 * 24 * 60 * 60 // 90 days, present by second
	cookie = new(http.Cookie)
	cookie.Name = "token"
	cookie.MaxAge = age
	cookie.HttpOnly = true

	claims := infrastructure.TokenClaims{
		user.ID,
		"role",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 90).Unix(),
			Issuer:    "web",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(infrastructure.TokenSigningKey))
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	cookie.Value = tokenString

	return

}
