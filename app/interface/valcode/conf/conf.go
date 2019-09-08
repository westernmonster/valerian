package conf

import (
	"errors"

	flag "github.com/spf13/pflag"

	"valerian/library/cache/memcache"
	"valerian/library/conf"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/naming/discovery"
	"valerian/library/net/http/mars"
	xtime "valerian/library/time"
	"valerian/library/tracing"

	"github.com/BurntSushi/toml"
)

// Conf global variable.
var (
	Conf     = &Config{}
	client   *conf.Client
	confPath string
)

type Config struct {
	Log      *log.Config
	Mars     *mars.ServerConfig
	Tracer   *tracing.Config
	Memcache *Memcache

	Aliyun    *Aliyun
	Discovery *discovery.Config
}

type Aliyun struct {
	AccessKeyID     string
	AccessKeySecret string
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

// DB db config.
type DB struct {
	Auth *sqalx.Config
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func remote() (err error) {
	if client, err = conf.New(); err != nil {
		return
	}
	if err = load(); err != nil {
		return
	}
	go func() {
		for range client.Event() {
			log.Info("config event")
		}
	}()
	return
}

func load() (err error) {
	var (
		s       string
		ok      bool
		tmpConf *Config
	)
	if s, ok = client.Toml2(); !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}

func init() {
	flag.StringVar(&confPath, "config", "", "default config path")
}

// Init int config
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
}
