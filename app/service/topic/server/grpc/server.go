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
				zap.String("error", fmt.Sprintf("%+v", err)))
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

	stat, err := s.svr.GetTopicStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return api.FromTopic(resp, stat), nil
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

func (s *server) GetTopicMeta(ctx context.Context, req *api.TopicMetaReq) (resp *api.TopicMetaInfo, err error) {
	meta, err := s.svr.GetTopicMeta(ctx, req.AccountID, req.TopicID)
	if err != nil {
		return nil, err
	}

	return api.FromTopicMeta(meta), nil
}

func (s *server) GetTopicPermission(ctx context.Context, req *api.TopicPermissionReq) (resp *api.TopicPermissionInfo, err error) {
	isMember, role, editPermission, err := s.svr.GetTopicPermission(ctx, req.AccountID, req.TopicID)
	if err != nil {
		return nil, err
	}

	return &api.TopicPermissionInfo{
		IsMember:       isMember,
		MemberRole:     role,
		EditPermission: editPermission,
	}, nil
}

func (s *server) GetUserTopicsPaged(ctx context.Context, req *api.UserTopicsReq) (resp *api.UserTopicsResp, err error) {
	items, err := s.svr.GetUserTopicsPaged(ctx, req.AccountID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}
	resp = &api.UserTopicsResp{
		Items: make([]*api.TopicInfo, len(items)),
	}

	for i, v := range items {
		stat, err := s.svr.GetTopicStat(ctx, v.ID)
		if err != nil {
			return nil, err
		}

		resp.Items[i] = api.FromTopic(v, stat)
	}

	return
}
