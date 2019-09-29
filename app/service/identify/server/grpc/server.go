package grpc

import (
	"context"
	"fmt"
	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/service"
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
			log.For(ctx).Info("",
				zap.String("path", info.FullMethod),
				zap.String("caller", metadata.String(ctx, metadata.Caller)),
				zap.String("args", fmt.Sprintf("%v", req)),
				zap.String("error", fmt.Sprintf("%+v", err)))
		}
		return
	})
	api.RegisterIdentifyServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetTokenInfo(ctx context.Context, req *api.TokenReq) (*api.AuthReply, error) {
	return s.svr.GetTokenInfo(ctx, req.Token)
}
