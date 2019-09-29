package server

import (
	"context"
	"fmt"

	"valerian/app/service/account/api"
	"valerian/app/service/account/service"
	"valerian/library/log"
	"valerian/library/net/metadata"
	"valerian/library/net/rpc/warden"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// New Identify warden rpc server
func New(cfg *warden.ServerConfig, s *service.Service) *warden.Server {
	w := warden.NewServer(cfg)
	w.Use(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if resp, err = handler(ctx, req); err != nil {
			log.For(ctx).Info("rpc call",
				zap.String("path", info.FullMethod),
				zap.String("caller", metadata.String(ctx, metadata.Caller)),
				zap.String("args", fmt.Sprintf("%v", req)),
				zap.String("error", fmt.Sprintf("%+v", err)))
		}
		return
	})
	api.RegisterAccountServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) BasicInfo(ctx context.Context, req *api.AidReq) (*api.BaseInfoReply, error) {
	resp, err := s.svr.BaseInfo(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	return api.FromBaseInfo(resp), nil
}

func (s *server) BaseInfos(ctx context.Context, req *api.AidsReq) (*api.BaseInfosReply, error) {
	resp, err := s.svr.BatchBaseInfo(ctx, req.Aids)
	if err != nil {
		return nil, err
	}

	baseInfos := make(map[int64]*api.BaseInfoReply, len(resp))
	baseInfosReply := &api.BaseInfosReply{
		BaseInfos: baseInfos,
	}

	for k, v := range resp {
		baseInfos[k] = api.FromBaseInfo(v)
	}

	return baseInfosReply, nil
}

func (s *server) AccountStat(ctx context.Context, req *api.AidReq) (*api.AccountStatInfo, error) {
	resp, err := s.svr.GetAccountStat(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	return api.FromStat(resp), nil
}
