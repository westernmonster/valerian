package dao

import (
	"container/list"
	"context"
	"valerian/app/service/msm/model"
)

// Codes get all codes.
func (d *Dao) Codes(c context.Context) (codes map[int]string, lcode *model.Code, err error) {
	return
}

// Diff get change codes.
func (d *Dao) Diff(c context.Context, ver int64) (vers *list.List, err error) {
	return
}

// CodesLang get all codes.
func (d *Dao) CodesLang(c context.Context) (codes map[int]map[string]string, lcode *model.CodeLangs, err error) {
	return
}
