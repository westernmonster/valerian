package client

import (
	"context"
	"valerian/app/infra/config/model"
	"valerian/library/net/rpc"
)

const (
	_appid     = "config.service"
	_push      = "RPC.Push"
	_force     = "RPC.Force"
	_hosts     = "RPC.Hosts"
	_setToken  = "RPC.SetToken"
	_clearHost = "RPC.ClearHost"
)

var (
	_noArg = &struct{}{}
)

//Service service.
type Service struct {
	client *rpc.Client2
}

// New2 new a config service.
func New(c *rpc.ClientConfig) (s *Service) {
	s = &Service{}
	s.client = rpc.NewDiscoveryCli(_appid, c)
	return
}

// Push push new ver to config-service
func (s *Service) Push(c context.Context, arg *model.ArgConf) (err error) {
	err = s.client.Boardcast(c, _push, arg, _noArg)
	return
}

// SetTokenV4 update token in config-service
func (s *Service) SetToken(c context.Context, arg *model.ArgToken) (err error) {
	err = s.client.Boardcast(c, _setToken, arg, _noArg)
	return
}

//Hosts get host list.
func (s *Service) Hosts(c context.Context, svr string) (hosts []*model.Host, err error) {
	err = s.client.Call(c, _hosts, svr, &hosts)
	return
}

// ClearHost update token in config-service
func (s *Service) ClearHost(c context.Context, svr string) (err error) {
	err = s.client.Call(c, _clearHost, svr, _noArg)
	return
}

// Force push new host ver to config-service
func (s *Service) Force(c context.Context, arg *model.ArgConf) (err error) {
	err = s.client.Boardcast(c, _force, arg, _noArg)
	return
}
