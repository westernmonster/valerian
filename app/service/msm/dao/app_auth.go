package dao

import (
	"context"

	"valerian/app/service/msm/model"
)

// AllAppsInfo AllAppsInfo.
func (d *Dao) AllAppsInfo(c context.Context) (res map[int64]*model.AppInfo, err error) {
	return
}

// AllAppsAuth get all app auth info.
func (d *Dao) AllAppsAuth(c context.Context) (res map[int64]map[int64]*model.AppAuth, err error) {
	return
}
