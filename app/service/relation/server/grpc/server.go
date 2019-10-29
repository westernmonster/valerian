package server

import (
	"context"
	"fmt"

	"valerian/app/service/relation/api"
	"valerian/app/service/relation/service"
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
	api.RegisterRelationServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetFollowingPaged(ctx context.Context, req *api.RelationReq) (*api.FollowingResp, error) {
	resp, err := s.svr.FollowingPaged(ctx, req.AccountID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	return api.FromFollowingResp(resp), nil
}

func (s *server) GetFansPaged(ctx context.Context, req *api.RelationReq) (*api.FansResp, error) {
	resp, err := s.svr.FansPaged(ctx, req.AccountID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	return api.FromFansResp(resp), nil
}

func (s *server) Follow(ctx context.Context, req *api.FollowReq) (*api.EmptyStruct, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	err := s.svr.Follow(ctx, req.AccountID, req.TargetAccountID)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) Unfollow(ctx context.Context, req *api.FollowReq) (*api.EmptyStruct, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	err := s.svr.Unfollow(ctx, req.AccountID, req.TargetAccountID)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) IsFollowing(ctx context.Context, req *api.FollowReq) (*api.IsFollowingResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	isFollowing, err := s.svr.IsFollowing(ctx, req.AccountID, req.TargetAccountID)
	if err != nil {
		return nil, err
	}

	resp := &api.IsFollowingResp{
		IsFollowing: isFollowing,
	}
	return resp, nil

}

func (s *server) GetFansIDs(ctx context.Context, req *api.AidReq) (*api.IDsResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	ids, err := s.svr.GetFansIDs(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}
