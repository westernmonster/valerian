package proxy_test

import (
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/proxy"
)

// This example create several reverse proxy to show how to use proxy middleware.
// We proxy three path to `api.stonote.cn` and return response without any changes.
func Example() {
	proxies := map[string]string{
		"/index":        "http://api.stonote.cn/html/index",
		"/ping":         "http://api.stonote.cn/api/ping",
		"/api/versions": "http://api.stonote.cn/api/web/versions",
	}

	engine := mars.Default()
	for path, ep := range proxies {
		engine.GET(path, proxy.NewAlways(ep))
	}
	engine.Run(":18080")
}

// This example create several reverse proxy to show how to use jd proxy middleware.
// The request will be proxied to destination only when request is from specified datacenter.
func ExampleNewZoneProxy() {
	proxies := map[string]string{
		"/index":        "http://api.stonote.cn/html/index",
		"/ping":         "http://api.stonote.cn/api/ping",
		"/api/versions": "http://api.stonote.cn/api/web/versions",
	}

	engine := mars.Default()
	// proxy to specified destination
	for path, ep := range proxies {
		engine.GET(path, proxy.NewZoneProxy("sh004", ep), func(ctx *mars.Context) {
			ctx.String(200, "Origin")
		})
	}
	// proxy with request path
	ug := engine.Group("/update", proxy.NewZoneProxy("sh004", "http://sh001-api.stonote.cn"))
	ug.POST("/name", func(ctx *mars.Context) {
		ctx.String(500, "Should not be accessed")
	})
	ug.POST("/sign", func(ctx *mars.Context) {
		ctx.String(500, "Should not be accessed")
	})
	engine.Run(":18080")
}
