package dao

import (
	"valerian/app/service/msm/conf"
	"valerian/library/database/sqalx"
	"valerian/library/net/http/mars"
)

// Dao dao.
type Dao struct {
	client     *mars.Client
	db         sqalx.Node
	treeHost   string
	platformID string
}

// New new dao.
func New(c *conf.Config) *Dao {
	d := &Dao{
		db:         sqalx.NewMySQL(c.DB.Main),
		client:     mars.NewClient(c.HTTPClient),
		treeHost:   c.Tree.Host,
		platformID: c.Tree.PlatformID,
	}
	return d
}

// Close close mysql resource.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}