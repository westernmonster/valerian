package conf

import (
	"errors"
	"flag"

	"valerian/library/cache/memcache"
	"valerian/library/conf"
	"valerian/library/database/sqalx"
	ecode "valerian/library/ecode/tip"
	"valerian/library/log"
	"valerian/library/naming/discovery"
	"valerian/library/net/http/mars"
	xtime "valerian/library/time"
	"valerian/library/tracing"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf     = &Config{}
	client   *conf.Client
)

type Config struct {
	Log       *log.Config
	Tracer    *tracing.Config
	Ecode     *ecode.Config
	DB        *DB
	Memcache  *Memcache
	Mars      *mars.ServerConfig
	Discovery *discovery.Config
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
	if confPath != "" {
		return local()
	}
	return remote()
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
			log.Info("config reload")
			if load() != nil {
				log.Errorf("config reload error (%v)", err)
			}
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
