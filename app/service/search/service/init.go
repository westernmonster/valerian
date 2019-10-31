package service

import (
	"context"
	"time"
	"valerian/app/service/search/model"
	"valerian/library/database/sqalx"
)

func (p *Service) Init(c context.Context) (err error) {
	c, cancelFunc := context.WithTimeout(c, 3*60*time.Second)
	defer cancelFunc()
	c = sqalx.NewContext(c, true)
	if err = p.d.CreateAccountIndices(c); err != nil {
		return
	}

	if err = p.d.CreateTopicIndices(c); err != nil {
		return
	}

	if err = p.d.CreateArticleIndices(c); err != nil {
		return
	}

	if err = p.d.CreateDiscussionIndices(c); err != nil {
		return
	}

	var accounts []*model.Account
	if accounts, err = p.d.GetAccounts(c, p.d.DB()); err != nil {
		return
	}

	iAcc := make([]*model.ESAccount, len(accounts))
	for i, v := range accounts {
		item := &model.ESAccount{
			ID:           v.ID,
			Mobile:       &v.Mobile,
			Email:        &v.Email,
			UserName:     &v.UserName,
			Role:         v.Role,
			Gender:       &v.Gender,
			BirthYear:    &v.BirthYear,
			BirthMonth:   &v.BirthMonth,
			BirthDay:     &v.BirthDay,
			Location:     &v.Location,
			Introduction: &v.Introduction,
			Avatar:       &v.Avatar,
			Source:       &v.Source,
			CreatedAt:    &v.CreatedAt,
			UpdatedAt:    &v.UpdatedAt,
		}

		idCert := bool(v.IDCert)
		workCert := bool(v.WorkCert)
		isOrg := bool(v.IsOrg)
		isVIP := bool(v.IsVip)
		item.IDCert = &idCert
		item.WorkCert = &workCert
		item.IsVIP = &isVIP
		item.IsOrg = &isOrg

		iAcc[i] = item
	}

	if err = p.d.BulkAccount2ES(c, iAcc); err != nil {
		return
	}

	var topics []*model.Topic
	if topics, err = p.d.GetTopics(c, p.d.DB()); err != nil {
		return
	}

	iTopic := make([]*model.ESTopic, len(topics))
	for i, v := range topics {
		item := &model.ESTopic{
			ID:              v.ID,
			Name:            &v.Name,
			Avatar:          &v.Avatar,
			Bg:              &v.Bg,
			Introduction:    &v.Introduction,
			ViewPermission:  &v.ViewPermission,
			EditPermission:  &v.EditPermission,
			JoinPermission:  &v.JoinPermission,
			CatalogViewType: &v.CatalogViewType,
			CreatedAt:       &v.CreatedAt,
			UpdatedAt:       &v.UpdatedAt,
		}

		allowDiscuss := bool(v.AllowDiscuss)
		allowChat := bool(v.AllowChat)
		isPrivate := bool(v.IsPrivate)
		item.AllowDiscuss = &allowDiscuss
		item.AllowChat = &allowChat
		item.IsPrivate = &isPrivate

		var acc *model.Account
		if acc, err = p.d.GetAccountByID(c, p.d.DB(), v.CreatedBy); err != nil {
			return
		} else if acc == nil {
			return
		}

		item.Creator = &model.ESCreator{
			ID:           acc.ID,
			UserName:     &acc.UserName,
			Avatar:       &acc.Avatar,
			Introduction: &acc.Introduction,
		}

		iTopic[i] = item
	}

	if err = p.d.BulkTopic2ES(c, iTopic); err != nil {
		return
	}

	var articles []*model.Article
	if articles, err = p.d.GetArticles(c, p.d.DB()); err != nil {
		return
	}

	iArticles := make([]*model.ESArticle, len(articles))
	for i, v := range articles {
		item := &model.ESArticle{
			ID:          v.ID,
			Title:       &v.Title,
			Content:     &v.Content,
			ContentText: &v.ContentText,
			CreatedAt:   &v.CreatedAt,
			UpdatedAt:   &v.UpdatedAt,
		}

		disableRevise := bool(v.DisableRevise)
		item.DisableRevise = &disableRevise
		disableComment := bool(v.DisableComment)
		item.DisableComment = &disableComment

		var acc *model.Account
		if acc, err = p.d.GetAccountByID(c, p.d.DB(), v.CreatedBy); err != nil {
			return
		} else if acc == nil {
			return
		}

		item.Creator = &model.ESCreator{
			ID:           acc.ID,
			UserName:     &acc.UserName,
			Avatar:       &acc.Avatar,
			Introduction: &acc.Introduction,
		}

		iArticles[i] = item

	}

	if err = p.d.BulkArticle2ES(c, iArticles); err != nil {
		return
	}

	var discussions []*model.Discussion
	if discussions, err = p.d.GetDiscussions(c, p.d.DB()); err != nil {
		return
	}

	iDiscuss := make([]*model.ESDiscussion, len(discussions))
	for i, v := range discussions {
		item := &model.ESDiscussion{
			ID:          v.ID,
			Title:       &v.Title,
			Content:     &v.Content,
			ContentText: &v.ContentText,
			CreatedAt:   &v.CreatedAt,
			UpdatedAt:   &v.UpdatedAt,
		}

		var acc *model.Account
		if acc, err = p.d.GetAccountByID(c, p.d.DB(), v.CreatedBy); err != nil {
			return
		} else if acc == nil {
			return
		}

		item.Creator = &model.ESCreator{
			ID:           acc.ID,
			UserName:     &acc.UserName,
			Avatar:       &acc.Avatar,
			Introduction: &acc.Introduction,
		}

		var t *model.Topic
		if t, err = p.d.GetTopicByID(c, p.d.DB(), v.TopicID); err != nil {
			return
		} else if t == nil {
			return
		}

		item.Topic = &model.ESDiscussionTopic{
			ID:           t.ID,
			Name:         &t.Name,
			Avatar:       &t.Avatar,
			Introduction: &t.Introduction,
		}

		if v.CategoryID != -1 {
			var cate *model.DiscussCategory
			if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
				return
			} else if cate == nil {
				return
			}
			item.Category = &model.ESDiscussionCategory{
				ID:   cate.ID,
				Name: &cate.Name,
				Seq:  &cate.Seq,
			}
		}

		iDiscuss[i] = item
	}

	if err = p.d.BulkDiscussion2ES(c, iDiscuss); err != nil {
		return
	}

	return
}
