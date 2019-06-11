package service

import (
	"context"
	"fmt"
	"valerian/app/conf"
	"valerian/app/interface/certification/dao"
	"valerian/app/interface/certification/model"
	"valerian/library/cloudauth"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		AddIDCertification(c context.Context, node sqalx.Node, item *model.IDCertification) (err error)
		UpdateIDCertification(c context.Context, node sqalx.Node, item *model.IDCertification) (err error)
		DelIDCertification(c context.Context, node sqalx.Node, id int64) (err error)
		GetUserIDCertification(c context.Context, node sqalx.Node, aid int64) (item *model.IDCertification, err error)

		Ping(c context.Context) (err error)
		Close()
		DB() sqalx.Node
	}
	cloudauth interface {
		GetVerifyToken(c context.Context, ticketID string) (resp *cloudauth.GetVerifyTokenResponse, err error)
		GetStatus(c context.Context, ticketID string) (resp *cloudauth.GetStatusResponse, err error)
		SubmitVerification(c context.Context, ticketID string, realName, idcardNumber, idcardFrontImage, idcardBackImage string) (resp *cloudauth.SubmitVerificationResponse, err error)
		GetMaterials(c context.Context, ticketID string) (resp *cloudauth.GetMaterialsResponse, err error)
	}
	missch chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		missch: make(chan func(), 1024),
	}

	if aliClient, err := sdk.NewClientWithAccessKey("cn-hangzhou", c.Aliyun.AccessKeyID, c.Aliyun.AccessKeySecret); err != nil {
		log.Error(fmt.Sprintf("init aliyun client failed(%+v)", err))
		panic(err)
	} else {
		s.cloudauth = &cloudauth.CloudAuthClient{Client: aliClient}
	}
	go s.cacheproc()
	return
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return s.d.Ping(c)
}

// Close dao.
func (s *Service) Close() {
	s.d.Close()
}

func (s *Service) addCache(f func()) {
	select {
	case s.missch <- f:
	default:
		log.Warn("cacheproc chan full")
	}
}

// cacheproc is a routine for executing closure.
func (s *Service) cacheproc() {
	for {
		f := <-s.missch
		f()
	}
}
