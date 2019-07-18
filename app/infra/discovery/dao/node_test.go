package dao

import (
	"context"
	"strings"
	"testing"
	"time"

	dc "valerian/app/infra/discovery/conf"
	"valerian/app/infra/discovery/model"
	"valerian/library/ecode"
	mars "valerian/library/net/http/mars"
	"valerian/library/net/netutil/breaker"
	xtime "valerian/library/time"

	. "github.com/smartystreets/goconvey/convey"
	gock "gopkg.in/h2non/gock.v1"
)

func TestCall(t *testing.T) {
	Convey("test call", t, func() {
		var res *model.Instance
		node := newNode(&dc.Config{HTTPClient: &mars.ClientConfig{Breaker: &breaker.Config{Window: xtime.Duration(time.Second),
			Sleep:   xtime.Duration(time.Millisecond * 100),
			Bucket:  10,
			Ratio:   0.5,
			Request: 100}, Timeout: xtime.Duration(time.Second), App: &mars.App{Key: "0c4b8fe3ff35a4b6", Secret: "b370880d1aca7d3a289b9b9a7f4d6812"}}, Mars: &dc.HTTPServers{Inner: &mars.ServerConfig{Address: "127.0.0.1:7171"}}, Nodes: []string{"127.0.0.1:7171"}}, "api.flywk.com")
		node.client.SetTransport(gock.DefaultTransport)
		httpMock("POST", "http://api.flywk.com/discovery/register").Reply(200).JSON(`{"ts":1514341945,"code":409,"data":{"region":"shsb","zone":"fuck","appid":"main.arch.account-service","env":"pre","hostname":"cs4sq","http":"","rpc":"0.0.0.0:18888","weight":2}}`)
		i := model.NewInstance(reg)
		err := node.call(context.TODO(), model.Register, i, "http://api.flywk.com/discovery/register", &res)
		So(err, ShouldResemble, ecode.Conflict)
		So(res.Appid, ShouldResemble, "main.arch.account-service")
	})
}

func TestNodeCancel(t *testing.T) {
	Convey("test node renew 409 error", t, func() {
		i := model.NewInstance(reg)
		node := newNode(&dc.Config{HTTPClient: &mars.ClientConfig{Breaker: &breaker.Config{Window: xtime.Duration(time.Second),
			Sleep:   xtime.Duration(time.Millisecond * 100),
			Bucket:  10,
			Ratio:   0.5,
			Request: 100}, Timeout: xtime.Duration(time.Second), App: &mars.App{Key: "0c4b8fe3ff35a4b6", Secret: "b370880d1aca7d3a289b9b9a7f4d6812"}}, Mars: &dc.HTTPServers{Inner: &mars.ServerConfig{Address: "127.0.0.1:7171"}}, Nodes: []string{"127.0.0.1:7171"}}, "api.flywk.com")
		node.pRegisterURL = "http://127.0.0.1:7171/discovery/register"
		node.client.SetTransport(gock.DefaultTransport)
		httpMock("POST", "http://api.flywk.com/discovery/cancel").Reply(200).JSON(`{"code":0}`)
		err := node.Cancel(context.TODO(), i)
		So(err, ShouldBeNil)
	})
}

func TestNodeRenew(t *testing.T) {
	Convey("test node renew 409 error", t, func() {
		i := model.NewInstance(reg)
		node := newNode(&dc.Config{HTTPClient: &mars.ClientConfig{Breaker: &breaker.Config{Window: xtime.Duration(time.Second),
			Sleep:   xtime.Duration(time.Millisecond * 100),
			Bucket:  10,
			Ratio:   0.5,
			Request: 100}, Timeout: xtime.Duration(time.Second), App: &mars.App{Key: "0c4b8fe3ff35a4b6", Secret: "b370880d1aca7d3a289b9b9a7f4d6812"}}, Mars: &dc.HTTPServers{Inner: &mars.ServerConfig{Address: "127.0.0.1:7171"}}, Nodes: []string{"127.0.0.1:7171"}}, "api.flywk.com")
		node.pRegisterURL = "http://127.0.0.1:7171/discovery/register"
		node.client.SetTransport(gock.DefaultTransport)
		httpMock("POST", "http://api.flywk.com/discovery/renew").Reply(200).JSON(`{"code":409,"data":{"region":"shsb","zone":"fuck","appid":"main.arch.account-service","env":"pre","hostname":"cs4sq","http":"","rpc":"0.0.0.0:18888","weight":2}}`)
		httpMock("POST", "http://127.0.0.1:7171/discovery/register").Reply(200).JSON(`{"code":0}`)
		err := node.Renew(context.TODO(), i)
		So(err, ShouldBeNil)
	})
}

func TestNodeRenew2(t *testing.T) {
	Convey("test node renew 404 error", t, func() {
		i := model.NewInstance(reg)
		node := newNode(&dc.Config{HTTPClient: &mars.ClientConfig{Breaker: &breaker.Config{Window: xtime.Duration(time.Second),
			Sleep:   xtime.Duration(time.Millisecond * 100),
			Bucket:  10,
			Ratio:   0.5,
			Request: 100}, Timeout: xtime.Duration(time.Second), App: &mars.App{Key: "0c4b8fe3ff35a4b6", Secret: "b370880d1aca7d3a289b9b9a7f4d6812"}}, Mars: &dc.HTTPServers{Inner: &mars.ServerConfig{Address: "127.0.0.1:7171"}}, Nodes: []string{"127.0.0.1:7171"}}, "api.flywk.com")
		node.client.SetTransport(gock.DefaultTransport)
		httpMock("POST", "http://api.flywk.com/discovery/renew").Reply(200).JSON(`{"code":404}`)
		httpMock("POST", "http://api.flywk.com/discovery/register").Reply(200).JSON(`{"code":0}`)
		err := node.Renew(context.TODO(), i)
		So(err, ShouldBeNil)
	})
}

func TestSet(t *testing.T) {
	Convey("test set", t, func() {
		node := newNode(&dc.Config{HTTPClient: &mars.ClientConfig{Breaker: &breaker.Config{Window: xtime.Duration(time.Second),
			Sleep:   xtime.Duration(time.Millisecond * 100),
			Bucket:  10,
			Ratio:   0.5,
			Request: 100}, Timeout: xtime.Duration(time.Second), App: &mars.App{Key: "0c4b8fe3ff35a4b6", Secret: "b370880d1aca7d3a289b9b9a7f4d6812"}}, Mars: &dc.HTTPServers{Inner: &mars.ServerConfig{Address: "127.0.0.1:7171"}}, Nodes: []string{"127.0.0.1:7171"}}, "api.flywk.com")
		node.client.SetTransport(gock.DefaultTransport)
		httpMock("POST", "http://api.flywk.com/discovery/set").Reply(200).JSON(`{"ts":1514341945,"code":0}`)
		set := &model.ArgSet{
			Region:   "shsb",
			Env:      "pre",
			Appid:    "main.arch.account-service",
			Hostname: []string{"test1"},
			Status:   []int64{1},
		}
		err := node.Set(context.TODO(), set)
		So(err, ShouldBeNil)
	})
}

func httpMock(method, url string) *gock.Request {
	r := gock.New(url)
	r.Method = strings.ToUpper(method)
	return r
}
