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
	status, err := s.svr.Follow(ctx, arg)
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

func (s *server) GetCatalogTaxonomiesHierarchy(ctx context.Context, arg *api.IDReq) (*api.CatalogsResp, error) {
	items, err := s.svr.GetCatalogTaxonomiesHierarchy(ctx, arg.ID)
	if err != nil {
		return nil, err
	}
	return &api.CatalogsResp{Items: items}, nil
}

func (s *server) GetCatalogsHierarchy(ctx context.Context, arg *api.IDReq) (*api.CatalogsResp, error) {
	items, err := s.svr.GetCatalogsHierarchy(ctx, arg.ID)
	if err != nil {
		return nil, err
	}
	return &api.CatalogsResp{Items: items}, nil
}

func (s *server) SaveCatalogs(ctx context.Context, arg *api.ArgSaveCatalogs) (*api.EmptyStruct, error) {
	err := s.svr.SaveCatalogs(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) GetUserCanEditTopicIDs(ctx context.Context, arg *api.AidReq) (*api.IDsResp, error) {
	ids, err := s.svr.GetUserCanEditTopicIDs(ctx, arg.AccountID)
	if err != nil {
		return nil, err
	}
	return &api.IDsResp{IDs: ids}, nil
}

func (s *server) HasTaxonomy(ctx context.Context, arg *api.TopicReq) (*api.BoolResp, error) {
	has, err := s.svr.HasTaxonomy(ctx, arg.ID)
	if err != nil {
		return nil, err
	}
	return &api.BoolResp{Result: has}, nil
}

func (s *server) IsTopicMember(ctx context.Context, arg *api.ArgIsTopicMember) (*api.BoolResp, error) {
	has, err := s.svr.IsTopicMember(ctx, arg.AccountID, arg.TopicID)
	if err != nil {
		return nil, err
	}
	return &api.BoolResp{Result: has}, nil
}

func (s *server) HasInvite(ctx context.Context, arg *api.ArgHasInvite) (*api.BoolResp, error) {
	has, err := s.svr.HasInvited(ctx, arg.AccountID, arg.TopicID)
	if err != nil {
		return nil, err
	}
	return &api.BoolResp{Result: has}, nil
}

func (s *server) Invite(ctx context.Context, arg *api.ArgTopicInvite) (*api.EmptyStruct, error) {
	err := s.svr.Invite(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) ProcessInvite(ctx context.Context, arg *api.ArgProcessInvite) (*api.EmptyStruct, error) {
	err := s.svr.ProcessInvite(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) Leave(ctx context.Context, arg *api.TopicReq) (*api.EmptyStruct, error) {
	err := s.svr.Leave(ctx, arg.Aid, arg.ID)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) BulkSaveMembers(ctx context.Context, arg *api.ArgBatchSavedTopicMember) (*api.EmptyStruct, error) {
	err := s.svr.BulkSaveMembers(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) GetTopicMembersPaged(ctx context.Context, arg *api.ArgTopicMembers) (*api.TopicMembersPagedResp, error) {
	resp, err := s.svr.GetTopicMembersPaged(ctx, arg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) GetTopicInfo(ctx context.Context, req *api.TopicReq) (*api.TopicInfo, error) {
	resp, err := s.svr.GetTopicInfo(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
	if resp, err = s.svr.GetUserTopicsPaged(ctx, req.AccountID, int(req.Limit), int(req.Offset)); err != nil {
		return
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

func (s *server) GetFollowedTopicsIDs(ctx context.Context, req *api.AidReq) (*api.IDsResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	ids, err := s.svr.GetFollowedTopicsIDs(ctx, req.AccountID)
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

func (s *server) CanView(ctx context.Context, req *api.TopicReq) (*api.BoolResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	ret, err := s.svr.CanView(ctx, req.Aid, req.ID)
	if err != nil {
		return nil, err
	}

	return &api.BoolResp{Result: ret}, nil
}

func (s *server) CanEdit(ctx context.Context, req *api.TopicReq) (*api.BoolResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	ret, err := s.svr.CanEdit(ctx, req.Aid, req.ID)
	if err != nil {
		return nil, err
	}

	return &api.BoolResp{Result: ret}, nil
}

func (s *server) GetRecommendTopicsIDs(ctx context.Context, req *api.EmptyStruct) (*api.IDsResp, error) {
	ids, err := s.svr.GetRecommendTopicsIDs(ctx)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}

func (s *server) GetRecommendAuthTopicsIDs(ctx context.Context, req *api.IDsReq) (*api.IDsResp, error) {
	ids, err := s.svr.GetRecommendAuthTopicsIDs(ctx, req.IDs)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}

func (s *server) GetRecommendMemberIDs(ctx context.Context, req *api.IDsReq) (*api.IDsResp, error) {
	ids, err := s.svr.GetRecommendMemberIDs(ctx, req.IDs)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}

func (s *server) AddRecommendTopic(ctx context.Context, req *api.TopicReq) (*api.EmptyStruct, error) {
	err := s.svr.AddRecommendTopic(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) DelRecommendTopic(ctx context.Context, req *api.TopicReq) (*api.EmptyStruct, error) {
	err := s.svr.DelRecommendTopic(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) GetAuthed2CurrentTopicIDs(ctx context.Context, req *api.TopicReq) (*api.IDsResp, error) {
	ids, err := s.svr.GetAuthed2CurrentTopicIDs(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &api.IDsResp{IDs: ids}, nil
}
