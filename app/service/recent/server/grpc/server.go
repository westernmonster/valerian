package server

import (
	"context"
	"fmt"
	"valerian/app/service/recent/api"
	"valerian/app/service/recent/service"
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
	api.RegisterRecentServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetRecentPubsPaged(c context.Context, req *api.RecentPubsReq) (*api.RecentPubsResp, error) {
	items, err := s.svr.GetUserRecentPubsPaged(c, req.AccountID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	resp := &api.RecentPubsResp{
		Items: make([]*api.RecentPubInfo, len(items)),
	}

	for i, v := range items {
		info := &api.RecentPubInfo{
			ID:         v.ID,
			AccountID:  v.AccountID,
			TargetID:   v.TargetID,
			TargetType: v.TargetType,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		resp.Items[i] = info

	}

	return resp, nil
}

func (s *server) GetRecentViewsPaged(c context.Context, req *api.RecentViewsReq) (*api.RecentViewsResp, error) {
	items, err := s.svr.GetUserRecentViewsPaged(c, req.AccountID, req.TargetType, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	resp := &api.RecentViewsResp{
		Items: make([]*api.RecentViewInfo, len(items)),
	}

	for i, v := range items {
		info := &api.RecentViewInfo{
			ID:         v.ID,
			AccountID:  v.AccountID,
			TargetID:   v.TargetID,
			TargetType: v.TargetType,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		resp.Items[i] = info

	}

	return resp, nil
}
