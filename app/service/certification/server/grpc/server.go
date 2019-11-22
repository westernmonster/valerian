package server

import (
	"context"
	"fmt"
	"valerian/app/service/certification/api"
	"valerian/app/service/certification/model"
	"valerian/app/service/certification/service"
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
	api.RegisterCertificationServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
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
	err := p.svr.AuditWorkCert(c, &model.ArgAuditWorkCert{
		AccountID:   req.AccountID,
		ManagerID:   req.ManagerID,
		AuditResult: req.AuditResult,
		Approve:     req.Approve,
	})
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
