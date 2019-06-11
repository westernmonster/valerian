package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"valerian/app/conf"
	"valerian/library/conf/env"
	"valerian/library/locale"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/tracing"

	httpAccount "valerian/app/interface/account/http"
	authMiddleware "valerian/app/interface/auth"
	httpCertification "valerian/app/interface/certification/http"
	httpFile "valerian/app/interface/file/http"
	httpLocale "valerian/app/interface/locale/http"
	httpLocation "valerian/app/interface/location/http"
	httpAuth "valerian/app/interface/passport-auth/http"
	httpLogin "valerian/app/interface/passport-login/http"
	httpRegister "valerian/app/interface/passport-register/http"
	httpTopic "valerian/app/interface/topic/http"
	httpValcode "valerian/app/interface/valcode/http"

	"github.com/joho/godotenv"
)

// @title 飞行百科 API
// @version 1.0
// @description 飞行百科 API
// @description 所有返回结果以如下JSON对象返回
// @description <pre>
// @description {
// @description    "code": 0,           // 如果没有另行约定，一般为200
// @description    "message": "string", // 一般用户发生错误或者验证失败时候返回的消息
// @description    "result": {},        // 返回结果，所有文档所示的返回结果处于这个字段
// @description    "success": true      // 用于判断是否执行成功
// @description }
// @description </pre>
// @description
// @description
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
// @host dev.flywk.com
// @BasePath /api/v1
func main() {
	flag.Parse()

	// Load Environment Variables
	godotenv.Load()
	env.Init()

	// Load locale files
	locale.LoadTranslateFile()

	conf.Init()
	log.Init(nil)
	tracing.Init(nil)

	initHTTP(conf.Conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info(fmt.Sprintf("web-interface get a signal %s", s.String()))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("web-interface exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}

func initHTTP(c *conf.Config) {
	engine := mars.DefaultServer(c.HTTPServer)
	authMiddleware.Init(c)

	httpLogin.Init(c, engine)
	httpValcode.Init(c, engine)
	httpLocation.Init(c, engine)
	httpRegister.Init(c, engine)
	httpAuth.Init(c, engine)
	httpAccount.Init(c, engine)
	httpFile.Init(c, engine)
	httpTopic.Init(c, engine)
	httpLocale.Init(c, engine)
	httpCertification.Init(c, engine)

	if err := engine.Start(); err != nil {
		log.Error(fmt.Sprintf("engine.Start error(%v)", err))
		panic(err)
	}
}
