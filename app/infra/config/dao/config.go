package dao

import "valerian/app/infra/config/model"

// UpdateConfValue update config state/
func (d *Dao) UpdateConfValue(ID int64, value string) (err error) {
	return
}

// UpdateConfState update config state/
func (d *Dao) UpdateConfState(ID int64, state int8) (err error) {
	return
}

// ConfigsByIDs get Config by IDs.
func (d *Dao) ConfigsByIDs(ids []int64) (confs []*model.Value, err error) {
	return
}
