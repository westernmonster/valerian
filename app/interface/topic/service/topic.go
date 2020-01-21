package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"valerian/app/interface/topic/model"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) CreateTopic(c context.Context, arg *model.ArgCreateTopic) (topicID int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
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
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
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

	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	// 检测查看权限
	if err = p.checkViewPermission(c, aid, topicID); err != nil {
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
		UpdatedAt:          t.UpdatedAt,
		HasCatalogTaxonomy: t.HasCatalogTaxonomy,
		MemberCount:        t.Stat.MemberCount,
		Members:            make([]*model.TopicMemberResp, 0),
		Catalogs:           make([]*model.TopicRootCatalog, 0),
		AuthTopics:         make([]*model.AuthTopicResp, 0),
		DiscussCategories:  make([]*model.DiscussCategoryResp, 0),
	}

	inc := includeParam(include)

	if t.Stat != nil {
		item.TopicStat = &model.TopicStatResp{
			MemberCount:     t.Stat.MemberCount,
			ArticleCount:    t.Stat.ArticleCount,
			DiscussionCount: t.Stat.DiscussionCount,
		}
	}

	if inc["creator"] {
		if t.Creator != nil {
			item.Creator = &model.Creator{
				ID:           t.Creator.ID,
				Avatar:       t.Creator.Avatar,
				UserName:     t.Creator.UserName,
				Introduction: t.Creator.Introduction,
			}
		}
	}

	if dl, ok := c.Deadline(); ok {
		ctimeout := time.Until(dl)
		fmt.Println(ctimeout)
	}
	if inc["members"] {
		var data *topic.TopicMembersPagedResp
		if data, err = p.d.GetTopicMembersPaged(c, &topic.ArgTopicMembers{TopicID: topicID, Page: 1, PageSize: 9}); err != nil {
			return
		}

		if data.Data != nil {
			for _, v := range data.Data {
				item.Members = append(item.Members, &model.TopicMemberResp{
					AccountID: v.AccountID,
					Role:      v.Role,
					Avatar:    v.Avatar,
					UserName:  v.UserName,
				})
			}
		}
	}

	if dl, ok := c.Deadline(); ok {
		ctimeout := time.Until(dl)
		fmt.Println(ctimeout)
	}
	if inc["catalogs"] {
		fmt.Println(11111111)
		if dl, ok := c.Deadline(); ok {
			ctimeout := time.Until(dl)
			fmt.Println(ctimeout)
		}
		var resp *topic.CatalogsResp
		if resp, err = p.d.GetCatalogsHierarchy(c, &topic.IDReq{ID: topicID, Aid: aid}); err != nil {
			return
		}

		item.Catalogs = p.FromCatalogs(resp.Items)

	}

	if inc["auth_topics"] {
		fmt.Println(22222222)
		if dl, ok := c.Deadline(); ok {
			ctimeout := time.Until(dl)
			fmt.Println(ctimeout)
		}

		var resp *topic.AuthTopicsResp
		if resp, err = p.d.GetAuthTopics(c, &topic.IDReq{ID: topicID, Aid: aid}); err != nil {
			return
		}

		if resp.Items != nil {
			for _, v := range resp.Items {
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

	}

	if inc["discuss_categories"] {
		fmt.Println(33333333)
		if dl, ok := c.Deadline(); ok {
			ctimeout := time.Until(dl)
			fmt.Println(ctimeout)
		}
		var resp *discuss.CategoriesResp
		if resp, err = p.d.GetDiscussionCategories(c, topicID); err != nil {
			return
		}

		if resp.Items != nil {
			for _, v := range resp.Items {
				item.DiscussCategories = append(item.DiscussCategories, &model.DiscussCategoryResp{
					ID:      v.ID,
					TopicID: v.TopicID,
					Name:    v.Name,
					Seq:     v.Seq,
				})
			}
		}

	}

	if inc["meta"] {
		fmt.Println(44444444)
		var m *topic.TopicMetaInfo
		if m, err = p.d.GetTopicMeta(c, aid, topicID); err != nil {
			return
		}

		item.TopicMeta = &model.TopicMeta{
			CanFollow:    m.CanFollow,
			CanEdit:      m.CanEdit,
			Fav:          m.Fav,
			CanView:      m.CanView,
			FollowStatus: (m.FollowStatus),
			IsMember:     m.CanView,
			MemberRole:   m.MemberRole,
		}

	}

	return
}

func (p *Service) FormCatelogTopic(v *topic.TopicInfo) (resp *model.TargetTopic) {
	resp = &model.TargetTopic{
		ID:           v.ID,
		Name:         v.Name,
		Introduction: v.Introduction,
		Avatar:       v.Avatar,
		Creator:      nil,
	}
	if v.Stat != nil {
		resp.MemberCount = v.Stat.MemberCount
		resp.ArticleCount = v.Stat.ArticleCount
		resp.DiscussionCount = v.Stat.DiscussionCount
	}
	if v.Creator != nil {
		resp.Creator = &model.Creator{
			ID: v.Creator.ID,
		}
	}
	return
}

func (p *Service) FromCatalogArticle(v *topic.TargetArticle) (resp *model.TargetArticle) {
	resp = &model.TargetArticle{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		ChangeDesc:   v.ChangeDesc,
		LikeCount:    v.LikeCount,
		DislikeCount: v.DislikeCount,
		ReviseCount:  v.ReviseCount,
		CommentCount: v.CommentCount,
		RelationIDs:  make([]string, 0),
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
	}
	if v.RelationIDs != nil {
		for _, x := range v.RelationIDs {
			resp.RelationIDs = append(resp.RelationIDs, strconv.FormatInt(x, 10))
		}
	}
	return
}

func (p *Service) FromCatalogs(items []*topic.TopicRootCatalogInfo) (resp []*model.TopicRootCatalog) {
	resp = make([]*model.TopicRootCatalog, 0)
	for _, v := range items {
		root := &model.TopicRootCatalog{
			ID:        &v.ID,
			Name:      v.Name,
			Type:      v.Type,
			RefID:     v.RefID,
			IsPrimary: v.IsPrimary,
			Children:  make([]*model.TopicParentCatalog, 0),
		}
		if v.Article != nil {
			root.Article = p.FromCatalogArticle(v.Article)
		}

		if v.Topic != nil {
			root.Topic = p.FormCatelogTopic(v.Topic)
		}

		if v.Children != nil {
			for _, x := range v.Children {
				parent := &model.TopicParentCatalog{
					ID:        &x.ID,
					Name:      x.Name,
					Type:      x.Type,
					RefID:     x.RefID,
					IsPrimary: v.IsPrimary,
					Children:  make([]*model.TopicChildCatalog, 0),
				}
				if x.Article != nil {
					parent.Article = p.FromCatalogArticle(x.Article)
				}
				if x.Topic != nil {
					parent.Topic = p.FormCatelogTopic(x.Topic)
				}
				if x.Children != nil {
					for _, j := range x.Children {
						child := &model.TopicChildCatalog{
							ID:        &j.ID,
							Name:      j.Name,
							Type:      j.Type,
							RefID:     j.RefID,
							IsPrimary: v.IsPrimary,
						}
						if j.Article != nil {
							child.Article = p.FromCatalogArticle(j.Article)
						}
						if j.Topic != nil {
							child.Topic = p.FormCatelogTopic(j.Topic)
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
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if err = p.d.DelTopic(c, &topic.IDReq{ID: topicID, Aid: aid}); err != nil {
		return
	}
	return
}

func (p *Service) GetTopicBasicInfo(c context.Context, topicID int64) (item *model.TopicBasicInfo, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var t *topic.TopicResp
	if t, err = p.d.GetTopicResp(c, &topic.IDReq{ID: topicID, Aid: aid, Include: ""}); err != nil {
		return
	}

	item = &model.TopicBasicInfo{
		ID:               t.ID,
		Name:             t.Name,
		Bg:               t.Bg,
		Avatar:           t.Avatar,
		Introduction:     t.Introduction,
		CatalogViewType:  t.CatalogViewType,
		TopicHome:        t.TopicHome,
		IsPrivate:        t.IsPrivate,
		AllowChat:        t.AllowChat,
		AllowDiscuss:     t.AllowDiscuss,
		MuteNotification: t.MuteNotification,
		EditPermission:   t.EditPermission,
		JoinPermission:   t.JoinPermission,
		ViewPermission:   t.ViewPermission,
		Important:        t.Important,
		CreatedAt:        t.CreatedAt,
		Members:          make([]*model.TopicMemberResp, 0),
		AuthTopics:       make([]*model.AuthTopicResp, 0),
	}

	if t.Creator != nil {
		item.Creator = &model.Creator{
			ID:           t.Creator.ID,
			Avatar:       t.Creator.Avatar,
			UserName:     t.Creator.UserName,
			Introduction: t.Creator.Introduction,
		}
	}

	if t.Stat != nil {
		item.TopicStat = &model.TopicStatResp{
			MemberCount:     t.Stat.MemberCount,
			ArticleCount:    t.Stat.ArticleCount,
			DiscussionCount: t.Stat.DiscussionCount,
		}
	}

	var data *topic.TopicMembersPagedResp
	if data, err = p.d.GetTopicMembersPaged(c, &topic.ArgTopicMembers{TopicID: topicID, Page: 1, PageSize: 9}); err != nil {
		return
	}

	if data.Data != nil {
		for _, v := range data.Data {
			item.Members = append(item.Members, &model.TopicMemberResp{
				AccountID: v.AccountID,
				Role:      v.Role,
				Avatar:    v.Avatar,
				UserName:  v.UserName,
			})
		}
	}

	var resp *topic.AuthTopicsResp
	if resp, err = p.d.GetAuthTopics(c, &topic.IDReq{ID: topicID, Aid: aid}); err != nil {
		return
	}

	if resp.Items != nil {
		for _, v := range resp.Items {
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

	var m *topic.TopicMetaInfo
	if m, err = p.d.GetTopicMeta(c, aid, topicID); err != nil {
		return
	}

	item.TopicMeta = &model.TopicMeta{
		CanFollow:    m.CanFollow,
		CanEdit:      m.CanEdit,
		Fav:          m.Fav,
		CanView:      m.CanView,
		FollowStatus: (m.FollowStatus),
		IsMember:     m.CanView,
		MemberRole:   m.MemberRole,
	}

	return
}
