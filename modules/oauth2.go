package modules

import (
	"fmt"
	"net/http"
	"sync"

	"git.flywk.com/flywiki/api/infrastructure/db"
	"git.flywk.com/flywiki/api/infrastructure/osin"
	"git.flywk.com/flywiki/api/modules/repo"
	"git.flywk.com/flywiki/api/modules/usecase"
	"github.com/gin-gonic/gin"
)

var (
	onceServer  sync.Once
	oauthServer *osin.Server
)

func StartOAUTHServer() {
	onceServer.Do(func() {
		db, node, err := db.InitDatabase()
		if err != nil {
			panic(err)
			return
		}
		sconfig := osin.NewServerConfig()
		sconfig.AccessExpiration = 3600
		sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
		sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE,
			osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS}
		sconfig.AllowGetAccessRequest = true
		sconfig.AllowClientSecretInParams = true
		//sconfig.RequirePKCEForPublicClients = true // with a clientID and secret it's enough

		osinStorage := &usecase.OAuthStorage{
			Node:                    node,
			DB:                      db,
			AuthClientRepository:    &repo.AuthClientRepository{},
			AuthAccessRepository:    &repo.AuthAccessRepository{},
			AuthRefreshRepository:   &repo.AuthRefreshRepository{},
			AuthExpiresRepository:   &repo.AuthExpiresRepository{},
			AuthAuthorizeRepository: &repo.AuthAuthorizeRepository{},
		}

		oauthServer = osin.NewServer(sconfig, osinStorage)
		authService := NewAuthenticatorValidator()
		oauthServer.AccessTokenGen = authService
	})

}

func HandleAuthorizeRequest(c *gin.Context) {
	resp := oauthServer.NewResponse()
	defer resp.Close()

	if ar := oauthServer.HandleAuthorizeRequest(resp, c.Request); ar != nil {
		if !HandleLoginPage(ar, c.Writer, c.Request) {
			return
		}
		ar.UserData = struct{ Login string }{Login: "test"}
		ar.Authorized = true
		oauthServer.FinishAuthorizeRequest(resp, c.Request, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		resp.Output["custom_parameter"] = 187723
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
}

func HandleLoginPage(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	if r.Method == "POST" && r.FormValue("login") == "test" && r.FormValue("password") == "test" {
		return true
	}

	return false
}

func HandleTokenRequest(c *gin.Context) {
	resp := oauthServer.NewResponse()
	defer resp.Close()

	if ar := oauthServer.HandleAccessRequest(resp, c.Request); ar != nil {
		switch ar.Type {
		case osin.AUTHORIZATION_CODE:
			ar.Authorized = true
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		case osin.PASSWORD:
			if ar.Username == "test" && ar.Password == "test" {
				ar.Authorized = true
			}
		case osin.CLIENT_CREDENTIALS:
			ar.Authorized = true
		case osin.ASSERTION:
			if ar.AssertionType == "urn:osin.example.complete" && ar.Assertion == "osin.data" {
				ar.Authorized = true
			}
		}
		oauthServer.FinishAccessRequest(resp, c.Request, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		resp.Output["custom_parameter"] = 19923
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
}
