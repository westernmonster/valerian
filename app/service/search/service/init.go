package service

import (
	"context"
	"time"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/search/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
)

func (p *Service) Init(c context.Context) (err error) {
	c, cancelFunc := context.WithTimeout(c, 3*60*time.Second)
	defer cancelFunc()
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

	var accounts []*account.DBAccount
	if accounts, err = p.d.GetAllAccounts(c); err != nil {
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
		isVIP := bool(v.IsVIP)
		item.IDCert = &idCert
		item.WorkCert = &workCert
		item.IsVIP = &isVIP
		item.IsOrg = &isOrg

		iAcc[i] = item
	}

	if err = p.d.BulkAccount2ES(c, iAcc); err != nil {
		return
	}

	var topics []*topic.TopicInfo
	if topics, err = p.d.GetAllTopics(c); err != nil {
		return
	}

	for _, v := range topics {
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

		item.Creator = &model.ESCreator{
			ID:           v.Creator.ID,
			UserName:     &v.Creator.UserName,
			Avatar:       &v.Creator.Avatar,
			Introduction: &v.Creator.Introduction,
		}

		if err = p.d.PutTopic2ES(c, item); err != nil {
			return
		}
	}

	var articles []*article.DBArticle
	if articles, err = p.d.GetAllArticles(c); err != nil {
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
		disableComment := bool(v.DisableComment)
		item.DisableRevise = &disableRevise
		item.DisableComment = &disableComment

		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.CreatedBy); err != nil {
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

	var discussions []*discuss.DiscussionInfo
	if discussions, err = p.d.GetAllDiscussions(c); err != nil {
		return
	}

	for _, v := range discussions {
		item := &model.ESDiscussion{
			ID:    v.ID,
			Title: &v.Title,
			// Content:     &v.Content,
			// ContentText: &v.ContentText,
			CreatedAt: &v.CreatedAt,
			UpdatedAt: &v.UpdatedAt,
		}

		item.Creator = &model.ESCreator{
			ID:           v.Creator.ID,
			UserName:     &v.Creator.UserName,
			Avatar:       &v.Creator.Avatar,
			Introduction: &v.Creator.Introduction,
		}

		var t *model.Topic
		if t, err = p.d.GetTopicByID(c, p.d.DB(), v.TopicID); err != nil {
			return
		} else if t == nil {
			err = ecode.TopicNotExist
			return
		}

		item.Topic = &model.ESDiscussionTopic{
			ID:           t.ID,
			Name:         &t.Name,
			Avatar:       t.Avatar,
			Introduction: &t.Introduction,
		}

		if v.CategoryID != -1 {
			var cate *model.DiscussCategory
			if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
				return
			} else if cate == nil {
				err = ecode.DiscussCategoryNotExist
				return
			}

			item.Category = &model.ESDiscussionCategory{
				ID:   cate.ID,
				Name: &cate.Name,
				Seq:  &cate.Seq,
			}
		}

		if err = p.d.PutDiscussion2ES(c, item); err != nil {
			return
		}
	}

	return
}
