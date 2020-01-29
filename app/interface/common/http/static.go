package http

import (
	"valerian/library/net/http/mars"

	"github.com/gobuffalo/packr"
)

// @Summary User Agreement
// @Description User Agreement
// @Tags editor
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200  "HTML"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /static/user-agreement [get]
func agreement(c *mars.Context) {

	data := packr.NewBox("./static").Bytes("user-agreement.html")

	c.Writer.WriteHeader(200)
	c.Writer.Write(data)
}

// @Summary User Privacy
// @Description User Privacy
// @Tags editor
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200  "HTML"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /static/privacy [get]
func privacy(c *mars.Context) {
	data := packr.NewBox("./static").Bytes("privacy.html")

	c.Writer.WriteHeader(200)
	c.Writer.Write(data)
}
