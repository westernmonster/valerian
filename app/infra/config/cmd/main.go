package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"valerian/app/infra/config/conf"
	"valerian/app/infra/config/http"
	rpc "valerian/app/infra/config/rpc/server"
	"valerian/app/infra/config/service"
	"valerian/library/conf/env"
	"valerian/library/log"
	"valerian/library/naming"
	"valerian/library/naming/discovery"
	xip "valerian/library/net/ip"
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
	// service init
	srv := service.New(conf.Conf)
	rpcSvr := rpc.New(conf.Conf, srv)
	http.Init(conf.Conf, srv, rpcSvr)
	// start discovery register
	var (
		err    error
		cancel context.CancelFunc
	)
	if env.IP == "" {
		ip := xip.InternalIP()
		hn, _ := os.Hostname()
		dis := discovery.New(nil)
		ins := &naming.Instance{
			Zone:     env.Zone,
			Env:      env.DeployEnv,
			AppID:    "config.service",
			Hostname: hn,
			Addrs: []string{
				"http://" + ip + ":" + env.HTTPPort,
				"gorpc://" + ip + ":" + env.GORPCPort,
			},
		}
		if cancel, err = dis.Register(context.Background(), ins); err != nil {
			panic(err)
		}
	}
	// end discovery register

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("config-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if cancel != nil {
				cancel()
			}
			rpcSvr.Close()
			srv.Close()
			log.Info("config-service exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
