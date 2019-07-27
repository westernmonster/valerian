package rpc

import (
	"valerian/app/infra/config/conf"
	"valerian/app/infra/config/model"
	"valerian/app/infra/config/service"

	"valerian/library/net/rpc"
	"valerian/library/net/rpc/context"
)

// RPC export rpc service
type RPC struct {
	srv *service.Service
}

// New new rpc server.
func New(c *conf.Config, srv *service.Service) (svr *rpc.Server) {
	r := &RPC{srv: srv}
	svr = rpc.NewServer(c.RPCServer)
	if err := svr.Register(r); err != nil {
		panic(err)
	}
	return
}

// Ping check connection success.
func (r *RPC) Ping(c context.Context, arg *struct{}, res *struct{}) (err error) {
	return
}

// PushV4 push new config change to config-service
func (r *RPC) Push(c context.Context, a *model.ArgConf, res *struct{}) (err error) {
	service := &model.Service{Name: a.App, BuildVersion: a.BuildVer, Version: a.Ver}
	err = r.srv.Push(c, service)
	return
}

//SetTokenV4 update Token
func (r *RPC) SetToken(c context.Context, a *model.ArgToken, res *struct{}) (err error) {
	r.srv.SetToken(c, a.App, a.Token)
	return
}

//Hosts get host list.
func (r *RPC) Hosts(c context.Context, svr string, res *[]*model.Host) (err error) {
	*res, err = r.srv.Hosts(c, svr)
	return
}

//ClearHost clear host.
func (r *RPC) ClearHost(c context.Context, svr string, res *struct{}) error {
	return r.srv.ClearHost(c, svr)
}

// Force push new host config change to config-service
func (r *RPC) Force(c context.Context, a *model.ArgConf, res *struct{}) (err error) {
	service := &model.Service{Name: a.App, BuildVersion: a.BuildVer, Version: a.Ver}
	err = r.srv.Force(c, service, a.Hosts, a.SType)
	return
}
