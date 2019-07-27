package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"valerian/app/infra/config/model"
	"valerian/library/log"
)

// SetFile set config file.
func (d *Dao) SetFile(c context.Context, name string, conf *model.Content) (err error) {
	b, err := json.Marshal(conf)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("json.Marshal(%v) error(%v)", conf, err))
		return
	}
	p := path.Join(d.pathCache, name)
	if err = ioutil.WriteFile(p, b, 0644); err != nil {
		log.For(c).Error(fmt.Sprintf("ioutil.WriteFile(%s) error(%v)", p, err))
	}
	return
}

// File return config file.
func (d *Dao) File(c context.Context, name string) (res *model.Content, err error) {
	p := path.Join(d.pathCache, name)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("ioutil.ReadFile(%s) error(%v)", p, err))
		return
	}
	res = &model.Content{}
	if err = json.Unmarshal(b, &res); err != nil {
		log.For(c).Error(fmt.Sprintf("json.Unmarshal(%s) error(%v)", b, err))
	}
	return
}

// DelFile delete file cache.
func (d *Dao) DelFile(c context.Context, name string) (err error) {
	p := path.Join(d.pathCache, name)
	if err = os.Remove(p); err != nil {
		log.For(c).Error(fmt.Sprintf("os.Remove(%s) error(%v)", p, err))
	}
	return
}

// SetFileStr save string file.
func (d *Dao) SetFileStr(c context.Context, name string, val string) (err error) {
	p := path.Join(d.pathCache, name)
	if err = ioutil.WriteFile(p, []byte(val), 0644); err != nil {
		log.For(c).Error(fmt.Sprintf("ioutil.WriteFile(%s) error(%v)", p, err))
	}
	return
}

// FileStr get string file.
func (d *Dao) FileStr(c context.Context, name string) (file string, err error) {
	p := path.Join(d.pathCache, name)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("ioutil.ReadFile(%s) error(%v)", p, err))
		return
	}
	file = string(b)
	return
}
