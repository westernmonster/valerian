package server

import (
	"context"
	"fmt"

	"valerian/app/service/like/api"
	"valerian/app/service/like/service"
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
	api.RegisterLikeServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) IsLike(ctx context.Context, req *api.LikeReq) (*api.LikeInfo, error) {
	isLike, err := s.svr.IsLike(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.LikeInfo{IsLike: isLike}, nil
}

func (s *server) Like(ctx context.Context, req *api.LikeReq) (*api.EmptyStruct, error) {
	err := s.svr.Like(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) CancelLike(ctx context.Context, req *api.LikeReq) (*api.EmptyStruct, error) {
	err := s.svr.CancelLike(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) IsDislike(ctx context.Context, req *api.DislikeReq) (*api.DislikeInfo, error) {
	isDislike, err := s.svr.IsDislike(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.DislikeInfo{IsDislike: isDislike}, nil
}

func (s *server) Dislike(ctx context.Context, req *api.DislikeReq) (*api.EmptyStruct, error) {
	err := s.svr.Dislike(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) CancelDislike(ctx context.Context, req *api.DislikeReq) (*api.EmptyStruct, error) {
	err := s.svr.CancelDislike(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}
