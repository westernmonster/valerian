package grpc

import (
	"context"
	"fmt"
	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/service"
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
			log.For(ctx).Info("",
				zap.String("path", info.FullMethod),
				zap.String("caller", metadata.String(ctx, metadata.Caller)),
				zap.String("args", fmt.Sprintf("%v", req)),
				zap.String("error", fmt.Sprintf("%+v", err)))
		}
		return
	})
	api.RegisterIdentifyServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetTokenInfo(ctx context.Context, req *api.TokenReq) (*api.AuthReply, error) {
	return s.svr.GetTokenInfo(ctx, req.Token)
}

func (s *server) AdminAuth(c context.Context, req *api.AdminAuthReq) (*api.AdminAuthResp, error) {
	sid, aid, uname, err := s.svr.AuthAdmin(c, req.Sid)
	if err != nil {
		return nil, err
	}

	resp := &api.AdminAuthResp{
		Aid:      aid,
		Sid:      sid,
		Username: uname,
	}
	return resp, nil
}

func (s *server) Permissions(ctx context.Context, req *api.AdminPermissionReq) (*api.AdminPermissionResp, error) {
	return nil, nil
}

func (s *server) RenewToken(ctx context.Context, req *api.RenewTokenReq) (*api.RenewTokenResp, error) {
	data, err := s.svr.RenewToken(ctx, req.RefreshToken, req.ClientId)
	if err != nil {
		return nil, err
	}

	resp := &api.RenewTokenResp{
		Aid:          data.AccountID,
		AccessToken:  data.AccessToken,
		ExpiresIn:    int32(data.ExpiresIn),
		TokenType:    data.TokenType,
		RefreshToken: data.RefreshToken,
	}

	return resp, nil
}

func (s *server) Logout(ctx context.Context, req *api.LogoutReq) (*api.EmptyStruct, error) {
	err := s.svr.Logout(ctx, req.Aid, req.ClientId)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func (s *server) EmailLogin(ctx context.Context, req *api.EmailLoginReq) (*api.LoginResp, error) {
	return s.svr.EmailLogin(ctx, req)
}

func (s *server) MobileLogin(ctx context.Context, req *api.MobileLoginReq) (*api.LoginResp, error) {
	return s.svr.MobileLogin(ctx, req)
}

func (s *server) DigitLogin(ctx context.Context, req *api.DigitLoginReq) (*api.LoginResp, error) {
	return s.svr.DigitLogin(ctx, req)
}

func (s *server) EmailRegister(ctx context.Context, req *api.EmailRegisterReq) (*api.LoginResp, error) {
	return s.svr.EmailRegister(ctx, req)
}

func (s *server) MobileRegister(ctx context.Context, req *api.MobileRegisterReq) (*api.LoginResp, error) {
	return s.svr.MobileRegister(ctx, req)
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

func (s *server) AccountLock(ctx context.Context, req *api.LockReq) (*api.EmptyStruct, error) {
	err := s.svr.AccountLock(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) AccountUnlock(ctx context.Context, req *api.LockReq) (*api.EmptyStruct, error) {
	err := s.svr.AccountUnlock(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) Deactive(ctx context.Context, req *api.DeactiveReq) (*api.EmptyStruct, error) {
	err := s.svr.Deactive(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}

func (s *server) SetPassword(ctx context.Context, req *api.SetPasswordReq) (*api.EmptyStruct, error) {
	err := s.svr.SetPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.EmptyStruct{}, nil
}
