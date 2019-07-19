package http

import (
	"io"
	"strconv"
	"strings"
	"valerian/app/infra/config/conf"
	"valerian/app/infra/config/service"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/antispam"

	"github.com/dgryski/go-farm"
)

var (
	cnf *conf.Config
	// verify   *v.Verify
	confSvc2 *service.Service
	anti     *antispam.Antispam
)

// Init init.
func Init(c *conf.Config, srv *service.Service, rpcCloser io.Closer) {
	initService(c)
	// verify = v.New(c.Verify)
	cnf = c
	// confSvc = s
	confSvc2 = srv
	engine := mars.DefaultServer(c.Mars)
	innerRouter(engine)
	if err := engine.Start(); err != nil {
		log.Errorf("engine.Start() error(%v)", err)
		panic(err)
	}
}

// innerRouter init inner router.
func innerRouter(e *mars.Engine) {
	e.Ping(ping)
	e.Register(register)
	// b := e.Group("/", verify.Verify)
	noAuth := e.Group("/")
	{
		noAuth.GET("config/versions", versions)
		noAuth.GET("config/builds", builds)
		noAuth.GET("config/check", check)
		noAuth.GET("config/get", setMid, anti.ServeHTTP, config)
		noAuth.GET("config/file.so", file)
		noAuth.GET("config/latest", latest)
		noAuth.POST("config/host/clear", clearhost)
	}
}
func setMid(c *mars.Context) {
	var (
		token   string
		service string
		query   = c.Request.URL.Query()
		hash    uint64
	)
	service = query.Get("service")
	if service == "" {
		token = query.Get("token")
		if token == "" {
			c.JSON(nil, ecode.RequestErr)
			c.Abort()
			return
		}
		hash = farm.Hash64([]byte(token))
	} else {
		arrs := strings.Split(service, "_")
		if len(arrs) != 3 {
			c.JSON(nil, ecode.RequestErr)
			c.Abort()
			return
		}
		_, err := strconv.ParseInt(arrs[0], 10, 64)
		if err != nil {
			c.JSON(nil, ecode.RequestErr)
			c.Abort()
			return
		}
		hash = farm.Hash64([]byte(service))
	}
	c.Set("mid", int64(hash))
}

func initService(c *conf.Config) {
	anti = antispam.New(c.Antispam)
}
