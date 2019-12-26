package service

import "context"

// Logout 退出
func (p *Service) Logout(c context.Context, aid int64, clientID string) (err error) {
	p.deleteToken(c, clientID, aid)
	return
}
