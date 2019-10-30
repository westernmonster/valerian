package server

import (
	"context"
	"fmt"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/service"
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

func (s *server) CreateTopic(ctx context.Context, arg *api.ArgCreateTopic) (*api.IDResp, error) {
	id, err := s.svr.CreateTopic(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.IDResp{ID: id}, nil
}

func (s *server) UpdateTopic(ctx context.Context, arg *api.ArgUpdateTopic) (*api.EmptyStruct, error) {
	err := s.svr.UpdateTopic(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) DelTopic(ctx context.Context, arg *api.IDReq) (*api.EmptyStruct, error) {
	err := s.svr.DelTopic(ctx, arg.ID)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) ChangeOwner(ctx context.Context, arg *api.ArgChangeOwner) (*api.EmptyStruct, error) {
	err := s.svr.ChangeOwner(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) SaveAuthTopics(ctx context.Context, arg *api.ArgSaveAuthTopics) (*api.EmptyStruct, error) {
	err := s.svr.SaveAuthTopics(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) GetAuthTopics(ctx context.Context, arg *api.IDReq) (*api.AuthTopicsResp, error) {
	items, err := s.svr.GetAuthTopics(ctx, arg.ID)
	if err != nil {
		return nil, err
	}
	return &api.AuthTopicsResp{Items: items}, nil
}

func (s *server) GetTopicResp(ctx context.Context, arg *api.IDReq) (*api.TopicResp, error) {
	resp, err := s.svr.GetTopicResp(ctx, arg.Aid, arg.ID, arg.Include)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) FollowTopic(ctx context.Context, arg *api.ArgTopicFollow) (*api.StatusResp, error) {
	status, err := s.svr.FollowTopic(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.StatusResp{Status: status}, nil
}

func (s *server) AuditFollow(ctx context.Context, arg *api.ArgAuditFollow) (*api.EmptyStruct, error) {
	err := s.svr.AuditFollow(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
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

	acc, err := s.svr.GetAccountBaseInfo(ctx, resp.CreatedBy)
	if err != nil {
		return nil, err
	}

	ret := api.FromTopic(resp, stat, acc)
	return ret, nil
}

func (s *server) GetTopicStat(ctx context.Context, req *api.TopicReq) (*api.TopicStat, error) {

	stat, err := s.svr.GetTopicStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	resp := &api.TopicStat{
		MemberCount:     int32(stat.MemberCount),
		ArticleCount:    int32(stat.ArticleCount),
		DiscussionCount: int32(stat.DiscussionCount),
	}
	return resp, err
}

func (s *server) GetTopicMemberRole(ctx context.Context, req *api.TopicMemberRoleReq) (resp *api.MemberRoleReply, err error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
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
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	meta, err := s.svr.GetTopicMeta(ctx, req.AccountID, req.TopicID)
	if err != nil {
		return nil, err
	}

	return meta, err
}

func (s *server) GetTopicPermission(ctx context.Context, req *api.TopicPermissionReq) (resp *api.TopicPermissionInfo, err error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
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

		m, err := s.svr.GetAccountBaseInfo(ctx, v.CreatedBy)
		if err != nil {
			return nil, err
		}

		resp.Items[i] = api.FromTopic(v, stat, m)
	}

	return
}

func (s *server) GetBelongsTopicIDs(ctx context.Context, req *api.AidReq) (*api.IDsResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	ids, err := s.svr.GetBelongsTopicIDs(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}

func (s *server) GetTopicMemberIDs(ctx context.Context, req *api.TopicReq) (*api.IDsResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	ids, err := s.svr.GetTopicMemberIDs(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}

func (s *server) GetAllTopics(ctx context.Context, req *api.EmptyStruct) (resp *api.AllTopicsResp, err error) {
	items, err := s.svr.GetAllTopics(ctx)
	if err != nil {
		return nil, err
	}
	resp = &api.AllTopicsResp{
		Items: make([]*api.TopicInfo, len(items)),
	}

	for i, v := range items {
		stat, err := s.svr.GetTopicStat(ctx, v.ID)
		if err != nil {
			return nil, err
		}

		m, err := s.svr.GetAccountBaseInfo(ctx, v.CreatedBy)
		if err != nil {
			return nil, err
		}

		resp.Items[i] = api.FromTopic(v, stat, m)
	}

	return
}
