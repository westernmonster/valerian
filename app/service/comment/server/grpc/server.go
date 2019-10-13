package server

import (
	"context"
	"fmt"
	"valerian/app/service/comment/api"
	"valerian/app/service/comment/service"
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
	comment, err := s.svr.GetComment(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	stat, err := s.svr.GetCommentStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	m, err := s.svr.GetAccountBaseInfo(ctx, comment.CreatedBy)
	if err != nil {
		return nil, err
	}

	resp := &api.CommentInfo{
		ID:         comment.ID,
		Content:    comment.Content,
		Deleted:    bool(comment.Deleted),
		Featured:   bool(comment.Featured),
		OwnerID:    comment.OwnerID,
		ResourceID: comment.ResourceID,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
		Stat: &api.CommentStat{
			ChildrenCount: int32(stat.ChildrenCount),
			LikeCount:     int32(stat.LikeCount),
			DislikeCount:  int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:       m.ID,
			UserName: m.UserName,
			Avatar:   m.Avatar,
		},
	}

	if m.Introduction != nil {
		resp.Creator.Introduction = &api.Creator_IntroductionValue{m.GetIntroductionValue()}
	}

	return resp, nil
}
