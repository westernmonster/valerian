package conf

import (
	"time"
	"valerian/library/cache/redis"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/antispam"
	"valerian/library/net/rpc"
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
	PollTimeout time.Duration
	// local cache
	PathCache string
	//BM
	Mars mars.ServerConfig
	// Antispam
	Antispam *antispam.Config
}
