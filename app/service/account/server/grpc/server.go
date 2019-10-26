package server

import (
	"context"
	"fmt"

	"valerian/app/service/account/api"
	"valerian/app/service/account/service"
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
	api.RegisterAccountServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) AllAccounts(ctx context.Context, req *api.EmptyStruct) (*api.AllAccountsResp, error) {
	items, err := s.svr.GetAllAccounts(ctx)
	if err != nil {
		return nil, err
	}

	resp := &api.AllAccountsResp{
		Items: make([]*api.DBAccount, len(items)),
	}

	for i, v := range items {
		item := &api.DBAccount{
			ID:           v.ID,
			Mobile:       v.Mobile,
			Email:        v.Email,
			UserName:     v.UserName,
			Role:         v.Role,
			Gender:       v.Gender,
			BirthYear:    v.BirthYear,
			BirthMonth:   v.BirthMonth,
			BirthDay:     v.BirthDay,
			Location:     v.Location,
			Introduction: v.Introduction,
			Avatar:       v.Avatar,
			Source:       int32(v.Source),
			IP:           v.IP,
			IDCert:       bool(v.IDCert),
			WorkCert:     bool(v.WorkCert),
			IsOrg:        bool(v.IsOrg),
			IsVIP:        bool(v.IsVip),
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
		resp.Items[i] = item
	}

	return resp, nil
}

func (s *server) SelfProfileInfo(ctx context.Context, req *api.AidReq) (*api.SelfProfile, error) {
	resp, err := s.svr.GetSelfProfile(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	reply := api.FromSelfProfile(resp)

	stat, err := s.svr.GetAccountStat(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	statInfo := api.FromStat(stat)
	reply.Stat = statInfo

	setting, err := s.svr.GetAccountSetting(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	settingInfo := api.FromSetting(setting)
	reply.Setting = settingInfo

	return reply, nil
}

func (s *server) MemberInfo(ctx context.Context, req *api.AidReq) (*api.MemberInfoReply, error) {
	resp, err := s.svr.GetMemberProfile(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	reply := api.FromProfileInfo(resp)

	stat, err := s.svr.GetAccountStat(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	statInfo := api.FromStat(stat)
	reply.Stat = statInfo

	return reply, nil
}

func (s *server) AccountInfo(ctx context.Context, req *api.AidReq) (*api.DBAccount, error) {
	v, err := s.svr.GetAccountByID(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	item := &api.DBAccount{
		ID:           v.ID,
		Mobile:       v.Mobile,
		Email:        v.Email,
		UserName:     v.UserName,
		Role:         v.Role,
		Gender:       v.Gender,
		BirthYear:    v.BirthYear,
		BirthMonth:   v.BirthMonth,
		BirthDay:     v.BirthDay,
		Location:     v.Location,
		Introduction: v.Introduction,
		Avatar:       v.Avatar,
		Source:       int32(v.Source),
		IP:           v.IP,
		IDCert:       bool(v.IDCert),
		WorkCert:     bool(v.WorkCert),
		IsOrg:        bool(v.IsOrg),
		IsVIP:        bool(v.IsVip),
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}

	return item, nil
}

func (s *server) BasicInfo(ctx context.Context, req *api.AidReq) (*api.BaseInfoReply, error) {
	resp, err := s.svr.BaseInfo(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	return api.FromBaseInfo(resp), nil
}

func (s *server) BaseInfos(ctx context.Context, req *api.AidsReq) (*api.BaseInfosReply, error) {
	resp, err := s.svr.BatchBaseInfo(ctx, req.Aids)
	if err != nil {
		return nil, err
	}

	baseInfos := make(map[int64]*api.BaseInfoReply, len(resp))
	baseInfosReply := &api.BaseInfosReply{
		BaseInfos: baseInfos,
	}

	for k, v := range resp {
		baseInfos[k] = api.FromBaseInfo(v)
	}

	return baseInfosReply, nil
}

func (s *server) AccountStat(ctx context.Context, req *api.AidReq) (*api.AccountStatInfo, error) {
	resp, err := s.svr.GetAccountStat(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	return api.FromStat(resp), nil
}

func (s *server) UpdateSetting(ctx context.Context, req *api.SettingReq) (*api.EmptyStruct, error) {
	err := s.svr.UpdateAccountSetting(ctx, req.Aid, req.Settings, req.Language)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) SettingInfo(ctx context.Context, req *api.AidReq) (*api.Setting, error) {
	data, err := s.svr.GetAccountSetting(ctx, req.Aid)
	if err != nil {
		return nil, err
	}
	return &api.Setting{
		ActivityLike:         data.ActivityLike,
		ActivityComment:      data.ActivityComment,
		ActivityFollowTopic:  data.ActivityFollowTopic,
		ActivityFollowMember: data.ActivityFollowMember,
		NotifyLike:           data.NotifyLike,
		NotifyComment:        data.NotifyComment,
		NotifyNewFans:        data.NotifyNewFans,
		NotifyNewMember:      data.NotifyNewMember,
		Language:             data.Language,
	}, nil
}
