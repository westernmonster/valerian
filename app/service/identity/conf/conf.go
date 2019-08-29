package conf

import (
	"errors"
	"flag"
	"valerian/library/cache/memcache"
	"valerian/library/conf"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/net/rpc/warden"
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
	Log        *log.Config
	Mars       *mars.ServerConfig
	HTTPClient *mars.ClientConfig
	Tracer     *tracing.Config

	// IdentifyConfig
	Identify *IdentifyConfig

	// grpc server
	WardenServer *warden.ServerConfig

	Memcache *memcache.Config
	// MemcacheLoginLog
	MemcacheLoginLog *memcache.Config
}

// IdentifyConfig identify config
type IdentifyConfig struct {
	AuthHost string
	// LoginLogConsumerSize goroutine size
	LoginLogConsumerSize int
	// LoginCacheExpires login check cache expires
	LoginCacheExpires int32
	// IntranetCIDR
	IntranetCIDR []string
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
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init int config
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
}
