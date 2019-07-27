package service

import (
	"context"
	"sync"

	"valerian/app/infra/config/conf"
	"valerian/app/infra/config/dao"
	"valerian/app/infra/config/model"
	"valerian/library/database/sqalx"
	xtime "valerian/library/time"
)

// Service service.
type Service struct {
	d interface {
		GetAppsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.App, err error)
		GetApps(c context.Context, node sqalx.Node) (items []*model.App, err error)
		GetAppByID(c context.Context, node sqalx.Node, id int64) (item *model.App, err error)
		GetAppByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.App, err error)
		AddApp(c context.Context, node sqalx.Node, item *model.App) (err error)
		UpdateApp(c context.Context, node sqalx.Node, item *model.App) (err error)
		DelApp(c context.Context, node sqalx.Node, id int64) (err error)

		GetBuildsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Build, err error)
		GetBuildsByAppIDs(c context.Context, node sqalx.Node, appIDs []int64) (items []*model.Build, err error)
		GetBuilds(c context.Context, node sqalx.Node) (items []*model.Build, err error)
		GetBuildByID(c context.Context, node sqalx.Node, id int64) (item *model.Build, err error)
		GetBuildByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Build, err error)
		AddBuild(c context.Context, node sqalx.Node, item *model.Build) (err error)
		UpdateBuild(c context.Context, node sqalx.Node, item *model.Build) (err error)
		DelBuild(c context.Context, node sqalx.Node, id int64) (err error)

		GetConfigsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Config, err error)
		GetConfigs(c context.Context, node sqalx.Node) (items []*model.Config, err error)
		GetConfigByID(c context.Context, node sqalx.Node, id int64) (item *model.Config, err error)
		GetConfigsByIDs(c context.Context, node sqalx.Node, ids []int64) (items []*model.Config, err error)
		GetConfigByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Config, err error)
		AddConfig(c context.Context, node sqalx.Node, item *model.Config) (err error)
		UpdateConfig(c context.Context, node sqalx.Node, item *model.Config) (err error)
		DelConfig(c context.Context, node sqalx.Node, id int64) (err error)

		GetTagsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Tag, err error)
		GetTags(c context.Context, node sqalx.Node) (items []*model.Tag, err error)
		GetTagByID(c context.Context, node sqalx.Node, id int64) (item *model.Tag, err error)
		GetTagByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Tag, err error)
		AddTag(c context.Context, node sqalx.Node, item *model.Tag) (err error)
		UpdateTag(c context.Context, node sqalx.Node, item *model.Tag) (err error)
		DelTag(c context.Context, node sqalx.Node, id int64) (err error)

		GetForcesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Force, err error)
		GetForces(c context.Context, node sqalx.Node) (items []*model.Force, err error)
		GetForceByID(c context.Context, node sqalx.Node, id int64) (item *model.Force, err error)
		GetForceByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Force, err error)
		AddForce(c context.Context, node sqalx.Node, item *model.Force) (err error)
		UpdateForce(c context.Context, node sqalx.Node, item *model.Force) (err error)
		DelForce(c context.Context, node sqalx.Node, id int64) (err error)

		SetFile(c context.Context, name string, conf *model.Content) (err error)
		File(c context.Context, name string) (res *model.Content, err error)
		DelFile(c context.Context, name string) (err error)
		SetFileStr(c context.Context, name string, val string) (err error)
		FileStr(c context.Context, name string) (file string, err error)

		Hosts(c context.Context, svr string) (hosts []*model.Host, err error)
		SetHost(c context.Context, host *model.Host, svr string) (err error)
		ClearHost(c context.Context, svr string) (err error)

		Ping(c context.Context) (err error)
		Close()
		DB() sqalx.Node
	}
	PollTimeout xtime.Duration
	//config2 app
	aLock    sync.RWMutex
	bLock    sync.RWMutex
	apps     map[string]*model.App
	services map[string]*model.App
	//config2 tags
	tLock sync.RWMutex
	tags  map[string]*curTag

	eLock  sync.RWMutex
	events map[string]chan *model.Diff
	//force version
	fLock  sync.RWMutex
	forces map[string]int64

	//config2 tag force
	tfLock    sync.RWMutex
	forceType map[string]int
	// config2 tagID
	tagIDLock sync.RWMutex
	tagID     map[string]int64

	//config2 last force version
	lfvLock   sync.RWMutex
	lfvforces map[string]int64
}

// New new a service.
func New(c *conf.Config) (s *Service) {
	s = new(Service)
	s.d = dao.New(c)
	s.PollTimeout = c.PollTimeout
	s.tags = make(map[string]*curTag)
	s.apps = make(map[string]*model.App)
	s.services = make(map[string]*model.App)
	s.events = make(map[string]chan *model.Diff)
	s.forces = make(map[string]int64)
	s.forceType = make(map[string]int)
	s.tagID = make(map[string]int64)
	s.lfvforces = make(map[string]int64)
	return
}

// Ping check is ok.
func (s *Service) Ping(c context.Context) (err error) {
	return s.d.Ping(c)
}

// Close close resources.
func (s *Service) Close() {
	s.d.Close()
}
