package conf

import (
	"errors"
	"valerian/library/conf"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/naming/discovery"
	"valerian/library/net/http/mars"
	"valerian/library/net/rpc"
	"valerian/library/tracing"

	flag "github.com/spf13/pflag"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf conf
	Conf   = &Config{}
	client *conf.Client
)

type Config struct {
	Log        *log.Config
	Mars       *mars.ServerConfig
	HTTPClient *mars.ClientConfig
	Tracer     *tracing.Config
	DB         *DB
	Tree       *ServiceTree
	ConfSvr    *rpc.ClientConfig
	Discovery  *discovery.Config
}

// DB db config.
type DB struct {
	Main *sqalx.Config
	Auth *sqalx.Config
	Apm  *sqalx.Config
}

// ServiceTree ServiceTree.
type ServiceTree struct {
	Host       string
	PlatformID string
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
