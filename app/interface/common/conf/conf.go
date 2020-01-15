package conf

import (
	"errors"

	flag "github.com/spf13/pflag"

	"valerian/library/cache/memcache"
	"valerian/library/conf"
	"valerian/library/database/sqalx"
	ecode "valerian/library/ecode/tip"
	"valerian/library/log"
	"valerian/library/mq"
	"valerian/library/naming/discovery"
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/middleware/auth"
	"valerian/library/net/rpc/warden"
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
	Log    *log.Config
	Mars   *mars.ServerConfig
	Tracer *tracing.Config
	DB     *DB
	Ecode  *ecode.Config
	// Auth
	Auth      *auth.Config
	Memcache  *Memcache
	Discovery *discovery.Config
	Nats      *mq.Config

	Aliyun *Aliyun
	OSS    *OSS

	LikeRPC     *warden.ClientConfig
	AccountRPC  *warden.ClientConfig
	DiscussRPC  *warden.ClientConfig
	TopicRPC    *warden.ClientConfig
	ArticleRPC  *warden.ClientConfig
	CommentRPC  *warden.ClientConfig
	RelationRPC *warden.ClientConfig
}

type Aliyun struct {
	AccessKeyID     string
	AccessKeySecret string
	RegionID        string
	RoleArn         string
	RoleSessionName string
}

type OSS struct {
	Host             string
	CallbackURL      string
	ImageDir         string
	FileDir          string
	CertificationDir string
	OtherDir         string
}

// DB db config.
type DB struct {
	Main *sqalx.Config
}

// Memcache memcache config.
type Memcache struct {
	Main *MC
}

// MC .
type MC struct {
	*memcache.Config
	Expire xtime.Duration
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
