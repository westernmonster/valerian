package usecase

import (
	"strconv"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/database/sqlx/types"

	"github.com/ztrue/tracerr"

	"valerian/infrastructure/berr"
	"valerian/infrastructure/biz"
	"valerian/infrastructure/gid"
	"valerian/models"
	"valerian/modules/repo"
)

type TopicUsecase struct {
	sqalx.Node
	*sqlx.DB

	TopicRepository interface {
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.Topic, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.Topic, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.Topic) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.Topic) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)

		SearchTopics(node sqalx.Node, cond map[string]string) (items []*models.TopicSearchResult, err error)

		GetTopicVersions(node sqalx.Node, topicSetID int64) (items []*models.TopicVersion, err error)
	}

	TopicCategoryRepository interface {
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.TopicCategory, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.TopicCategory, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.TopicCategory, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.TopicCategory) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.TopicCategory) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}

	TopicMemberRepository interface {
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.TopicMember, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.TopicMember, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.TopicMember) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.TopicMember) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)

		GetTopicMembers(node sqalx.Node, topicID int64, limit int) (items []*models.TopicMember, err error)

		GetTopicMembersPaged(node sqalx.Node, topicID int64, page, pageSize int) (count int, items []*models.TopicMember, err error)

		GetTopicMembersCount(node sqalx.Node, topicID int64) (count int, err error)
	}

	TopicSetRepository interface {
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.TopicSet) (err error)
	}

	TopicFollowerRepository interface {
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.TopicFollower, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.TopicFollower) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.TopicFollower) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
	}

	TopicTypeRepository interface {
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int) (item *repo.TopicType, exist bool, err error)
	}

	TopicRelationRepository interface {
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.TopicRelation, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.TopicRelation, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.TopicRelation) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.TopicRelation) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)

		GetAllRelatedTopics(node sqalx.Node, topicID int64) (items []*models.RelatedTopicShort, err error)

		GetAllRelatedTopicsDetail(node sqalx.Node, topicID int64) (items []*models.RelatedTopic, err error)
	}
}

func (p *TopicUsecase) GetTopicMembersPaged(ctx *biz.BizContext, topicID int64, page, pageSize int) (resp *models.TopicMembersPagedResp, err error) {
	count, data, err := p.TopicMemberRepository.GetTopicMembersPaged(p.Node, topicID, page, pageSize)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	resp = &models.TopicMembersPagedResp{
		Data:     data,
		Count:    count,
		PageSize: pageSize,
	}

	return
}

func (p *TopicUsecase) FollowTopic(ctx *biz.BizContext, topicID int64, isFollowed bool) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.TopicFollowerRepository.GetByCondition(tx, map[string]string{
		"topic_id":     strconv.FormatInt(topicID, 10),
		"followers_id": strconv.FormatInt(*ctx.AccountID, 10),
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if isFollowed {
		if exist {
			return
		}

		sid, errInner := gid.NextID()
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
		follower := &repo.TopicFollower{
			ID:          sid,
			TopicID:     topicID,
			FollowersID: *ctx.AccountID,
		}

		errInner = p.TopicFollowerRepository.Insert(tx, follower)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

	} else {
		if !exist {
			return
		}

		errInner := p.TopicFollowerRepository.Delete(tx, item.ID)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicUsecase) SeachTopics(ctx *biz.BizContext, topicID int64, query string) (items []*models.TopicSearchResult, err error) {
	items, err = p.TopicRepository.SearchTopics(p.Node, map[string]string{
		"id":    strconv.FormatInt(topicID, 10),
		"query": query,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicUsecase) GetAllRelatedTopics(ctx *biz.BizContext, topicID int64) (items []*models.RelatedTopic, err error) {
	items, err = p.TopicRelationRepository.GetAllRelatedTopicsDetail(p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	for _, v := range items {
		_, exist, errInner := p.TopicFollowerRepository.GetByCondition(p.Node, map[string]string{
			"topic_id":     strconv.FormatInt(v.TopicID, 10),
			"followers_id": strconv.FormatInt(*ctx.AccountID, 10),
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if exist {
			v.IsFollowed = true
		}

		count, errInner := p.TopicMemberRepository.GetTopicMembersCount(p.Node, v.TopicID)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		v.MembersCount = count
	}

	return
}

func (p *TopicUsecase) Get(ctx *biz.BizContext, topicID int64) (item *models.Topic, err error) {
	t, exist, err := p.TopicRepository.GetByID(p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("话题不存在")
		return
	}

	item = &models.Topic{
		ID:               t.ID,
		TopicSetID:       t.TopicSetID,
		Cover:            t.Cover,
		Name:             t.Name,
		Introduction:     t.Introduction,
		CategoryViewType: t.CategoryViewType,
		TopicType:        t.TopicType,
		TopicHome:        t.TopicHome,
		VersionName:      t.VersionName,
		VersionLanguage:  t.VersionLanguage,
		IsPrivate:        bool(t.IsPrivate),
		AllowChat:        bool(t.AllowChat),
		EditPermission:   t.EditPermission,
		ViewPermission:   t.ViewPermission,
		JoinPermission:   t.JoinPermission,
		Important:        bool(t.Important),
		MuteNotification: bool(t.MuteNotification),
		CreatedAt:        t.CreatedAt,
	}

	item.Members = make([]*models.TopicMember, 0)
	item.RelatedTopics = make([]*models.RelatedTopicShort, 0)
	item.Categories = make([]*models.TopicCategoryParentItem, 0)

	members, err := p.TopicMemberRepository.GetTopicMembers(p.Node, topicID, 10)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item.Members = members

	relatedTopics, err := p.TopicRelationRepository.GetAllRelatedTopics(p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item.RelatedTopics = relatedTopics

	categories, err := p.getCategoriesHierarchyOfAll(p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item.Categories = categories

	_, exist, err = p.TopicFollowerRepository.GetByCondition(p.Node, map[string]string{
		"topic_id":     strconv.FormatInt(topicID, 10),
		"followers_id": strconv.FormatInt(*ctx.AccountID, 10),
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if exist {
		item.IsFollowed = true
	}
	return
}

func (p *TopicUsecase) Create(ctx *biz.BizContext, req *models.CreateTopicReq) (id int64, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	id, err = gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	item := &repo.Topic{
		ID:               id,
		Name:             req.Name,
		Cover:            req.Cover,
		Introduction:     req.Introduction,
		IsPrivate:        types.BitBool(req.IsPrivate),
		AllowChat:        types.BitBool(req.AllowChat),
		EditPermission:   req.EditPermission,
		ViewPermission:   req.ViewPermission,
		JoinPermission:   req.JoinPermission,
		Important:        types.BitBool(req.Important),
		MuteNotification: types.BitBool(req.MuteNotification),
		CategoryViewType: req.CategoryViewType,
		TopicHome:        req.TopicHome,
		VersionName:      req.VersionName,
		VersionLanguage:  req.VersionLanguage,
		CreatedBy:        *ctx.AccountID,
	}

	_, exist, err := p.TopicTypeRepository.GetByID(tx, req.TopicType)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("话题类型不存在")
		return
	}

	item.TopicType = req.TopicType

	if req.TopicSetID != nil {
		_, exist, errInner := p.TopicRepository.GetByCondition(tx, map[string]string{
			"topic_set_id": strconv.FormatInt(*req.TopicSetID, 10),
			"version_name": req.VersionName,
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if exist {
			err = berr.Errorf("话题版本已经存在")
			return
		}

		item.TopicSetID = *req.TopicSetID
	} else {
		sid, errInner := gid.NextID()
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
		set := &repo.TopicSet{
			ID: sid,
		}
		item.TopicSetID = sid

		errInner = p.TopicSetRepository.Insert(tx, set)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
	}

	err = p.TopicRepository.Insert(tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkSaveCategories(tx, *ctx.AccountID, id, req.Categories)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkCreateMembers(tx, *ctx.AccountID, id, req)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkSaveRelations(tx, *ctx.AccountID, id, req.RelatedTopics)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicUsecase) Delete(ctx *biz.BizContext, topicID int64) (err error) {
	return
}

func (p *TopicUsecase) Update(ctx *biz.BizContext, topicID int64, req *models.UpdateTopicReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.TopicRepository.GetByID(tx, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("话题不存在")
		return
	}

	if req.TopicType != nil {
		_, exist, errInner := p.TopicTypeRepository.GetByID(tx, *req.TopicType)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
		if !exist {
			err = berr.Errorf("话题类型不存在")
			return
		}

		item.TopicType = *req.TopicType
	}

	if req.Cover != nil && *req.Cover != "" {
		item.Cover = *req.Cover
	}

	if req.Name != nil && *req.Name != "" {
		item.Name = *req.Name
	}

	if req.Introduction != nil && *req.Introduction != "" {
		item.Introduction = *req.Introduction
	}

	if req.VersionName != nil && *req.VersionName != "" {
		item.VersionName = *req.VersionName

		dbItem, exist, errInner := p.TopicRepository.GetByCondition(tx, map[string]string{
			"topic_set_id": strconv.FormatInt(item.TopicSetID, 10),
			"version_name": *req.VersionName,
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if exist && dbItem.ID != item.ID {
			err = berr.Errorf("话题版本已经存在")
			return
		}
	}

	if req.VersionLanguage != nil && *req.VersionLanguage != "" {
		item.VersionLanguage = *req.VersionLanguage
	}

	if req.JoinPermission != nil && *req.JoinPermission != "" {
		item.JoinPermission = *req.JoinPermission
	}

	if req.EditPermission != nil && *req.EditPermission != "" {
		item.EditPermission = *req.EditPermission
	}

	if req.ViewPermission != nil && *req.ViewPermission != "" {
		item.ViewPermission = *req.ViewPermission
	}

	if req.CategoryViewType != nil && *req.CategoryViewType != "" {
		item.CategoryViewType = *req.CategoryViewType
	}

	if req.TopicHome != nil && *req.TopicHome != "" {
		item.TopicHome = *req.TopicHome
	}

	if req.IsPrivate != nil {
		item.IsPrivate = types.BitBool(*req.IsPrivate)
	}

	if req.AllowChat != nil {
		item.AllowChat = types.BitBool(*req.AllowChat)
	}

	if req.Important != nil {
		item.Important = types.BitBool(*req.Important)
	}

	if req.MuteNotification != nil {
		item.MuteNotification = types.BitBool(*req.MuteNotification)
	}

	err = p.TopicRepository.Update(tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkSaveCategories(tx, *ctx.AccountID, item.ID, req.Categories)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkSaveRelations(tx, *ctx.AccountID, item.ID, req.RelatedTopics)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicUsecase) bulkSaveRelations(node sqalx.Node, accountID, topicID int64, relations []*models.RelatedTopicReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	for _, v := range relations {
		relation, exist, errInner := p.TopicRelationRepository.GetByCondition(tx, map[string]string{
			"from_topic_id": strconv.FormatInt(topicID, 10),
			"top_topic_id":  strconv.FormatInt(v.TopicID, 10),
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if exist {
			relation.Relation = v.Type
			errInner = p.TopicRelationRepository.Update(tx, relation)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
		} else {
			id, errInner := gid.NextID()
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
			item := &repo.TopicRelation{
				ID:          id,
				FromTopicID: topicID,
				ToTopicID:   v.TopicID,
				Relation:    v.Type,
			}
			errInner = p.TopicRelationRepository.Insert(tx, item)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicUsecase) bulkCreateMembers(node sqalx.Node, accountID, topicID int64, req *models.CreateTopicReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	for _, v := range req.Members {
		member, exist, errInner := p.TopicMemberRepository.GetByCondition(tx, map[string]string{
			"topic_id":   strconv.FormatInt(topicID, 10),
			"account_id": strconv.FormatInt(v.AccountID, 10),
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if exist {
			member.Role = v.Role
			errInner = p.TopicMemberRepository.Update(tx, member)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
		} else {
			id, errInner := gid.NextID()
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
			item := &repo.TopicMember{
				ID:        id,
				AccountID: v.AccountID,
				Role:      v.Role,
				TopicID:   topicID,
			}
			errInner = p.TopicMemberRepository.Insert(tx, item)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicUsecase) bulkSaveCategories(node sqalx.Node, accountID int64, topicID int64, categories []*models.TopicCategoryParentItem) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	dicMap := make(map[int64]bool)
	items, err := p.TopicCategoryRepository.GetAllByCondition(tx, map[string]string{
		"topic_id": strconv.FormatInt(topicID, 10),
	})
	for _, v := range items {
		dicMap[v.ID] = true
	}

	for _, parent := range categories {
		var parentID int64
		// update
		if parent.ID != nil {
			parentID = *parent.ID
			// exist in db
			if _, ok := dicMap[*parent.ID]; ok {
				// update record
				item, _, errInner := p.TopicCategoryRepository.GetByID(tx, *parent.ID)
				if errInner != nil {
					err = tracerr.Wrap(errInner)
					return
				}

				item.Name = parent.Name
				item.Seq = parent.Seq

				errInner = p.TopicCategoryRepository.Update(tx, item)
				if errInner != nil {
					err = tracerr.Wrap(errInner)
					return
				}
			} else {
				err = berr.Errorf("未找到该父级分类")
				return
			}
		} else {

			id, errInner := gid.NextID()
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}

			parentID = id

			item := &repo.TopicCategory{
				ID:        id,
				TopicID:   topicID,
				Name:      parent.Name,
				ParentID:  0,
				Seq:       parent.Seq,
				CreatedBy: accountID,
			}

			errInner = p.TopicCategoryRepository.Insert(tx, item)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
		}

		for _, child := range parent.Children {
			// update
			if child.ID != nil {
				// exist in db
				if _, ok := dicMap[*child.ID]; ok {
					// update record
					item, _, errInner := p.TopicCategoryRepository.GetByID(tx, *child.ID)
					if errInner != nil {
						err = tracerr.Wrap(errInner)
						return
					}

					item.Name = child.Name
					item.Seq = child.Seq
					item.ParentID = parentID

					errInner = p.TopicCategoryRepository.Update(tx, item)
					if errInner != nil {
						err = tracerr.Wrap(errInner)
						return
					}
				} else {
					err = berr.Errorf("未找到该父级分类")
					return
				}
			} else {

				id, errInner := gid.NextID()
				if errInner != nil {
					err = tracerr.Wrap(errInner)
					return
				}

				item := &repo.TopicCategory{
					ID:        id,
					TopicID:   topicID,
					Name:      child.Name,
					ParentID:  parentID,
					Seq:       child.Seq,
					CreatedBy: accountID,
				}

				errInner = p.TopicCategoryRepository.Insert(tx, item)
				if errInner != nil {
					err = tracerr.Wrap(errInner)
					return
				}
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicUsecase) getCategoriesHierarchyOfAll(node sqalx.Node, topicID int64) (items []*models.TopicCategoryParentItem, err error) {
	parents, err := p.TopicCategoryRepository.GetAllByCondition(node, map[string]string{
		"topic_id":  strconv.FormatInt(topicID, 10),
		"parent_id": "0",
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	// TODO: 是否有数据所有权的权限认证？
	items = make([]*models.TopicCategoryParentItem, 0)

	for _, v := range parents {
		parent := &models.TopicCategoryParentItem{
			ID:       &v.ID,
			Name:     v.Name,
			Seq:      v.Seq,
			Children: make([]*models.TopicCategoryChildItem, 0),
		}

		children, errInner := p.TopicCategoryRepository.GetAllByCondition(node, map[string]string{
			"topic_id":  strconv.FormatInt(topicID, 10),
			"parent_id": strconv.FormatInt(v.ID, 10),
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		for _, x := range children {
			child := &models.TopicCategoryChildItem{
				ID:   &x.ID,
				Name: x.Name,
				Seq:  x.Seq,
			}
			parent.Children = append(parent.Children, child)
		}

		items = append(items, parent)

	}

	return

}

func (p *TopicUsecase) GetTopicVersions(ctx *biz.BizContext, topicSetID int64) (items []*models.TopicVersion, err error) {
	return p.TopicRepository.GetTopicVersions(p.Node, topicSetID)
}

func (p *TopicUsecase) BulkSaveMembers(ctx *biz.BizContext, topicID int64, req *models.BatchSavedTopicMemberReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	for _, v := range req.Members {

		member, exist, errInner := p.TopicMemberRepository.GetByCondition(tx, map[string]string{
			"topic_id":   strconv.FormatInt(topicID, 10),
			"account_id": strconv.FormatInt(v.AccountID, 10),
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		switch v.Opt {
		case "C":
			if exist {
				continue
			}

			id, errInner := gid.NextID()
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
			item := &repo.TopicMember{
				ID:        id,
				AccountID: v.AccountID,
				Role:      v.Role,
				TopicID:   topicID,
			}
			errInner = p.TopicMemberRepository.Insert(tx, item)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}

			break
		case "U":
			if !exist {
				continue
			}
			member.Role = v.Role
			errInner = p.TopicMemberRepository.Update(tx, member)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
			break
		case "D":
			if !exist {
				continue
			}
			errInner = p.TopicMemberRepository.Delete(tx, member.ID)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
			break
		}
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}
