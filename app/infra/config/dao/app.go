package dao

import "valerian/app/infra/config/model"

// AppByTree get token by Name.
func (d *Dao) AppByTree(zone, env string, treeID int64) (app *model.App, err error) {
	return
}

// AppsByNameEnv get token by Name.
func (d *Dao) AppsByNameEnv(name, env string) (apps []*model.DBApp, err error) {
	return
}

// AppGet ...
func (d *Dao) AppGet(zone, env, token string) (app *model.App, err error) {
	return
}
