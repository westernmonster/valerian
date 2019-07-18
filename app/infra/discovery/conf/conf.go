package conf

import (
	"errors"
	"flag"
	"net"

	"valerian/library/conf"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	xip "valerian/library/net/ip"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf conf
	Conf = &Config{}
	// ConfCh for update node of server.
	ConfCh    = make(chan struct{}, 1)
	configKey = "discovery-service.toml"
)

// Config config
type Config struct {
	Nodes      []string
	Zones      map[string][]string // zone -> nodes
	Mars       *HTTPServers
	Log        *log.Config
	HTTPClient *mars.ClientConfig
}

func (c *Config) fix() (err error) {
	// check ip
	host, port, err := net.SplitHostPort(c.Mars.Inner.Address)
	if err != nil {
		return
	}
	if host == "0.0.0.0" || host == "127.0.0.1" || host == "" {
		host = xip.InternalIP()
	}
	c.Mars.Inner.Address = host + ":" + port
	return
}

// HTTPServers Http Servers
type HTTPServers struct {
	Inner *mars.ServerConfig
}

func init() {
	// flag.StringVar(&confPath, "conf", "discovery-example.toml", "config path")
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf
func Init() (err error) {
	if confPath != "" {
		if _, err = toml.DecodeFile(confPath, &Conf); err != nil {
			return
		}
		return Conf.fix()
	}
	err = remote()
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
				continue
			}
			// to change the node of server
			ConfCh <- struct{}{}
		}
	}()
	return
}

func load() (err error) {
	s, ok := client.Value(configKey)
	if !ok {
		return errors.New("load config center error")
	}
	var tmpConf *Config
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	if err = tmpConf.fix(); err != nil {
		return
	}
	// copy
	*Conf = *tmpConf
	return
}
