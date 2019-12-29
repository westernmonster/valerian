package service

import (
	"context"

	"valerian/app/interface/discuss/model"
	discussion "valerian/app/service/discuss/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetDiscussionFiles(c context.Context, discussionID int64) (items []*model.DiscussionFileResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data *discussion.DiscussionFilesResp
	if data, err = p.d.GetDiscussionFiles(c, aid, discussionID); err != nil {
		return
	}

	items = make([]*model.DiscussionFileResp, len(data.Items))

	for i, v := range data.Items {
		items[i] = &model.DiscussionFileResp{
			ID:        v.ID,
			FileName:  v.FileName,
			FileURL:   v.FileURL,
			PdfURL:    v.PdfURL,
			FileType:  v.FileType,
			Seq:       v.Seq,
			CreatedAt: v.CreatedAt,
		}
	}
	return
}

func (p *Service) SaveDiscussionFiles(c context.Context, arg *model.ArgSaveDiscussionFiles) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &discussion.ArgSaveDiscussionFiles{
		Aid:          aid,
		DiscussionID: arg.DiscussionID,
		Items:        make([]*discussion.ArgDiscussionFile, 0),
	}

	for _, v := range arg.Items {
		file := &discussion.ArgDiscussionFile{
			FileName: v.FileName,
			FileType: v.FileType,
			FileURL:  v.FileURL,
			Seq:      v.Seq,
		}
		if v.ID != nil {
			file.ID = &discussion.ArgDiscussionFile_IDValue{*v.ID}
		}
		req.Items = append(req.Items, file)
	}

	if err = p.d.SaveDiscussionFiles(c, req); err != nil {
		return
	}
	return
}
