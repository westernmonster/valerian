package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"valerian/app/infra/discovery/conf"
	"valerian/app/infra/discovery/http"
	"valerian/app/infra/discovery/service"
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

	// init log
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("discovery start")

	tracing.Init(nil)
	// service init
	rand.Seed(time.Now().UnixNano())
	svc, cancel := service.New(conf.Conf)
	http.Init(conf.Conf, svc)
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("discovery get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cancel()
			time.Sleep(time.Second)
			log.Info("discovery exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
