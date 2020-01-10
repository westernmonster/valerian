package http

import (
	"valerian/library/net/http/mars"

	"github.com/gobuffalo/packr"
)

func agreement(c *mars.Context) {

	data := packr.NewBox("./static").Bytes("user-agreement.html")

	c.Writer.WriteHeader(200)
	c.Writer.Write(data)
}

func privacy(c *mars.Context) {
	data := packr.NewBox("./static").Bytes("privacy.html")

	c.Writer.WriteHeader(200)
	c.Writer.Write(data)
}
