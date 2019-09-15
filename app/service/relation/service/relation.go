package service

import (
	"context"
)

// Fans 分页获取粉丝列表
func (s *Service) Fans(c context.Context, aid int64, limit, offset int) (fans []int64, err error) {
	return
}

// Fans 分页获取关注列表
func (s *Service) Following(c context.Context, aid int64, limit, offset int) (following []int64, err error) {
	return
}

// Follow 关注
func (s *Service) Follow(c context.Context, aid int64, fid int64) (err error) {
	return
}

// Unfollow 取关
func (s *Service) Unfollow(c context.Context, aid int64, fid int64) (err error) {
	return
}
