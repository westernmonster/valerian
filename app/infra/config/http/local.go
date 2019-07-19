package http

import (
	"net/http"

	"valerian/library/log"
	"valerian/library/net/http/mars"
)

// ping check server ok.
func ping(c *mars.Context) {
	var (
		err error
	)
	if err = confSvc2.Ping(c); err != nil {
		log.Errorf("config service ping error(%v)", err)
		c.JSON(nil, err)
		http.Error(c.Writer, "", http.StatusServiceUnavailable)
	}
}

// register check server ok.
func register(c *mars.Context) {
	c.JSON(map[string]struct{}{}, nil)
}
