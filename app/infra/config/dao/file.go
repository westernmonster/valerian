package dao

import "valerian/app/infra/config/model"

// SetFile set config file.
func (d *Dao) SetFile(name string, conf *model.Content) (err error) {
	return
}

// File return config file.
func (d *Dao) File(name string) (res *model.Content, err error) {
	return
}

// DelFile delete file cache.
func (d *Dao) DelFile(name string) (err error) {
	return
}

// SetFileStr save string file.
func (d *Dao) SetFileStr(name string, val string) (err error) {
	return
}

// FileStr get string file.
func (d *Dao) FileStr(name string) (file string, err error) {
	return
}
