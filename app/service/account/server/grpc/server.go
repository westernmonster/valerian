package server

import (
	"context"
	"fmt"

	"valerian/app/service/account/api"
	"valerian/app/service/account/model"
	"valerian/app/service/account/service"
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

func (s *server) AllAccountsIDs(ctx context.Context, req *api.AidReq) (*api.IDsResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
	}
	data, err := s.svr.GetAllAccountIDs(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.IDsResp{IDs: data}, nil
}

func (s *server) AllAccountsIDsPaged(ctx context.Context, req *api.AccountsPagedReq) (*api.IDsResp, error) {
	data, err := s.svr.GetAllAccountIDsPaged(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.IDsResp{IDs: data}, nil
}

func (s *server) AllAccountsPaged(ctx context.Context, req *api.AccountsPagedReq) (*api.AdminAccountsResp, error) {
	return s.svr.GetAllAccountsPaged(ctx, req)
}

func (p *server) RequestIDCert(c context.Context, req *api.AidReq) (*api.RequestIDCertResp, error) {
	if req.UseMaster {
		c = sqalx.NewContext(c, true)
		defer func() {
			c = sqalx.NewContext(c, false)
		}()
	}
	data, err := p.svr.RequestIDCert(c, req.Aid)
	if err != nil {
		return nil, err
	}

	resp := &api.RequestIDCertResp{
		CloudauthPageUrl: data.CloudauthPageUrl,
		STSToken: &api.STSToken{
			AccessKeyId:     data.STSToken.AccessKeyId,
			AccessKeySecret: data.STSToken.AccessKeySecret,
			Expiration:      data.STSToken.Expiration,
			EndPoint:        data.STSToken.EndPoint,
			BucketName:      data.STSToken.BucketName,
			Path:            data.STSToken.Path,
			Token:           data.STSToken.Token,
		},
		VerifyToken: &api.VerifyToken{
			Token:           data.VerifyToken.Token,
			DurationSeconds: int32(data.VerifyToken.DurationSeconds),
		},
	}

	return resp, nil
}

func (p *server) RefreshIDCertStatus(c context.Context, req *api.AidReq) (*api.IDCertStatus, error) {
	if req.UseMaster {
		c = sqalx.NewContext(c, true)
		defer func() {
			c = sqalx.NewContext(c, false)
		}()
	}
	status, err := p.svr.RefreshIDCertStatus(c, req.Aid)
	if err != nil {
		return nil, err
	}

	return &api.IDCertStatus{Status: int32(status)}, nil
}

func (p *server) GetIDCert(c context.Context, req *api.AidReq) (*api.IDCertInfo, error) {
	if req.UseMaster {
		c = sqalx.NewContext(c, true)
		defer func() {
			c = sqalx.NewContext(c, false)
		}()
	}
	data, err := p.svr.GetIDCert(c, req.Aid)
	if err != nil {
		return nil, err
	}

	return &api.IDCertInfo{
		AccountID:            data.AccountID,
		Status:               int32(data.Status),
		AuditConclusions:     data.AuditConclusions,
		Name:                 data.Name,
		IdentificationNumber: data.IdentificationNumber,
		IDCardType:           data.IDCardType,
		IDCardStartDate:      data.IDCardStartDate,
		IDCardExpiry:         data.IDCardExpiry,
		Address:              data.Address,
		Sex:                  data.Sex,
		IDCardFrontPic:       data.IDCardFrontPic,
		IDCardBackPic:        data.IDCardBackPic,
		FacePic:              data.FacePic,
		EthnicGroup:          data.EthnicGroup,
		CreatedAt:            data.CreatedAt,
		UpdatedAt:            data.UpdatedAt,
	}, nil
}

func (p *server) GetIDCertStatus(c context.Context, req *api.AidReq) (*api.IDCertStatus, error) {
	if req.UseMaster {
		c = sqalx.NewContext(c, true)
		defer func() {
			c = sqalx.NewContext(c, false)
		}()
	}
	status, err := p.svr.GetIDCertStatus(c, req.Aid)
	if err != nil {
		return nil, err
	}

	return &api.IDCertStatus{Status: int32(status)}, nil
}

func (p *server) RequestWorkCert(c context.Context, req *api.WorkCertReq) (*api.EmptyStruct, error) {
	err := p.svr.RequestWorkCert(c, &model.ArgAddWorkCert{
		AccountID:  req.AccountID,
		WorkPic:    req.WorkPic,
		OtherPic:   req.OtherPic,
		Company:    req.Company,
		Department: req.Department,
		Position:   req.Position,
		ExpiresAt:  req.ExpiresAt,
	})
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (p *server) AuditWorkCert(c context.Context, req *api.AuditWorkCertReq) (*api.EmptyStruct, error) {
	err := p.svr.AuditWorkCert(c, req)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (p *server) GetWorkCert(c context.Context, req *api.AidReq) (*api.WorkCertInfo, error) {
	if req.UseMaster {
		c = sqalx.NewContext(c, true)
		defer func() {
			c = sqalx.NewContext(c, false)
		}()
	}
	data, err := p.svr.GetWorkCert(c, req.Aid)
	if err != nil {
		return nil, err
	}

	return &api.WorkCertInfo{
		AccountID:   data.AccountID,
		Status:      int32(data.Status),
		WorkPic:     data.WorkPic,
		OtherPic:    data.OtherPic,
		Company:     data.Company,
		Department:  data.Department,
		Position:    data.Position,
		ExpiresAt:   data.ExpiresAt,
		AuditResult: data.AuditResult,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

func (p *server) GetWorkCertStatus(c context.Context, req *api.AidReq) (*api.WorkCertStatus, error) {
	if req.UseMaster {
		c = sqalx.NewContext(c, true)
		defer func() {
			c = sqalx.NewContext(c, false)
		}()
	}
	status, err := p.svr.GetWorkCertStatus(c, req.Aid)
	if err != nil {
		return nil, err
	}

	return &api.WorkCertStatus{Status: int32(status)}, nil
}

func (p *server) GetWorkCertsPaged(c context.Context, req *api.WorkCertPagedReq) (*api.WorkCertPagedResp, error) {
	return p.svr.GetWorkCertificationsPaged(c, req)
}

func (p *server) GetWorkCertHistoriesPaged(c context.Context, req *api.WorkCertHistoriesPagedReq) (*api.WorkCertHistoriesPagedResp, error) {
	return p.svr.GetWorkCertHistoriesPaged(c, req)
}
