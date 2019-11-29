package server

import (
	"context"
	"fmt"

	"valerian/app/service/fav/api"
	"valerian/app/service/fav/service"
	"valerian/library/database/sqalx"
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
	api.RegisterFavServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) IsFav(ctx context.Context, req *api.FavReq) (*api.FavInfo, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	isFav, err := s.svr.IsFav(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.FavInfo{IsFav: isFav}, nil
}

func (s *server) Fav(ctx context.Context, req *api.FavReq) (*api.EmptyStruct, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	err := s.svr.Fav(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) Unfav(ctx context.Context, req *api.FavReq) (*api.EmptyStruct, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	err := s.svr.Unfav(ctx, req.AccountID, req.TargetID, req.TargetType)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) GetUserFavIDsPaged(ctx context.Context, req *api.UserFavsReq) (*api.IDsResp, error) {
	ids, err := s.svr.GetFavIDsPaged(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}
