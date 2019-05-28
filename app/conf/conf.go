package conf

import (
	"flag"

	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	xtime "valerian/library/time"
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
	Memcache   *Memcache
	Aliyun     *Aliyun
}

// DB db config.
type DB struct {
	Main *sqalx.Config
	Auth *sqalx.Config
}

// Memcache memcache config.
type Memcache struct {
	Auth *MC
	Main *MC
}

// MC .
type MC struct {
	*memcache.Config
	Expire xtime.Duration
}

type Aliyun struct {
	AccessKeyID     string
	AccessKeySecret string
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
func Init() {
	if confPath != "" {
		err := local()
		if err != nil {
			panic(err)
		}
		return
	}

	panic("load config file failed")
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)

	return
}
