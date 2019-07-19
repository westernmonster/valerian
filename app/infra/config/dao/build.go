package dao

// BuildsByAppID get builds by app id.
func (d *Dao) BuildsByAppID(appID int64) (builds []string, err error) {
	return
}

// BuildsByAppIDs get builds by app id.
func (d *Dao) BuildsByAppIDs(appIDs []int64) (builds []string, err error) {
	return
}

// TagID get TagID by ID.
func (d *Dao) TagID(appID int64, build string) (tagID int64, err error) {
	return
}

// BuildID get build by ID.
func (d *Dao) BuildID(appID int64, build string) (buildID int64, err error) {
	return
}
