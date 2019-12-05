package service

import (
	"valerian/app/admin/account/model"
	"valerian/library/net/http/mars"
)

func (s *Service) AccountLock(c *mars.Context, arg *model.ArgAccountLock) (err error) {
	//aid, ok := metadata.Value(c, metadata.Aid).(int64)
	//if !ok {
	//	err = ecode.AcquireAccountIDFailed
	//	return
	//}
	//fmt.Println("AccountLock Metadata %v", metadata.String(c, "aid"))
	if _, err := s.d.SetAccountLock(c, arg.AccountID); err != nil {
		return err
	}
	return
}

func (s *Service) AccountUnlock(c *mars.Context, arg *model.ArgAccountUnlock) (err error) {
	//aid, ok := metadata.Value(c, metadata.Aid).(int64)
	//if !ok {
	//	err = ecode.AcquireAccountIDFailed
	//	return
	//}
	if _, err := s.d.SetAccountUnlock(c, arg.AccountID); err != nil {
		return err
	}
	return
}
