package permit

import (
	"net/url"

	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/net/rpc/warden"

	"github.com/pkg/errors"
)

const (
	_verifyURI             = "/api/session/verify"
	_permissionURI         = "/x/admin/manager/permission"
	_sessIDKey             = "_AJSESSIONID"
	_sessUIDKey            = "uid"      // manager user_id
	_sessUnKey             = "username" // LDAP username
	_defaultDomain         = ".bilibili.co"
	_defaultCookieName     = "mng-go"
	_defaultCookieLifeTime = 2592000
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
	verifyURI       string
	permissionURI   string
	dashboardCaller string
	dsClient        *mars.Client // dashboard client
	maClient        *mars.Client // manager-admin client

	sm *SessionManager // user Session

	mng.PermitClient // mng grpc client
}

//Verify only export Verify function because of less configure
type Verify interface {
	Verify() mars.HandlerFunc
}

// Config2 .
type Config2 struct {
	MngClient *warden.ClientConfig
	Session   *SessionConfig
}

// New2 .
func New2(c *warden.ClientConfig) *Permit {
	permitClient, err := mng.NewClient(c)
	if err != nil {
		panic(errors.WithMessage(err, "Failed to dial mng rpc server"))
	}
	return &Permit{
		PermitClient: permitClient,
		sm:           &SessionManager{},
	}
}

// Verify2 check whether the user has logged in.
func (p *Permit) Verify2() mars.HandlerFunc {
	return func(ctx *mars.Context) {
		sid, username, err := p.login2(ctx)
		if err != nil {
			ctx.JSON(nil, ecode.Unauthorized)
			ctx.Abort()
			return
		}
		ctx.Set(_sessUnKey, username)
		p.sm.setHTTPCookie(ctx, _defaultCookieName, sid)
	}
}

// Permit2 same function as permit function but reply on grpc.
func (p *Permit) Permit2(permit string) mars.HandlerFunc {
	return func(ctx *mars.Context) {
		sid, username, err := p.login2(ctx)
		if err != nil {
			ctx.JSON(nil, ecode.Unauthorized)
			ctx.Abort()
			return
		}
		p.sm.setHTTPCookie(ctx, _defaultCookieName, sid)
		ctx.Set(_sessUnKey, username)
		if md, ok := metadata.FromContext(ctx); ok {
			md[metadata.Username] = username
		}
		reply, err := p.Permissions(ctx, &mng.PermissionReq{Username: username})
		if err != nil {
			if ecode.NothingFound.Equal(err) && permit != "" {
				ctx.JSON(nil, ecode.AccessDenied)
				ctx.Abort()
			}
			return
		}
		ctx.Set(_sessUIDKey, reply.Uid)
		if md, ok := metadata.FromContext(ctx); ok {
			md[metadata.Uid] = reply.Uid
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

// login2 .
func (p *Permit) login2(ctx *mars.Context) (sid, uname string, err error) {
	var dsbsid, mngsid string
	dsbck, err := ctx.Request.Cookie(_sessIDKey)
	if err == nil {
		dsbsid = dsbck.Value
	}
	if dsbsid == "" {
		err = ecode.Unauthorized
		return
	}
	mngck, err := ctx.Request.Cookie(_defaultCookieName)
	if err == nil {
		mngsid = mngck.Value
	}
	reply, err := p.Login(ctx, &mng.LoginReq{Mngsid: mngsid, Dsbsid: dsbsid})
	if err != nil {
		log.Error("mng rpc Login error(%v)", err)
		return
	}
	sid = reply.Sid
	uname = reply.Username
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

// permissions get user's permisssions from manager-admin.
func (p *Permit) permissions(ctx *mars.Context, username string) (uid int64, perms []string, err error) {
	params := url.Values{}
	params.Set(_sessUnKey, username)
	var res struct {
		Code int         `json:"code"`
		Data permissions `json:"data"`
	}
	if err = p.maClient.Get(ctx, p.permissionURI, metadata.String(ctx, metadata.RemoteIP), params, &res); err != nil {
		log.Error("dashboard get permissions url(%s) error(%v)", p.permissionURI+"?"+params.Encode(), err)
		return
	}
	if ecode.Int(res.Code) != ecode.OK {
		log.Error("dashboard get permissions url(%s) error(%v)", p.permissionURI+"?"+params.Encode(), res.Code)
		err = ecode.Int(res.Code)
		return
	}
	perms = res.Data.Perms
	uid = res.Data.UID
	return
}
