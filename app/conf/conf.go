package conf

import (
	"flag"
	"fmt"
	"time"

	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/tracing"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf config
	Conf = &Config{}
)

type Config struct {
	DC         *DC
	Log        *log.Config
	HTTPServer *mars.ServerConfig
	Tracer     *tracing.Config
	DB         *DB
}

// DB db config.
type DB struct {
	Main *sqalx.Config
	Auth *sqalx.Config
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

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init conf
func Init() error {
	fmt.Println(confPath)
	if confPath != "" {
		return local()
	}

	return nil
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)

	return
}
