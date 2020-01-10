package main

import (
	"os"
	"os/signal"
	"syscall"

	flag "github.com/spf13/pflag"

	"valerian/app/interface/like/conf"
	"valerian/app/interface/like/http"
	"valerian/app/interface/like/service"
	ecode "valerian/library/ecode/tip"
	"valerian/library/log"
	"valerian/library/tracing"
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

	tracing.Init(nil)

	// interface init
	srv := service.New(conf.Conf)
	http.Init(conf.Conf, srv)

	// ws := server.New(conf.Conf.WardenServer, srv)

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("like-interface get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			srv.Close()
			log.Info("like-interface exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
