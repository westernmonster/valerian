package auth

import (
	"net/http"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"git.flywk.com/flywiki/api/infrastructure/bootstrap"
	"git.flywk.com/flywiki/api/infrastructure/oauth2"
	"git.flywk.com/flywiki/api/infrastructure/oauth2/generates"
	"git.flywk.com/flywiki/api/infrastructure/oauth2/manage"
	oauth_models "git.flywk.com/flywiki/api/infrastructure/oauth2/models"
	"git.flywk.com/flywiki/api/infrastructure/oauth2/server"
	"git.flywk.com/flywiki/api/infrastructure/oauth2/store"
	"git.flywk.com/flywiki/api/models"
)

var (
	gServer *server.Server
	once    sync.Once
)

// InitServer Initialize the service
func InitServer(manager oauth2.Manager) *server.Server {
	once.Do(func() {
		gServer = server.NewDefaultServer(manager)
	})
	return gServer
}

// HandleAuthorizeRequest the authorization request handling
func HandleAuthorizeRequest(c *gin.Context) {
	err := gServer.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

// HandleTokenRequest token request handling
func HandleTokenRequest(c *gin.Context) {
	err := gServer.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

func Configure(p *bootstrap.Bootstrapper) {
	// OAUTH 2
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// Generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(models.JWTKey), jwt.SigningMethodHS512))

	// Client: 目前 OAuth2 只开放给自有服务，所以使用 password grant, client 也写死
	clientStore := store.NewClientStore()
	clientStore.Set(models.OAUTH2MobileClientID, &oauth_models.Client{
		ID:     models.OAUTH2MobileClientID,
		Secret: models.OAUTH2MobileClientSecret,
		Domain: models.OAUTH2MobileClientDomain,
	})
	clientStore.Set(models.OAUTH2WebClientID, &oauth_models.Client{
		ID:     models.OAUTH2WebClientID,
		Secret: models.OAUTH2WebClientSecret,
		Domain: models.OAUTH2WebClientDomain,
	})

	manager.MapClientStorage(clientStore)

	// // Initialize the gin oauth2 service
	InitServer(manager)
	SetAllowGetAccessRequest(true)
	SetClientInfoHandler(server.ClientFormHandler)
}
