package main

import (
	"flag"
	"valerian/app/conf"
	"valerian/library/log"
)

func main() {
	flag.Parse()
	log.Init(conf.Conf.Log)

	log.Info("web-interface start")

}
