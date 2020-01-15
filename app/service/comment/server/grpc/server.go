package server

import (
	"context"
	"fmt"
	"valerian/app/service/comment/api"
	"valerian/app/service/comment/service"
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
	api.RegisterCommentServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetCommentInfo(ctx context.Context, req *api.IDReq) (*api.CommentInfo, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}

	return s.svr.GetCommentInfo(ctx, req.Aid, req.ID)
}

func (s *server) GetCommentsPaged(ctx context.Context, req *api.CommentListReq) (*api.CommentListResp, error) {
	return s.svr.GetCommentsPaged(ctx, req)
}

func (s *server) AddComment(ctx context.Context, req *api.AddCommentReq) (*api.IDResp, error) {
	if id, err := s.svr.AddComment(ctx, req); err != nil {
		return nil, err
	} else {
		return &api.IDResp{ID: id}, nil
	}
}

func (s *server) DeleteComment(ctx context.Context, req *api.DeleteReq) (*api.EmptyStruct, error) {
	if err := s.svr.DelComment(ctx, req.Aid, req.ID); err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) GetAllChildrenComment(ctx context.Context, req *api.IDReq) (*api.ChildrenCommentListResp, error) {
	return s.svr.GetAllChildrenComments(ctx, req)
}
