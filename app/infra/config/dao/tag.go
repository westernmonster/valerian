package dao

import "valerian/app/infra/config/model"

//Tags get builds by app id.
func (d *Dao) Tags(appID int64) (tags []*model.ReVer, err error) {
	return
}

//ConfIDs get tag by id.
func (d *Dao) ConfIDs(ID int64) (ids []int64, err error) {
	return
}

// TagForce get force by tag.
func (d *Dao) TagForce(ID int64) (force int8, err error) {
	return
}

// LastForce ...
func (d *Dao) LastForce(appID, buildID int64) (lastForce int64, err error) {
	return
}

//TagAll ...
func (d *Dao) TagAll(tagID int64) (tag *model.DBTag, err error) {
	return
}
