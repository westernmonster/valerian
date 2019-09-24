package server

import (
	"context"
	"fmt"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/service"
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
				zap.String("args", fmt.Sprintf("%+v", err)))
		}
		return
	})
	api.RegisterTopicServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetTopicInfo(ctx context.Context, req *api.TopicReq) (*api.TopicInfo, error) {
	resp, err := s.svr.GetTopic(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return api.FromTopic(resp), nil
}

func (s *server) GetTopicMemberRole(ctx context.Context, req *api.TopicMemberRoleReq) (resp *api.MemberRoleReply, err error) {
	resp = &api.MemberRoleReply{}
	isMember, role, err := s.svr.GetTopicManagerRole(ctx, req.TopicID, req.AccountID)
	if err != nil {
		return nil, err
	}

	resp.IsMember = isMember
	resp.Role = role

	return resp, nil
}
