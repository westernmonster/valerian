package permit

import (
	"fmt"

	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
	"valerian/library/net/rpc/warden"

	mng "valerian/app/service/identify/api/grpc"

	"github.com/pkg/errors"
)

const (
	// _sessIDKey             = "_AJSESSIONID"
	_sessUIDKey   = "uid"      // manager user_id
	_sessUnameKey = "username" // LDAP username
	// CtxPermissions will be set into ctx.
	CtxPermissions = "permissions"
)

// permissions .
type permissions struct {
	UID   int64    `json:"uid"`
	Perms []string `json:"perms"`
}

// Permit is an auth middleware.
type Permit struct {
	sm *SessionManager // user Session

	mng.IdentifyClient // mng grpc client
}

type Config struct {
	Session     *SessionConfig
	IdentifyRPC *warden.ClientConfig
}

//Verify only export Verify function because of less configure
type Verify interface {
	Verify() mars.HandlerFunc
}

// New .
func New(c *Config) *Permit {
	identifyClient, err := mng.NewClient(c.IdentifyRPC)
	if err != nil {
		panic(errors.WithMessage(err, "Failed to dial mng rpc server"))
	}
	return &Permit{
		IdentifyClient: identifyClient,
		sm:             newSessionManager(c.Session),
	}
}

// Verify 验证管理员用户是否登录
func (p *Permit) Verify() mars.HandlerFunc {
	return func(ctx *mars.Context) {
		sid, uid, username, err := p.auth(ctx)
		if err != nil {
			fmt.Printf("permit.Verify() err(%+v)\n", err)
			ctx.JSON(nil, ecode.Unauthorized)
			ctx.Abort()
			return
		}
		ctx.Set(_sessUnameKey, username)
		ctx.Set(_sessUIDKey, uid)
		if md, ok := metadata.FromContext(ctx); ok {
			md[metadata.Uid] = uid
		}

		p.sm.setHTTPCookie(ctx, p.sm.c.CookieName, sid)
	}
}

// Permit2 same function as permit function but reply on grpc.
func (p *Permit) Permit(permit string) mars.HandlerFunc {
	return func(ctx *mars.Context) {
		sid, uid, username, err := p.auth(ctx)
		if err != nil {
			ctx.JSON(nil, ecode.Unauthorized)
			ctx.Abort()
			return
		}
		p.sm.setHTTPCookie(ctx, p.sm.c.CookieName, sid)
		ctx.Set(_sessUnameKey, username)
		ctx.Set(_sessUIDKey, uid)
		if md, ok := metadata.FromContext(ctx); ok {
			md[metadata.Uid] = uid
		}

		reply, err := p.Permissions(ctx, &mng.AdminPermissionReq{Sid: sid})
		if err != nil {
			if ecode.NothingFound.Equal(err) && permit != "" {
				ctx.JSON(nil, ecode.AccessDenied)
				ctx.Abort()
			}
			return
		}

		if len(reply.Perms) > 0 {
			ctx.Set(CtxPermissions, reply.Perms)
		}
		if !p.permit(permit, reply.Perms) {
			ctx.JSON(nil, ecode.AccessDenied)
			ctx.Abort()
			return
		}
	}
}

func (p *Permit) auth(ctx *mars.Context) (sid string, aid int64, username string, err error) {
	var cookieSid string
	sidVal, err := ctx.Request.Cookie(p.sm.c.CookieName)
	if err == nil {
		cookieSid = sidVal.Value
	}
	if cookieSid == "" {
		err = ecode.Unauthorized
		return
	}

	reply, err := p.AdminAuth(ctx, &mng.AdminAuthReq{Sid: cookieSid})
	if err != nil {
		log.For(ctx).Error(fmt.Sprintf("mng rpc AdminAuth error(%v)", err))
		return
	}

	sid = reply.Sid
	aid = reply.Aid
	username = reply.Username

	return
}

// permit check whether user has the access permission of the location.
func (p *Permit) permit(permit string, permissions []string) bool {
	if permit == "" {
		return true
	}
	for _, p := range permissions {
		if p == permit {
			// access the permit
			return true
		}
	}
	return false
}
