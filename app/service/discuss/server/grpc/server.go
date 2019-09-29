package server

import (
	"context"
	"fmt"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/service"
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
	api.RegisterDiscussionServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetDiscussionInfo(ctx context.Context, req *api.IDReq) (*api.DiscussionInfo, error) {
	resp, err := s.svr.GetDiscussion(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	stat, err := s.svr.GetDiscussionStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return api.FromDiscussion(resp, stat), nil
}

func (s *server) GetDiscussionCategories(ctx context.Context, req *api.CategoriesReq) (*api.CategoriesResp, error) {
	data, err := s.svr.GetDiscussCategories(ctx, req.TopicID)
	if err != nil {
		return nil, err
	}

	resp := &api.CategoriesResp{
		Items: make([]*api.CategoryInfo, len(data)),
	}

	for i, v := range data {
		item := api.FromCategory(v)
		resp.Items[i] = item
	}

	return resp, nil
}
