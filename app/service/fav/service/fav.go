package service

import (
	"context"
	"time"
	"valerian/app/service/fav/api"
	"valerian/app/service/fav/model"
	"valerian/library/gid"
)

// IsFav 是否被收藏
func (p *Service) IsFav(c context.Context, aid int64, targetID int64, targetType string) (isFav bool, err error) {
	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		isFav = true
		return
	}
	return
}

// Fav 收藏
func (p *Service) Fav(c context.Context, aid, targetID int64, targetType string) (err error) {

	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		return
	}

	fav = &model.Fav{
		ID:         gid.NewID(),
		AccountID:  aid,
		TargetID:   targetID,
		TargetType: targetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddFav(c, p.d.DB(), fav); err != nil {
		return
	}

	return
}

// Unfav 取消收藏
func (p *Service) Unfav(c context.Context, aid, targetID int64, targetType string) (err error) {

	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav == nil {
		return
	} else {
		if err = p.d.DelFav(c, p.d.DB(), fav.ID); err != nil {
			return
		}
	}

	return
}

// GetFavIDsPaged 获取指定收藏类型的ID列表
func (p *Service) GetFavIDsPaged(c context.Context, arg *api.UserFavsReq) (ids []int64, err error) {
	return p.d.GetFavIDsPaged(c, p.d.DB(), arg.AccountID, arg.TargetType, arg.Limit, arg.Offset)
}

// GetFavsPaged 分页获取指定收藏类型列表
func (p *Service) GetFavsPaged(c context.Context, arg *api.UserFavsReq) (resp *api.FavsResp, err error) {
	resp = &api.FavsResp{
		Items: make([]*api.FavItem, 0),
	}

	var data []*model.Fav
	if data, err = p.d.GetFavsPaged(c, p.d.DB(), arg.AccountID, arg.TargetType, arg.Limit, arg.Offset); err != nil {
		return
	}

	for _, v := range data {
		resp.Items = append(resp.Items, &api.FavItem{
			ID:         v.ID,
			TargetID:   v.TargetID,
			TargetType: v.TargetType,
			AccountID:  v.AccountID,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		})
	}

	return

}
