package proxy

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"

	"valerian/library/log"
	"valerian/library/net/http/mars"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init(nil)
}

func TestProxy(t *testing.T) {
	engine := mars.Default()
	engine.GET("/icon", NewAlways("http://api.flywk.com/x/web-interface/index/icon"))
	engine.POST("/x/web-interface/archive/like", NewAlways("http://api.flywk.com"))

	go engine.Run(":18080")
	defer func() {
		engine.Server().Shutdown(context.TODO())
	}()
	time.Sleep(time.Second)

	req, err := http.NewRequest("GET", "http://127.0.0.1:18080/icon", nil)
	assert.NoError(t, err)
	req.Host = "api.flywk.com"

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	// proxy form request
	form := url.Values{}
	form.Set("arg1", "1")
	form.Set("arg2", "2")
	req, err = http.NewRequest("POST", "http://127.0.0.1:18080/x/web-interface/archive/like?param=test", bytes.NewReader([]byte(form.Encode())))
	assert.NoError(t, err)
	req.Host = "api.flywk.com"

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	// proxy json request
	bs := []byte(`{"arg1": 1, "arg2": 2}`)
	req, err = http.NewRequest("POST", "http://127.0.0.1:18080/x/web-interface/archive/like?param=test", bytes.NewReader(bs))
	assert.NoError(t, err)
	req.Host = "api.flywk.com"
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}

func TestProxyRace(t *testing.T) {
	engine := mars.Default()
	engine.GET("/icon", NewAlways("http://api.flywk.com/x/web-interface/index/icon"))

	go engine.Run(":18080")
	defer func() {
		engine.Server().Shutdown(context.TODO())
	}()
	time.Sleep(time.Second)

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, err := http.NewRequest("GET", "http://127.0.0.1:18080/icon", nil)
			assert.NoError(t, err)
			req.Host = "api.flywk.com"

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()
			assert.Equal(t, 200, resp.StatusCode)
		}()
	}
	wg.Wait()
}

func TestZoneProxy(t *testing.T) {
	engine := mars.Default()
	engine.GET("/icon", NewZoneProxy("sh004", "http://api.flywk.com/x/web-interface/index/icon"), func(ctx *mars.Context) {
		ctx.AbortWithStatus(500)
	})
	engine.GET("/icon2", NewZoneProxy("none", "http://api.flywk.com/x/web-interface/index/icon2"), func(ctx *mars.Context) {
		ctx.AbortWithStatus(200)
	})
	ug := engine.Group("/update", NewZoneProxy("sh004", "http://api.flywk.com"))
	ug.POST("/name", func(ctx *mars.Context) {
		ctx.AbortWithStatus(500)
	})
	ug.POST("/sign", func(ctx *mars.Context) {
		ctx.AbortWithStatus(500)
	})

	go engine.Run(":18080")
	defer func() {
		engine.Server().Shutdown(context.TODO())
	}()
	time.Sleep(time.Second)

	req, err := http.NewRequest("GET", "http://127.0.0.1:18080/icon", nil)
	assert.NoError(t, err)
	req.Host = "api.flywk.com"
	req.Header.Set("X-FLYWK-SLB", "shjd-out-slb")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	req.URL.Path = "/icon2"
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	req.URL.Path = "/update/name"
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	req.URL.Path = "/update/sign"
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}

func BenchmarkProxy(b *testing.B) {
	engine := mars.Default()
	engine.GET("/icon", NewAlways("http://api.flywk.com/x/web-interface/index/icon"))

	go engine.Run(":18080")
	defer func() {
		engine.Server().Shutdown(context.TODO())
	}()
	time.Sleep(time.Second)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req, err := http.NewRequest("GET", "http://127.0.0.1:18080/icon", nil)
			assert.NoError(b, err)
			req.Host = "api.flywk.com"

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(b, err)
			defer resp.Body.Close()
			assert.Equal(b, 200, resp.StatusCode)
		}
	})
}
