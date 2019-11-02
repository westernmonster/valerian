package service

import (
	"context"
	"valerian/app/interface/article/model"
)

func (p *Service) GetReviseFiles(c context.Context, reviseID int64) (items []*model.ReviseFileResp, err error) {
	return
}

func (p *Service) SaveReviseFiles(c context.Context, arg *model.ArgSaveReviseFiles) (err error) {
	return
}
