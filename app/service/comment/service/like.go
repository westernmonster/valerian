package service

import (
	"context"
	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
)

func (p *Service) isLike(c context.Context, node sqalx.Node, aid int64, targetID int64, targetType string) (isLike bool, err error) {
	var fav *model.Like
	if fav, err = p.d.GetLikeByCond(c, node, map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		isLike = true
		return
	}
	return
}

func (p *Service) isDislike(c context.Context, node sqalx.Node, aid int64, targetID int64, targetType string) (isDislike bool, err error) {
	var fav *model.Dislike
	if fav, err = p.d.GetDislikeByCond(c, node, map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		isDislike = true
		return
	}

	return
}
