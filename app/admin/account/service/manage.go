package service

import (
	"valerian/app/admin/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
)

func (s *Service) AccountLock(c *mars.Context, arg *model.ArgAccountLock) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if _, err := s.d.SetAccountLock(c, aid, arg.AccountID); err != nil {
		return err
	}
	return
}

func (s *Service) AccountUnlock(c *mars.Context, arg *model.ArgAccountUnlock) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if _, err := s.d.SetAccountUnlock(c, aid, arg.AccountID); err != nil {
		return err
	}
	return
}
