package service

import (
	"context"
	"time"

	"valerian/app/interface/fav/model"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

func (p *Service) Fav(c context.Context, arg *model.ArgAddFav) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": arg.TargetType,
		"target_id":   arg.TargetID,
	}); err != nil {
		return
	} else if fav != nil {
		return
	}

	fav = &model.Fav{
		ID:         gid.NewID(),
		AccountID:  aid,
		TargetID:   arg.TargetID,
		TargetType: arg.TargetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddFav(c, p.d.DB(), fav); err != nil {
		return
	}

	return
}

func (p *Service) Unfav(c context.Context, arg *model.ArgDelFav) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": arg.TargetType,
		"target_id":   arg.TargetID,
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

func (p *Service) GetFavsPaged(c context.Context, targetType string, limit, offset int) (resp *model.FavListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data []*model.Fav
	if data, err = p.d.GetFavsPaged(c, p.d.DB(), aid, targetType, limit, offset); err != nil {
		return
	}

	resp = &model.FavListResp{
		Items:  make([]*model.FavItem, len(data)),
		Paging: &model.Paging{},
	}

	for i, v := range data {
		item := &model.FavItem{
			ID:         v.ID,
			TargetType: v.TargetType,
		}

		switch v.TargetType {
		case model.TargetTypeTopic:
			var t *topic.TopicInfo
			if t, err = p.d.GetTopic(c, v.TargetID); err != nil {
				return
			}

			item.Topic = &model.TargetTopic{
				ID:           t.ID,
				Name:         t.Name,
				Introduction: t.Introduction,
				MemberCount:  int(t.MemberCount),
			}
			avatar := t.GetAvatarValue()
			item.Topic.Avatar = &avatar
			break
		case model.TargetTypeArticle:
			break
		case model.TargetTypeDiscussion:
			var d *discuss.DiscussionInfo
			if d, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				return
			}
			item.Discussion = &model.TargetDiscuss{
				ID: d.ID,
			}
			break
		case model.TargetTypeRevise:
			break
		}
		resp.Items[i] = item
	}

	return
}
