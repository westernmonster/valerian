package service

import (
	"context"
	"encoding/hex"
	"math/rand"
	"valerian/library/net/http/mars/middleware/permit"
)

const (
	_sessUIDKey   = "uid"      // manager user_id
	_sessUnameKey = "username" // LDAP username
)

func (p *Service) session(ctx context.Context, sid string) (res *permit.Session) {
	if res, _ = p.d.Session(ctx, sid); res == nil {
		res = p.newSession(ctx)
	}
	return
}

func (p *Service) newSession(ctx context.Context) (res *permit.Session) {
	b := make([]byte, p.c.Session.SessionIDLength)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return
	}
	res = &permit.Session{
		Sid:    hex.EncodeToString(b),
		Values: make(map[string]interface{}),
	}
	return
}
