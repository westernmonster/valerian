package server

import (
	"context"
	"fmt"

	"valerian/app/service/account/api"
	"valerian/app/service/account/model"
	"valerian/app/service/account/service"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
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

func (s *server) SelfProfileInfo(ctx context.Context, req *api.AidReq) (*api.SelfProfile, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	resp, err := s.svr.GetSelfProfile(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *server) MemberInfo(ctx context.Context, req *api.AidReq) (*api.MemberInfoReply, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
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
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
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
		IsLock:       bool(v.IsLock),
		Deactive:     bool(v.Deactive),
	}

	return item, nil
}

func (s *server) GetAccountByMobile(ctx context.Context, req *api.MobileReq) (*api.DBAccount, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	v, err := s.svr.GetAccountByMobile(ctx, req.Prefix, req.Mobile)
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
		IsLock:       bool(v.IsLock),
		Deactive:     bool(v.Deactive),
		Password:     v.Password,
		Salt:         v.Salt,
	}

	return item, nil
}

func (s *server) GetAccountByEmail(ctx context.Context, req *api.EmailReq) (*api.DBAccount, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	v, err := s.svr.GetAccountByEmail(ctx, req.Email)
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
		IsLock:       bool(v.IsLock),
		Deactive:     bool(v.Deactive),
		Password:     v.Password,
		Salt:         v.Salt,
	}

	return item, nil
}

func (s *server) BasicInfo(ctx context.Context, req *api.AidReq) (*api.BaseInfoReply, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	resp, err := s.svr.BaseInfo(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *server) BaseInfos(ctx context.Context, req *api.AidsReq) (*api.BaseInfosReply, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	resp, err := s.svr.BatchBaseInfo(ctx, req.Aids)
	if err != nil {
		return nil, err
	}

	baseInfos := make(map[int64]*api.BaseInfoReply, len(resp))
	baseInfosReply := &api.BaseInfosReply{
		BaseInfos: baseInfos,
	}

	for k, v := range resp {
		baseInfos[k] = v
	}

	return baseInfosReply, nil
}

func (s *server) AccountStat(ctx context.Context, req *api.AidReq) (*api.AccountStatInfo, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	resp, err := s.svr.GetAccountStat(ctx, req.Aid)
	if err != nil {
		return nil, err
	}

	return api.FromStat(resp), nil
}

func (s *server) UpdateProfile(ctx context.Context, req *api.UpdateProfileReq) (*api.EmptyStruct, error) {
	err := s.svr.UpdateAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) ForgetPassword(ctx context.Context, req *api.ForgetPasswordReq) (*api.ForgetPasswordResp, error) {
	return s.svr.ForgetPassword(ctx, req)
}

func (s *server) ResetPassword(ctx context.Context, req *api.ResetPasswordReq) (*api.EmptyStruct, error) {
	err := s.svr.ResetPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) UpdateSetting(ctx context.Context, req *api.SettingReq) (*api.EmptyStruct, error) {
	err := s.svr.UpdateAccountSetting(ctx, req.Aid, req.Settings, req.Language)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) SettingInfo(ctx context.Context, req *api.AidReq) (*api.Setting, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
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

func (s *server) MobileExist(ctx context.Context, req *api.MobileReq) (*api.ExistResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	exist, err := s.svr.IsMobileExist(ctx, req.Prefix, req.Mobile)
	if err != nil {
		return nil, err
	}

	return &api.ExistResp{Exist: exist}, nil
}

func (s *server) EmailExist(ctx context.Context, req *api.EmailReq) (*api.ExistResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	exist, err := s.svr.IsEmailExist(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &api.ExistResp{Exist: exist}, nil
}

func (s *server) AddAccount(ctx context.Context, req *api.AddAccountReq) (*api.SelfProfile, error) {
	item := &model.Account{
		ID:           req.ID,
		Mobile:       req.Mobile,
		Email:        req.Email,
		UserName:     req.UserName,
		Role:         req.Role,
		Password:     req.Password,
		Salt:         req.Salt,
		Prefix:       req.Prefix,
		Gender:       req.Gender,
		BirthYear:    req.BirthYear,
		BirthMonth:   req.BirthMonth,
		BirthDay:     req.BirthDay,
		Location:     req.Location,
		Introduction: req.Introduction,
		Avatar:       req.Avatar,
		Source:       req.Source,
		IP:           req.IP,
		IDCert:       types.BitBool(req.IDCert),
		WorkCert:     types.BitBool(req.WorkCert),
		IsOrg:        types.BitBool(req.IsOrg),
		IsVip:        types.BitBool(req.IsVIP),
	}

	v, err := s.svr.AddAccount(ctx, item)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *server) AccountLock(ctx context.Context, req *api.AidReq) (*api.EmptyStruct, error) {
	err := s.svr.AccountLock(ctx, req.Aid)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) AccountUnlock(ctx context.Context, req *api.AidReq) (*api.EmptyStruct, error) {
	err := s.svr.AccountUnlock(ctx, req.Aid)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}
