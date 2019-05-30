package service

import (
	"context"
	"fmt"

	"valerian/app/conf"
	"valerian/app/interface/valcode/dao"
	"valerian/library/email"
	"valerian/library/log"
	"valerian/library/sms"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		SetMobileValcodeCache(c context.Context, vtype int, mobile, code string) (err error)
		MobileValcodeCache(c context.Context, vtype int, mobile string) (code string, err error)
		DelMobileCache(c context.Context, vtype int, mobile string) (err error)

		SetEmailValcodeCache(c context.Context, vtype int, mobile, code string) (err error)
		EmailValcodeCache(c context.Context, vtype int, mobile string) (code string, err error)
		DelEmailCache(c context.Context, vtype int, mobile string) (err error)

		Ping(c context.Context) (err error)
		Close()
	}
	sms interface {
		SendRegisterValcode(c context.Context, prefix, mobile string, valcode string) (err error)
		SendResetPasswordValcode(c context.Context, prefix, mobile string, valcode string) (err error)
		SendLoginValcode(c context.Context, prefix, mobile string, valcode string) (err error)
	}
	email interface {
		SendRegisterEmail(c context.Context, email string, valcode string) (err error)
		SendResetPasswordValcode(c context.Context, email string, valcode string) (err error)
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
		s.sms = &sms.SMSClient{Client: aliClient}
		s.email = &email.EmailClient{Client: aliClient}
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
