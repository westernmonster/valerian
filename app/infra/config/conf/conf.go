package conf

import (
	"valerian/library/cache/redis"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/naming/discovery"
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/antispam"
	"valerian/library/net/rpc"
	xtime "valerian/library/time"

	flag "github.com/spf13/pflag"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf init config
	Conf *Config
)

// Config config.
type Config struct {
	// log
	Log *log.Config
	//rpc server2
	RPCServer *rpc.ServerConfig
	// db
	DB *sqalx.Config
	// redis
	Redis *redis.Config
	// timeout
	PollTimeout xtime.Duration
	// local cache
	PathCache string
	//BM
	Mars *mars.ServerConfig
	// Antispam
	Antispam *antispam.Config

	Discovery *discovery.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "./config-service-example.toml", "config path")
}

// Init init.
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
