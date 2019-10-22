package tag_test

import (
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/tag"
)

// This example create a tag middleware instance and attach to a global.
// It will put a tag into Keys field of context by specified policy custom defines
// You can define a custom policy through any request param or http header.
// Register several tag middlewares to put several tags.
func Example() {
	var pf tag.PolicyFunc
	// create your tag policy
	pf = func(ctx *mars.Context) string {
		if ctx.Request.Form.Get("group") == "a" {
			return "a"
		}
		return "b"
	}
	t := tag.New("abtest", pf)

	engine := mars.Default()
	engine.Use(t)
	engine.GET("/abtest", HandlerMap)

	engine.Run(":18080")
}

func HandlerMap(ctx *mars.Context) {
	value, ok := tag.Value(ctx, "abtest")
	if !ok {
		ctx.String(-400, "failed to parse group")
		ctx.Abort()
		return
	}

	if value == "a" {
		HandlerA(ctx)
	}
	if value == "b" {
		HandlerB(ctx)
	}
}

func HandlerA(ctx *mars.Context) {
	// your business
	ctx.String(200, "group a")
	return
}

func HandlerB(ctx *mars.Context) {
	// your business
	ctx.String(200, "group b")
	return
}
