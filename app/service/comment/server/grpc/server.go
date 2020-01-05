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
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}

	return s.GetCommentInfo(ctx, req)
}

func (s *server) GetCommentsPaged(ctx context.Context, req *api.CommentListReq) (*api.CommentListResp, error) {
	return nil, nil
}

func (s *server) AddComment(ctx context.Context, req *api.AddCommentReq) (*api.EmptyStruct, error) {
	return nil, nil
}

func (s *server) DeleteComment(ctx context.Context, req *api.DeleteReq) (*api.EmptyStruct, error) {
	return nil, nil
}
