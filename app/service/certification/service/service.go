package service

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"valerian/app/service/certification/conf"
	"valerian/app/service/certification/dao"
	"valerian/library/cloudauth"
	"valerian/library/conf/env"
	"valerian/library/log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)

// Service struct of service
type Service struct {
	c         *conf.Config
	d         *dao.Dao
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

func includeParam(include string) (dic map[string]bool) {
	arr := strings.Split(include, ",")
	dic = make(map[string]bool)
	for _, v := range arr {
		dic[v] = true
	}

	return
}

func genURL(path string, param url.Values) (uri string, err error) {
	u, err := url.Parse(env.SiteURL + path)
	if err != nil {
		return
	}
	u.RawQuery = param.Encode()

	return u.String(), nil
}
