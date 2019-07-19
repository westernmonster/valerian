package conf

import (
	"valerian/library/conf"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/tracing"
)

var (
	confPath string
	// Conf conf
	Conf      = &Config{}
	configKey = "msm-service.toml"
	client    *conf.Client
)

type Config struct {
	Log        *log.Config
	Mars       *mars.ServerConfig
	HTTPClient *mars.ClientConfig
	Tracer     *tracing.Config
	DB         *DB
	Tree       *ServiceTree
}

// DB db config.
type DB struct {
	Main *sqalx.Config
	Auth *sqalx.Config
}

// ServiceTree ServiceTree.
type ServiceTree struct {
	Host       string
	PlatformID string
}
