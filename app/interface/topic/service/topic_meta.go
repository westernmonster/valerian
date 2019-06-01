package service

import (
	"context"
	"valerian/app/interface/topic/model"
)

func (p *Service) GetTopicMeta(c context.Context, t *model.TopicResp) (meta *model.TopicMeta, err error) {
	meta = new(model.TopicMeta)
	// aid, ok := metadata.Value(c, metadata.Aid).(int64)
	// if !ok {
	// 	meta = nil
	// 	return
	// }

	// var account *model.Account
	// if account, err = p.getAccountByID(c, aid); err != nil {
	// 	return
	// }

	// if t.ViewPermission == model.ViewPermissionPublic {
	// 	meta.CanView = true
	// }

	return
}
