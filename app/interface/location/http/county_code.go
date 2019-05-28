package http

import "valerian/library/net/http/mars"

func countryCodes(c *mars.Context) {
	c.JSON(srv.GetAllCountryCodes(c))
}
