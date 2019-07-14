package http

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/rate"
	"valerian/library/rate/limit"
	"valerian/library/rate/limit/bench/stress/conf"
	"valerian/library/rate/limit/bench/stress/service"
	"valerian/library/rate/vegas"
)

var (
	svc *service.Service
	req int64
	qps int64
)

// Init init
func Init(c *conf.Config) {
	rand.Seed(time.Now().Unix())
	initService(c)
	// init router
	engineInner := mars.DefaultServer(c.Mars.Inner)
	outerRouter(engineInner)
	if err := engineInner.Start(); err != nil {
		log.Error(fmt.Sprintf("xhttp.Serve error(%v)", err))
		panic(err)
	}
	engineLocal := mars.DefaultServer(c.Mars.Local)
	localRouter(engineLocal)
	if err := engineLocal.Start(); err != nil {
		log.Error(fmt.Sprintf("xhttp.Serve error(%v)", err))
		panic(err)
	}
	// if log.V(1) {
	go calcuQPS()
	// }
}

// initService init services.
func initService(c *conf.Config) {
	//	idfSvc = identify.New(c.Identify)
	svc = service.New(c)
}

// outerRouter init outer router api path.
func outerRouter(e *mars.Engine) {
	v := vegas.New()
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		defer ticker.Stop()
		for {
			<-ticker.C
			m := v.Stat()
			log.Info(fmt.Sprintf("vegas: limit(%d) inFlight(%d) minRtt(%v) rtt(%v)", m.Limit, m.InFlight, m.MinRTT, m.LastRTT))
		}
	}()
	l := limit.New(nil)
	//init api
	e.GET("/monitor/ping", ping)
	group := e.Group("/stress")
	group.GET("/normal", aqmTest)
	group.GET("/vegas", func(c *mars.Context) {
		start := time.Now()
		done, success := v.Acquire()
		if !success {
			done(time.Time{}, rate.Ignore)
			c.AbortWithStatus(509)
			return
		}
		defer done(start, rate.Success)
		c.Next()
	}, aqmTest)
	group.GET("/attack", func(c *mars.Context) {
		done, err := l.Allow(c)
		defer done(rate.Success)
		if err != nil {
			c.AbortWithStatus(509)
			return
		}
		c.Next()
	}, aqmTest)

}

func calcuQPS() {
	var creq, breq int64
	for {
		time.Sleep(time.Second * 5)
		creq = atomic.LoadInt64(&req)
		delta := creq - breq
		atomic.StoreInt64(&qps, delta/5)
		breq = creq
		log.Info(fmt.Sprintf("HTTP QPS:%d", atomic.LoadInt64(&qps)))
	}

}
func aqmTest(c *mars.Context) {
	params := c.Request.Form
	sleep, err := strconv.ParseInt(params.Get("sleep"), 10, 64)
	if err == nil {
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
	atomic.AddInt64(&req, 1)
	for i := 0; i < 3000+rand.Intn(3000); i++ {
		crc32.Checksum([]byte(`testasdwfwfsddsfgwddcscsc
			http://git.flywk.com/platform/go-common/merge_requests/new?merge_request%5Bsource_branch%5D=stress%2Fcodel`), crc32.IEEETable)
	}
}

// ping check server ok.
func ping(c *mars.Context) {
}

// innerRouter init local router api path.
func localRouter(e *mars.Engine) {
	//init api
	e.GET("/monitor/ping", ping)
	group := e.Group("/x/main/stress")
	{
		group.GET("", aqmTest)
	}
}
