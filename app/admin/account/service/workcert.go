package service

import (
	"valerian/app/admin/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
)

func (s *Service) WorkCert(c *mars.Context, arg *model.ArgWorkCert) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if _, err := s.d.SetWorkCert(c, arg, aid); err != nil {
		return err
	}
	return
}
