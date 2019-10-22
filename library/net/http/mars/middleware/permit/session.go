package permit

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
	"valerian/library/cache/memcache"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

func sessionKey(sid string) string {
	return fmt.Sprintf("sess_%d", sid)
}

// Session http session.
type Session struct {
	Sid string

	lock   sync.RWMutex
	Values map[string]interface{}
}

// SessionConfig config of Session.
type SessionConfig struct {
	SessionIDLength int
	CookieLifeTime  int
	CookieName      string
	Domain          string

	Memcache *memcache.Config
}

// SessionManager .
type SessionManager struct {
	mc *memcache.Pool // Session cache
	c  *SessionConfig
}

// newSessionManager .
func newSessionManager(c *SessionConfig) (s *SessionManager) {
	s = &SessionManager{
		mc: memcache.NewPool(c.Memcache),
		c:  c,
	}
	return
}

// SessionStart start session.
func (s *SessionManager) SessionStart(ctx *mars.Context) (si *Session) {
	// check manager Session id, if err or no exist need new one.
	if si, _ = s.cache(ctx); si == nil {
		si = s.newSession(ctx)
	}
	return
}

// SessionRelease flush session into store.
func (s *SessionManager) SessionRelease(ctx *mars.Context, sv *Session) {
	// set http cookie
	s.setHTTPCookie(ctx, s.c.CookieName, sv.Sid)
	// set mc
	conn := s.mc.Get(ctx)
	defer conn.Close()
	key := sessionKey(sv.Sid)
	item := &memcache.Item{
		Key:        key,
		Object:     sv,
		Flags:      memcache.FlagJSON,
		Expiration: int32(s.c.CookieLifeTime),
	}
	if err := conn.Set(item); err != nil {
		log.For(ctx).Error(fmt.Sprintf("SessionManager set error(%s,%v)", key, err))
	}
}

// SessionDestroy destroy session.
func (s *SessionManager) SessionDestroy(ctx *mars.Context, sv *Session) {
	conn := s.mc.Get(ctx)
	defer conn.Close()
	if err := conn.Delete(sessionKey(sv.Sid)); err != nil {
		log.For(ctx).Error(fmt.Sprintf("SessionManager delete error(%s,%v)", sv.Sid, err))
	}
}

func (s *SessionManager) cache(ctx *mars.Context) (res *Session, err error) {
	ck, err := ctx.Request.Cookie(s.c.CookieName)
	if err != nil || ck == nil {
		return
	}
	sid := ck.Value
	key := sessionKey(sid)
	// get from cache
	conn := s.mc.Get(ctx)
	defer conn.Close()
	r, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(ctx).Error(fmt.Sprintf("conn.Get(%s) error(%v)", sid, err))
		return
	}
	res = &Session{}
	if err = conn.Scan(r, res); err != nil {
		log.For(ctx).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(r.Value), err))
	}
	return
}

func (s *SessionManager) newSession(ctx context.Context) (res *Session) {
	b := make([]byte, s.c.SessionIDLength)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return nil
	}
	res = &Session{
		Sid:    hex.EncodeToString(b),
		Values: make(map[string]interface{}),
	}
	return
}

func (s *SessionManager) setHTTPCookie(ctx *mars.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		Path:     "/",
		HttpOnly: true,
		Domain:   s.c.Domain,
	}
	cookie.MaxAge = s.c.CookieLifeTime
	cookie.Expires = time.Now().Add(time.Duration(s.c.CookieLifeTime) * time.Second)
	http.SetCookie(ctx.Writer, cookie)
}

// Get get value by key.
func (s *Session) Get(key string) (value interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	value = s.Values[key]
	return
}

// Set set value into session.
func (s *Session) Set(key string, value interface{}) (err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Values[key] = value
	return
}
