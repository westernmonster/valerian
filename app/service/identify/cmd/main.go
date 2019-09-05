package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"valerian/app/service/identify/conf"
	"valerian/app/service/identify/server/grpc"
	"valerian/app/service/identify/server/http"
	"valerian/app/service/identify/service"
	"valerian/library/conf/env"
	ecode "valerian/library/ecode/tip"
	"valerian/library/log"
	"valerian/library/naming"
	"valerian/library/naming/discovery"
	xip "valerian/library/net/ip"
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

	// service init
	srv := service.New(conf.Conf)
	http.Init(conf.Conf, srv)

	ws := grpc.New(conf.Conf.WardenServer, srv)

	var (
		err    error
		cancel context.CancelFunc
	)
	{
		ip := xip.InternalIP()
		hn, _ := os.Hostname()
		dis := discovery.New(conf.Conf.Discovery)
		ins := &naming.Instance{
			Zone:     env.Zone,
			Env:      env.DeployEnv,
			AppID:    env.AppID,
			Hostname: hn,
			Addrs: []string{
				"http://" + ip + ":" + env.HTTPPort,
				"grpc://" + ip + ":" + env.GORPCPort,
			},
		}

		if cancel, err = dis.Register(context.Background(), ins); err != nil {
			panic(err)
		}
	}

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("identify-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if cancel != nil {
				cancel()
			}

			ws.Shutdown(context.Background())
			time.Sleep(time.Second * 2)
			srv.Close()
			log.Info("identify-service exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
