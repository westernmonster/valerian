package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	flag "github.com/spf13/pflag"

	"valerian/app/service/msm/conf"
	"valerian/app/service/msm/http"
	"valerian/app/service/msm/service"
	"valerian/library/conf/env"
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
	// init log
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("msm-service start")
	tracing.Init(nil)

	svr := service.New(conf.Conf)
	http.Init(conf.Conf, svr)

	// start discovery register
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
			},
		}

		if cancel, err = dis.Register(context.Background(), ins); err != nil {
			panic(err)
		}
	}
	// end discovery register

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("msm-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if cancel != nil {
				cancel()
			}
			svr.Close()
			log.Info("msm-service exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
