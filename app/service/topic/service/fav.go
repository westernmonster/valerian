package service

import (
	"context"

	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
)

// isFav 是否已收藏
func (p *Service) isFav(c context.Context, node sqalx.Node, aid int64, targetID int64, targetType string) (isFav bool, err error) {
	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, node, map[string]interface{}{
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
