package main

import (
	"os"
	"os/signal"
	"syscall"

	flag "github.com/spf13/pflag"

	"valerian/app/interface/locale/conf"
	"valerian/app/interface/locale/http"
	"valerian/app/interface/locale/service"
	ecode "valerian/library/ecode/tip"
	"valerian/library/log"
	"valerian/library/tracing"
)

var svc *service.Service

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Errorf("conf.Init() error(%v)", err)
		panic(err)
	}

	// init ecode
	ecode.Init(nil)
	// init log
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("app-locale start")
	// init trace
	tracing.Init(nil)
	// service init
	svc = service.New(conf.Conf)
	http.Init(conf.Conf, svc)

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("app-locale get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			log.Info("app-locale exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
