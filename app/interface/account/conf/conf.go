package conf

import (
	"flag"
	"time"
	"valerian/library/cache/memcache"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

type Config struct {
}

// DB db config.
type DB struct {
}

// Memcache memcache config.
type Memcache struct {
	Auth *MC
}

// MC .
type MC struct {
	*memcache.Config
	Expire time.Duration
}

// DC data center.
type DC struct {
	Num  int
	Desc string
}
