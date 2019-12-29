package server

import (
	"context"
	"fmt"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/service"
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

func (s *server) GetDiscussionStat(ctx context.Context, req *api.IDReq) (*api.DiscussionStat, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	stat, err := s.svr.GetDiscussionStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	resp := &api.DiscussionStat{
		LikeCount:    int32(stat.LikeCount),
		DislikeCount: int32(stat.DislikeCount),
		CommentCount: int32(stat.CommentCount),
	}
	return resp, err
}

func (s *server) GetDiscussionInfo(ctx context.Context, req *api.IDReq) (*api.DiscussionInfo, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}

	return s.svr.GetDiscussionInfo(ctx, req)
}

func (s *server) GetDiscussionCategories(ctx context.Context, req *api.CategoriesReq) (*api.CategoriesResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
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

func (s *server) GetUserDiscussionIDsPaged(ctx context.Context, req *api.UserDiscussionsReq) (*api.IDsResp, error) {
	var ids []int64
	ids, err := s.svr.GetUserDiscussionIDsPaged(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}

func (s *server) GetUserDiscussionsPaged(ctx context.Context, req *api.UserDiscussionsReq) (*api.DiscussionsResp, error) {
	items, err := s.svr.GetUserDiscussionsPaged(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.DiscussionsResp{
		Items: items,
	}

	return resp, nil
}

func (s *server) GetTopicDiscussionsPaged(ctx context.Context, req *api.TopicDiscussionsReq) (*api.DiscussionsResp, error) {
	items, err := s.svr.GetTopicDiscussionsPaged(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.DiscussionsResp{
		Items: items,
	}

	return resp, nil
}

func (s *server) AddDiscussion(ctx context.Context, req *api.ArgAddDiscussion) (*api.IDResp, error) {
	id, err := s.svr.AddDiscussion(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.IDResp{
		ID: id,
	}

	return resp, nil
}

func (s *server) UpdateDiscussion(ctx context.Context, req *api.ArgUpdateDiscussion) (*api.EmptyStruct, error) {
	err := s.svr.UpdateDiscussion(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) DelDiscussion(ctx context.Context, req *api.IDReq) (*api.EmptyStruct, error) {
	err := s.svr.DelDiscussion(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) SaveDiscussionCategories(ctx context.Context, req *api.ArgSaveDiscussCategories) (*api.EmptyStruct, error) {
	err := s.svr.SaveDiscussCategories(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) SaveDiscussionFiles(ctx context.Context, req *api.ArgSaveDiscussionFiles) (*api.EmptyStruct, error) {
	err := s.svr.SaveDiscussionFiles(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) GetDiscussionFiles(ctx context.Context, req *api.IDReq) (*api.DiscussionFilesResp, error) {
	return s.svr.GetDiscussionFiles(ctx, req)
}
