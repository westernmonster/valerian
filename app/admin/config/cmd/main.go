package main

import (
	"os"
	"os/signal"
	"syscall"

	flag "github.com/spf13/pflag"

	"valerian/app/admin/config/conf"
	"valerian/app/admin/config/http"
	_ "valerian/app/admin/config/pkg/lint/json"
	_ "valerian/app/admin/config/pkg/lint/toml"
	"valerian/app/admin/config/service"
	"valerian/library/log"
	"valerian/library/tracing"
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
	log.Info("config-admin start")
	tracing.Init(nil)

	svr := service.New(conf.Conf)
	http.Init(conf.Conf, svr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("config-admin get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svr.Close()
			log.Info("config-admin exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
