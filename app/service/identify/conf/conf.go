package conf

import (
	"errors"
	"flag"
	"valerian/library/cache/memcache"
	"valerian/library/conf"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/naming/discovery"
	"valerian/library/net/http/mars"
	"valerian/library/net/rpc/warden"
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
	Log    *log.Config
	Mars   *mars.ServerConfig
	Tracer *tracing.Config
	DB     *DB
	AuthMC *MC

	// grpc server
	WardenServer *warden.ServerConfig

	Discovery *discovery.Config
}

// MC .
type MC struct {
	*memcache.Config
	Expire xtime.Duration
}

// DB db config.
type DB struct {
	Main *sqalx.Config
	Auth *sqalx.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config file")
}

// Init init.
func Init() (err error) {
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
			if err := load(); err != nil {
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
