package main

import (
	"os"
	"os/signal"
	"syscall"
	"valerian/app/interface/locale/conf"
	"valerian/app/interface/locale/http"
	ecode "valerian/library/ecode/tip"
	"valerian/library/log"
	"valerian/library/tracing"

	flag "github.com/spf13/pflag"
)

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
	log.Info("app-tag start")
	// init trace
	tracing.Init(nil)
	// service init
	http.Init(conf.Conf)
	// init pprof conf.Conf.Perf
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("app-tag get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("app-tag exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
