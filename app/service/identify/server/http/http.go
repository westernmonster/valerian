package http

import (
	"net/http"
	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/conf"
	"valerian/app/service/identify/service"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

var (
	srv *service.Service
)

func Init(c *conf.Config, s *service.Service) {
	srv = s
	engine := mars.DefaultServer(c.Mars)

	innerRouter(engine)
	if err := engine.Start(); err != nil {
		log.Errorf("engine.Start error(%v)", err)
		panic(err)
	}
}

func innerRouter(e *mars.Engine) {
	e.Ping(ping)
	e.Register(register)
	group := e.Group("/x/internal/identify")
	{
		group.GET("token", accessToken)
	}
}

func accessToken(c *mars.Context) {
	token := new(api.TokenReq)
	if err := c.Bind(token); err != nil {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	res, err := srv.GetTokenInfo(c, token.Token)
	if err == nil {
		c.Set("aid", res.Aid)
	}
	c.JSON(res, err)
}

// ping check server ok.
func ping(c *mars.Context) {
	var (
		err error
	)
	if err = srv.Ping(c); err != nil {
		log.Errorf("config service ping error(%v)", err)
		c.JSON(nil, err)
		http.Error(c.Writer, "", http.StatusServiceUnavailable)
	}
}

// register support discovery.
func register(c *mars.Context) {
	c.JSON(map[string]struct{}{}, nil)
}
