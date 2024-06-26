package antispam

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"valerian/library/cache/redis"
	"valerian/library/net/http/mars"
	xtime "valerian/library/time"
)

func TestAntiSpamHandler(t *testing.T) {
	anti := New(
		&Config{
			On:     true,
			Second: 1,
			N:      1,
			Hour:   1,
			M:      1,
			Redis: &redis.Config{
				MaxActive:    10,
				MaxIdle:      10,
				IdleTimeout:  xtime.Duration(time.Second * 60),
				Name:         "test",
				Proto:        "tcp",
				Addr:         "127.0.0.1:6889",
				DialTimeout:  xtime.Duration(time.Second),
				ReadTimeout:  xtime.Duration(time.Second),
				WriteTimeout: xtime.Duration(time.Second),
			},
		},
	)

	engine := mars.New()
	engine.UseFunc(func(c *mars.Context) {
		mid, _ := strconv.ParseInt(c.Request.Form.Get("mid"), 10, 64)
		c.Set("mid", mid)
		c.Next()
	})
	engine.Use(anti.Handler())
	engine.GET("/antispam", func(c *mars.Context) {
		c.String(200, "pass")
	})
	go engine.Run(":18080")

	time.Sleep(time.Millisecond * 50)
	code, content, err := httpGet("http://127.0.0.1:18080/antispam?mid=11")
	if err != nil {
		t.Logf("http get failed, err:=%v", err)
		t.FailNow()
	}
	if code != 200 || string(content) != "pass" {
		t.Logf("request should pass by limiter, but blocked: %d, %v", code, content)
		t.FailNow()
	}

	_, content, err = httpGet("http://127.0.0.1:18080/antispam?mid=11")
	if err != nil {
		t.Logf("http get failed, err:=%v", err)
		t.FailNow()
	}
	if string(content) == "pass" {
		t.Logf("request should block by limiter, but passed")
		t.FailNow()
	}

	engine.Server().Shutdown(context.TODO())
}

func httpGet(url string) (code int, content []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	code = resp.StatusCode
	return
}

func TestConfigValidate(t *testing.T) {
	var conf *Config
	assert.Contains(t, conf.validate().Error(), "empty config")

	conf = &Config{
		Second: 0,
	}
	assert.Contains(t, conf.validate().Error(), "invalid Second")

	conf = &Config{
		Second: 1,
		Hour:   0,
	}
	assert.Contains(t, conf.validate().Error(), "invalid Hour")
}
