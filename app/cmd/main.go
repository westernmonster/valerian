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
	httpFile "valerian/app/interface/file/http"
	httpLocation "valerian/app/interface/location/http"
	httpAuth "valerian/app/interface/passport-auth/http"
	httpLogin "valerian/app/interface/passport-login/http"
	httpRegister "valerian/app/interface/passport-register/http"
	httpTopic "valerian/app/interface/topic/http"
	httpValcode "valerian/app/interface/valcode/http"

	"github.com/joho/godotenv"
)

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

	if err := engine.Start(); err != nil {
		log.Error(fmt.Sprintf("engine.Start error(%v)", err))
		panic(err)
	}
}
