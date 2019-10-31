package service

import (
	"context"

	"valerian/app/admin/topic/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) CreateTopic(c context.Context, arg *model.ArgCreateTopic) (topicID int64, err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &topic.ArgCreateTopic{
		Aid:             aid,
		Name:            arg.Name,
		Introduction:    arg.Introduction,
		CatalogViewType: arg.CatalogViewType,
		AllowDiscuss:    arg.AllowDiscuss,
		AllowChat:       arg.AllowChat,
	}
	if arg.Avatar != nil {
		item.Avatar = &topic.ArgCreateTopic_AvatarValue{*arg.Avatar}
	}
	if arg.Bg != nil {
		item.Bg = &topic.ArgCreateTopic_BgValue{*arg.Bg}
	}
	if topicID, err = p.d.CreateTopic(c, item); err != nil {
		return
	}

	return
}

func (p *Service) UpdateTopic(c context.Context, arg *model.ArgUpdateTopic) (err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &topic.ArgUpdateTopic{
		Aid: aid,
		ID:  arg.ID,
	}
	if arg.Avatar != nil {
		item.Avatar = &topic.ArgUpdateTopic_AvatarValue{*arg.Avatar}
	}
	if arg.Bg != nil {
		item.Bg = &topic.ArgUpdateTopic_BgValue{*arg.Bg}
	}

	if arg.Name != nil {
		item.Name = &topic.ArgUpdateTopic_NameValue{*arg.Name}
	}

	if arg.Introduction != nil {
		item.Introduction = &topic.ArgUpdateTopic_IntroductionValue{*arg.Introduction}
	}

	if arg.CatalogViewType != nil {
		item.CatalogViewType = &topic.ArgUpdateTopic_CatalogViewTypeValue{*arg.CatalogViewType}
	}

	if arg.AllowChat != nil {
		item.AllowChat = &topic.ArgUpdateTopic_AllowChatValue{*arg.AllowChat}
	}

	if arg.AllowDiscuss != nil {
		item.AllowDiscuss = &topic.ArgUpdateTopic_AllowDiscussValue{*arg.AllowDiscuss}
	}

	if arg.IsPrivate != nil {
		item.IsPrivate = &topic.ArgUpdateTopic_IsPrivateValue{*arg.IsPrivate}
	}
	if arg.ViewPermission != nil {
		item.ViewPermission = &topic.ArgUpdateTopic_ViewPermissionValue{*arg.ViewPermission}
	}
	if arg.EditPermission != nil {
		item.EditPermission = &topic.ArgUpdateTopic_EditPermissionValue{*arg.EditPermission}
	}
	if arg.JoinPermission != nil {
		item.JoinPermission = &topic.ArgUpdateTopic_JoinPermissionValue{*arg.JoinPermission}
	}
	if arg.Important != nil {
		item.Important = &topic.ArgUpdateTopic_ImportantValue{*arg.Important}
	}
	if arg.MuteNotification != nil {
		item.MuteNotification = &topic.ArgUpdateTopic_MuteNotificationValue{*arg.MuteNotification}
	}

	if err = p.d.UpdateTopic(c, item); err != nil {
		return
	}

	return
}

func (p *Service) GetTopic(c context.Context, topicID int64, include string) (item *model.TopicResp, err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var t *topic.TopicResp
	if t, err = p.d.GetTopicResp(c, &topic.IDReq{ID: topicID, Aid: aid, Include: include}); err != nil {
		return
	}

	item = &model.TopicResp{
		ID:                 t.ID,
		Name:               t.Name,
		Bg:                 t.Bg,
		Avatar:             t.Avatar,
		Introduction:       t.Introduction,
		CatalogViewType:    t.CatalogViewType,
		TopicHome:          t.TopicHome,
		IsPrivate:          t.IsPrivate,
		AllowChat:          t.AllowChat,
		AllowDiscuss:       t.AllowDiscuss,
		MuteNotification:   t.MuteNotification,
		EditPermission:     t.EditPermission,
		JoinPermission:     t.JoinPermission,
		ViewPermission:     t.ViewPermission,
		Important:          t.Important,
		CreatedAt:          t.CreatedAt,
		HasCatalogTaxonomy: t.HasCatalogTaxonomy,
		MemberCount:        t.Stat.MemberCount,
		Members:            make([]*model.TopicMemberResp, 0),
		Catalogs:           make([]*model.TopicRootCatalog, 0),
		AuthTopics:         make([]*model.AuthTopicResp, 0),
		DiscussCategories:  make([]*model.DiscussCategoryResp, 0),
	}

	if t.Creator != nil {
		item.Creator = &model.Creator{
			ID:           t.Creator.ID,
			Avatar:       t.Creator.Avatar,
			UserName:     t.Creator.UserName,
			Introduction: t.Creator.Introduction,
		}
	}

	if t.Members != nil {
		for _, v := range t.Members {
			item.Members = append(item.Members, &model.TopicMemberResp{
				AccountID: v.AccountID,
				Role:      v.Role,
				Avatar:    v.Avatar,
				UserName:  v.UserName,
			})
		}
	}

	if t.Catalogs != nil {
		item.Catalogs = p.FromCatalogs(t.Catalogs)
	}

	if t.AuthTopics != nil {
		for _, v := range t.AuthTopics {
			item.AuthTopics = append(item.AuthTopics, &model.AuthTopicResp{
				ToTopicID:      v.ToTopicID,
				EditPermission: v.EditPermission,
				Permission:     v.Permission,
				MemberCount:    v.MemberCount,
				Avatar:         v.Avatar,
				Name:           v.Name,
			})
		}
	}

	if t.DiscussCategories != nil {
		for _, v := range t.DiscussCategories {
			item.DiscussCategories = append(item.DiscussCategories, &model.DiscussCategoryResp{
				ID:      v.ID,
				TopicID: v.TopicID,
				Name:    v.Name,
				Seq:     v.Seq,
			})
		}
	}

	if t.Members != nil {
		for _, v := range t.Members {
			item.Members = append(item.Members, &model.TopicMemberResp{
				AccountID: v.AccountID,
				Role:      v.Role,
				UserName:  v.UserName,
				Avatar:    v.Avatar,
			})
		}
	}

	return
}

func (p *Service) FromCatalogs(items []*topic.TopicRootCatalogInfo) (resp []*model.TopicRootCatalog) {
	resp = make([]*model.TopicRootCatalog, 0)
	for _, v := range items {
		root := &model.TopicRootCatalog{
			ID:       &v.ID,
			Name:     v.Name,
			Type:     v.Type,
			RefID:    v.RefID,
			Children: make([]*model.TopicParentCatalog, 0),
		}
		if v.Article != nil {
			root.Article = &model.TargetArticle{
				ID:           v.Article.ID,
				Title:        v.Article.Title,
				Excerpt:      v.Article.Excerpt,
				LikeCount:    v.Article.LikeCount,
				DislikeCount: v.Article.DislikeCount,
				ReviseCount:  v.Article.ReviseCount,
				CommentCount: v.Article.CommentCount,
				CreatedAt:    v.Article.CreatedAt,
				UpdatedAt:    v.Article.UpdatedAt,
				Creator: &model.Creator{
					ID:           v.Article.Creator.ID,
					Avatar:       v.Article.Creator.Avatar,
					UserName:     v.Article.Creator.UserName,
					Introduction: v.Article.Creator.Introduction,
				},
			}
		}

		if v.Children != nil {
			for _, x := range v.Children {
				parent := &model.TopicParentCatalog{
					ID:       &x.ID,
					Name:     x.Name,
					Type:     x.Type,
					RefID:    x.RefID,
					Children: make([]*model.TopicChildCatalog, 0),
				}
				if x.Article != nil {
					root.Article = &model.TargetArticle{
						ID:           x.Article.ID,
						Title:        x.Article.Title,
						Excerpt:      x.Article.Excerpt,
						LikeCount:    x.Article.LikeCount,
						DislikeCount: x.Article.DislikeCount,
						ReviseCount:  x.Article.ReviseCount,
						CommentCount: x.Article.CommentCount,
						CreatedAt:    x.Article.CreatedAt,
						UpdatedAt:    x.Article.UpdatedAt,
						Creator: &model.Creator{
							ID:           x.Article.Creator.ID,
							Avatar:       x.Article.Creator.Avatar,
							UserName:     x.Article.Creator.UserName,
							Introduction: x.Article.Creator.Introduction,
						},
					}
				}
				if x.Children != nil {
					for _, j := range x.Children {
						child := &model.TopicChildCatalog{
							ID:    &j.ID,
							Name:  j.Name,
							Type:  j.Type,
							RefID: j.RefID,
						}
						if j.Article != nil {
							child.Article = &model.TargetArticle{
								ID:           j.Article.ID,
								Title:        j.Article.Title,
								Excerpt:      j.Article.Excerpt,
								LikeCount:    j.Article.LikeCount,
								DislikeCount: j.Article.DislikeCount,
								ReviseCount:  j.Article.ReviseCount,
								CommentCount: j.Article.CommentCount,
								CreatedAt:    j.Article.CreatedAt,
								UpdatedAt:    j.Article.UpdatedAt,
								Creator: &model.Creator{
									ID:           j.Article.Creator.ID,
									Avatar:       j.Article.Creator.Avatar,
									UserName:     j.Article.Creator.UserName,
									Introduction: j.Article.Creator.Introduction,
								},
							}
						}
						parent.Children = append(parent.Children, child)
					}
				}
				root.Children = append(root.Children, parent)
			}
		}
		resp = append(resp, root)
	}
	return
}

func (p *Service) DelTopic(c context.Context, topicID int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if err = p.d.DelTopic(c, &topic.IDReq{ID: topicID, Aid: aid}); err != nil {
		return
	}
	return
}
