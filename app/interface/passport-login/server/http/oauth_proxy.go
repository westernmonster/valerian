package http

import (
	"valerian/app/interface/passport-login/model"
	"valerian/library/net/http/mars"
)

func proxyEmailLogin(c mars.Context) {
	arg := new(model.ArgEmailLogin)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
	}
}
