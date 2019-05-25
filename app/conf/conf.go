package conf

import (
	"flag"
	"valerian/library/log"
	"valerian/library/net/http/mars"

	"github.com/BurntSushi/toml"
)

// global var
var (
	confPath string
	// client   *conf.Client
	// Conf config
	Conf = &Config{}
)

type Config struct {
	Log *log.Config

	// HTTPServer
	HTTPServer *mars.ServerConfig
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
