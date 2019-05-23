package usecase

import (
	"context"
	"strconv"
	"time"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/database/sqlx/types"
	"valerian/library/gid"

	"github.com/jinzhu/copier"
	"github.com/ztrue/tracerr"

	"valerian/infrastructure/berr"
	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
)

type TopicUsecase struct {
	sqalx.Node
	*sqlx.DB

	AccountRepository interface {
		// GetByID get a record by ID
		GetByID(ctx context.Context, node sqalx.Node, id int64) (item *repo.Account, exist bool, err error)
	}

	TopicRepository interface {
		// GetByID get a record by ID
		GetByID(ctx context.Context, node sqalx.Node, id int64) (item *repo.Topic, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *repo.Topic, exist bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.Topic) (err error)
		// Update update a exist record
		Update(ctx context.Context, node sqalx.Node, item *repo.Topic) (err error)

		Delete(ctx context.Context, node sqalx.Node, topicID int64) (err error)

		SearchTopics(ctx context.Context, node sqalx.Node, cond map[string]string) (items []*models.TopicSearchResult, err error)

		GetTopicVersions(ctx context.Context, node sqalx.Node, topicSetID int64) (items []*models.TopicVersion, err error)
	}

	TopicCategoryRepository interface {
		// GetAllByCondition get records by condition
		GetAllByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (items []*repo.TopicCategory, err error)
		// GetByID get a record by ID
		GetByID(ctx context.Context, node sqalx.Node, id int64) (item *repo.TopicCategory, exist bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.TopicCategory) (err error)
		// Update update a exist record
		Update(ctx context.Context, node sqalx.Node, item *repo.TopicCategory) (err error)
		// Delete logic delete a exist record
		Delete(ctx context.Context, node sqalx.Node, id int64) (err error)
	}

	TopicMemberRepository interface {
		// GetByCondition get a record by condition
		GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *repo.TopicMember, exist bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.TopicMember) (err error)
		// Update update a exist record
		Update(ctx context.Context, node sqalx.Node, item *repo.TopicMember) (err error)
		// Delete logic delete a exist record
		Delete(ctx context.Context, node sqalx.Node, id int64) (err error)

		GetAllTopicMembers(ctx context.Context, node sqalx.Node, topicID int64) (items []*repo.TopicMember, err error)

		GetTopicMembers(ctx context.Context, node sqalx.Node, topicID int64, limit int) (items []*models.TopicMember, err error)

		GetTopicMembersPaged(ctx context.Context, node sqalx.Node, topicID int64, page, pageSize int) (count int, items []*models.TopicMember, err error)

		GetTopicMembersCount(ctx context.Context, node sqalx.Node, topicID int64) (count int, err error)
	}

	TopicSetRepository interface {
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.TopicSet) (err error)
	}

	TopicFollowerRepository interface {
		// GetByCondition get a record by condition
		GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *repo.TopicFollower, exist bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.TopicFollower) (err error)
		// Delete logic delete a exist record
		Delete(ctx context.Context, node sqalx.Node, id int64) (err error)
	}

	TopicTypeRepository interface {
		// GetByID get a record by ID
		GetByID(ctx context.Context, node sqalx.Node, id int) (item *repo.TopicType, exist bool, err error)

		GetAll(ctx context.Context, node sqalx.Node) (items []*repo.TopicType, err error)
	}

	TopicRelationRepository interface {
		GetAllTopicRelations(ctx context.Context, node sqalx.Node, topicID int64) (items []*repo.TopicRelation, err error)
		// GetByCondition get a record by condition
		GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *repo.TopicRelation, exist bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.TopicRelation) (err error)
		// Update update a exist record
		Update(ctx context.Context, node sqalx.Node, item *repo.TopicRelation) (err error)

		GetAllRelatedTopics(ctx context.Context, node sqalx.Node, topicID int64) (items []*models.RelatedTopicShort, err error)

		GetAllRelatedTopicsDetail(ctx context.Context, node sqalx.Node, topicID int64) (items []*models.RelatedTopic, err error)
	}
}

func (p *TopicUsecase) GetTopicMembersPaged(c context.Context, ctx *biz.BizContext, topicID int64, page, pageSize int) (resp *models.TopicMembersPagedResp, err error) {
	count, data, err := p.TopicMemberRepository.GetTopicMembersPaged(c, p.Node, topicID, page, pageSize)
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

func (p *TopicUsecase) FollowTopic(c context.Context, ctx *biz.BizContext, topicID int64, isFollowed bool) (err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.TopicFollowerRepository.GetByCondition(c, tx, map[string]string{
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

		errInner = p.TopicFollowerRepository.Insert(c, tx, follower)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

	} else {
		if !exist {
			return
		}

		errInner := p.TopicFollowerRepository.Delete(c, tx, item.ID)
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

func (p *TopicUsecase) SeachTopics(c context.Context, ctx *biz.BizContext, query string) (items []*models.TopicSearchResult, err error) {
	items, err = p.TopicRepository.SearchTopics(c, p.Node, map[string]string{
		"query": query,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	for _, v := range items {
		versions, errInner := p.TopicRepository.GetTopicVersions(c, p.Node, v.TopicSetID)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		v.Versions = versions
	}
	return
}

func (p *TopicUsecase) GetAllRelatedTopics(c context.Context, ctx *biz.BizContext, topicID int64) (items []*models.RelatedTopic, err error) {
	items, err = p.TopicRelationRepository.GetAllRelatedTopicsDetail(c, p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	for _, v := range items {
		_, exist, errInner := p.TopicFollowerRepository.GetByCondition(c, p.Node, map[string]string{
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

		count, errInner := p.TopicMemberRepository.GetTopicMembersCount(c, p.Node, v.TopicID)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		v.MembersCount = count
	}

	return
}

func (p *TopicUsecase) Get(c context.Context, ctx *biz.BizContext, topicID int64) (item *models.Topic, err error) {
	return p.get(c, ctx, topicID)
}

func (p *TopicUsecase) get(c context.Context, ctx *biz.BizContext, topicID int64) (item *models.Topic, err error) {
	t, exist, err := p.TopicRepository.GetByID(c, p.Node, topicID)
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
		Bg:               t.Bg,
		Name:             t.Name,
		Introduction:     t.Introduction,
		CategoryViewType: t.CategoryViewType,
		TopicType:        t.TopicType,
		TopicHome:        t.TopicHome,
		VersionName:      t.VersionName,
		IsPrivate:        bool(t.IsPrivate),
		AllowChat:        bool(t.AllowChat),
		AllowDiscuss:     bool(t.AllowDiscuss),
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
	item.Versions = make([]*models.TopicVersion, 0)

	topicType, exist, err := p.TopicTypeRepository.GetByID(c, p.Node, t.TopicType)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if exist {
		item.TopicTypeName = topicType.Name
	}

	members, err := p.TopicMemberRepository.GetTopicMembers(c, p.Node, topicID, 10)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item.Members = members

	item.MembersCount, err = p.TopicMemberRepository.GetTopicMembersCount(c, p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	item.Versions, err = p.TopicRepository.GetTopicVersions(c, p.Node, t.TopicSetID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	relatedTopics, err := p.TopicRelationRepository.GetAllRelatedTopics(c, p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item.RelatedTopics = relatedTopics

	categories, err := p.getCategoriesHierarchyOfAll(c, p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item.Categories = categories

	_, exist, err = p.TopicFollowerRepository.GetByCondition(c, p.Node, map[string]string{
		"topic_id":     strconv.FormatInt(topicID, 10),
		"followers_id": strconv.FormatInt(*ctx.AccountID, 10),
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if exist {
		item.TopicMeta.IsFollowed = true
	}
	return
}

func (p *TopicUsecase) Create(c context.Context, ctx *biz.BizContext, req *models.CreateTopicReq) (id int64, err error) {
	tx, err := p.Node.Beginx(c)
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
		Bg:               req.Bg,
		Introduction:     req.Introduction,
		IsPrivate:        types.BitBool(req.IsPrivate),
		AllowChat:        types.BitBool(req.AllowChat),
		AllowDiscuss:     types.BitBool(req.AllowDiscuss),
		EditPermission:   req.EditPermission,
		ViewPermission:   req.ViewPermission,
		JoinPermission:   req.JoinPermission,
		Important:        types.BitBool(req.Important),
		MuteNotification: types.BitBool(req.MuteNotification),
		CategoryViewType: req.CategoryViewType,
		TopicHome:        req.TopicHome,
		VersionName:      req.VersionName,
		CreatedBy:        *ctx.AccountID,
	}

	_, exist, err := p.TopicTypeRepository.GetByID(c, tx, req.TopicType)
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
		_, exist, errInner := p.TopicRepository.GetByCondition(c, tx, map[string]string{
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

		errInner = p.TopicSetRepository.Insert(c, tx, set)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
	}

	err = p.TopicRepository.Insert(c, tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkSaveCategories(c, tx, *ctx.AccountID, id, req.Categories)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkCreateMembers(c, tx, *ctx.AccountID, id, req)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkSaveRelations(c, tx, *ctx.AccountID, id, req.RelatedTopics)
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

func (p *TopicUsecase) Delete(c context.Context, ctx *biz.BizContext, topicID int64) (err error) {
	return p.TopicRepository.Delete(c, p.Node, topicID)
}

func (p *TopicUsecase) Update(c context.Context, ctx *biz.BizContext, topicID int64, req *models.UpdateTopicReq) (err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.TopicRepository.GetByID(c, tx, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("话题不存在")
		return
	}

	if req.TopicType != nil {
		_, exist, errInner := p.TopicTypeRepository.GetByID(c, tx, *req.TopicType)
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
		item.Cover = req.Cover
	}

	if req.Bg != nil && *req.Bg != "" {
		item.Bg = req.Bg
	}

	if req.Name != nil && *req.Name != "" {
		item.Name = *req.Name
	}

	if req.Introduction != nil && *req.Introduction != "" {
		item.Introduction = *req.Introduction
	}

	if req.VersionName != nil && *req.VersionName != "" {
		item.VersionName = *req.VersionName

		dbItem, exist, errInner := p.TopicRepository.GetByCondition(c, tx, map[string]string{
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

	err = p.TopicRepository.Update(c, tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkSaveCategories(c, tx, *ctx.AccountID, item.ID, req.Categories)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.bulkSaveRelations(c, tx, *ctx.AccountID, item.ID, req.RelatedTopics)
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

func (p *TopicUsecase) bulkSaveRelations(c context.Context, node sqalx.Node, accountID, topicID int64, relations []*models.RelatedTopicReq) (err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	for _, v := range relations {
		relation, exist, errInner := p.TopicRelationRepository.GetByCondition(c, tx, map[string]string{
			"from_topic_id": strconv.FormatInt(topicID, 10),
			"top_topic_id":  strconv.FormatInt(v.TopicID, 10),
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if exist {
			relation.Relation = v.Type
			errInner = p.TopicRelationRepository.Update(c, tx, relation)
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
			errInner = p.TopicRelationRepository.Insert(c, tx, item)
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

func (p *TopicUsecase) bulkCreateMembers(c context.Context, node sqalx.Node, accountID, topicID int64, req *models.CreateTopicReq) (err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	for _, v := range req.Members {
		if v.AccountID == accountID {
			continue
		}

		if v.Role == models.MemberRoleOwner {
			err = berr.Errorf("主理人只能有一个")
			return
		}

		member, exist, errInner := p.TopicMemberRepository.GetByCondition(c, tx, map[string]string{
			"topic_id":   strconv.FormatInt(topicID, 10),
			"account_id": strconv.FormatInt(v.AccountID, 10),
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if exist {
			member.Role = v.Role
			errInner = p.TopicMemberRepository.Update(c, tx, member)
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
			errInner = p.TopicMemberRepository.Insert(c, tx, item)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}

			errInner = p.followTopic(c, tx, topicID, v.AccountID)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}

			break
		}
	}

	id, errInner := gid.NextID()
	if errInner != nil {
		err = tracerr.Wrap(errInner)
		return
	}
	item := &repo.TopicMember{
		ID:        id,
		AccountID: accountID,
		Role:      models.MemberRoleOwner,
		TopicID:   topicID,
	}
	errInner = p.TopicMemberRepository.Insert(c, tx, item)
	if errInner != nil {
		err = tracerr.Wrap(errInner)
		return
	}

	errInner = p.followTopic(c, tx, topicID, accountID)
	if errInner != nil {
		err = tracerr.Wrap(errInner)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicUsecase) followTopic(c context.Context, node sqalx.Node, topicID, accountID int64) (err error) {
	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	follower := &repo.TopicFollower{
		ID:          id,
		TopicID:     topicID,
		FollowersID: accountID,
	}

	err = p.TopicFollowerRepository.Insert(c, node, follower)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

func (p *TopicUsecase) bulkSaveCategories(c context.Context, node sqalx.Node, accountID int64, topicID int64, categories []*models.TopicCategoryParentItem) (err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	dicMap := make(map[int64]bool)
	items, err := p.TopicCategoryRepository.GetAllByCondition(c, tx, map[string]string{
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
				item, _, errInner := p.TopicCategoryRepository.GetByID(c, tx, *parent.ID)
				if errInner != nil {
					err = tracerr.Wrap(errInner)
					return
				}

				item.Name = parent.Name
				item.Seq = parent.Seq

				errInner = p.TopicCategoryRepository.Update(c, tx, item)
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

			errInner = p.TopicCategoryRepository.Insert(c, tx, item)
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
					item, _, errInner := p.TopicCategoryRepository.GetByID(c, tx, *child.ID)
					if errInner != nil {
						err = tracerr.Wrap(errInner)
						return
					}

					item.Name = child.Name
					item.Seq = child.Seq
					item.ParentID = parentID

					errInner = p.TopicCategoryRepository.Update(c, tx, item)
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

				errInner = p.TopicCategoryRepository.Insert(c, tx, item)
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

func (p *TopicUsecase) getCategoriesHierarchyOfAll(c context.Context, node sqalx.Node, topicID int64) (items []*models.TopicCategoryParentItem, err error) {
	parents, err := p.TopicCategoryRepository.GetAllByCondition(c, node, map[string]string{
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

		children, errInner := p.TopicCategoryRepository.GetAllByCondition(c, node, map[string]string{
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

func (p *TopicUsecase) GetTopicVersions(c context.Context, ctx *biz.BizContext, topicSetID int64) (items []*models.TopicVersion, err error) {
	return p.TopicRepository.GetTopicVersions(c, p.Node, topicSetID)
}

func (p *TopicUsecase) BulkSaveMembers(c context.Context, ctx *biz.BizContext, topicID int64, req *models.BatchSavedTopicMemberReq) (err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	_, exist, err := p.TopicRepository.GetByID(c, p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("话题不存在")
		return
	}

	for _, v := range req.Members {
		if v.Role == models.MemberRoleOwner {
			continue
		}
		member, exist, errInner := p.TopicMemberRepository.GetByCondition(c, tx, map[string]string{
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
			errInner = p.TopicMemberRepository.Insert(c, tx, item)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}

			errInner = p.followTopic(c, tx, topicID, v.AccountID)
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
			errInner = p.TopicMemberRepository.Update(c, tx, member)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
			break
		case "D":
			if !exist {
				continue
			}
			errInner = p.TopicMemberRepository.Delete(c, tx, member.ID)
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

func (p *TopicUsecase) GetAllTopicTypes(c context.Context, ctx *biz.BizContext) (items []*models.TopicType, err error) {
	data, err := p.TopicTypeRepository.GetAll(c, p.Node)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	items = make([]*models.TopicType, 0)

	copier.Copy(&items, &data)

	return
}

func (p *TopicUsecase) GetTopicMeta(c context.Context, node sqalx.Node, accountID int64, topicID int64) (meta models.TopicMeta, err error) {
	topic, exist, err := p.TopicRepository.GetByID(c, p.Node, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("话题不存在")
		return
	}

	account, exist, err := p.AccountRepository.GetByID(c, node, accountID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("未找到该账号")
		return
	}

	switch topic.JoinPermission {
	case models.JoinPermissionMember:
		meta.CanJoin = true
		break
	case models.JoinPermissionIDCert:
		if bool(account.IDCert) {
			meta.CanJoin = true
		}
		break
	case models.JoinPermissionWorkCert:
		if bool(account.WorkCert) {
			meta.CanJoin = true
		}
		break
	case models.JoinPermissionMemberApprove:
		break
	case models.JoinPermissionIDCertApprove:
		break
	case models.JoinPermissionWorkCertApprove:
		break
	case models.JoinPermissionAdminAdd:
		break
	case models.JoinPermissionPurchase:
		break
	case models.JoinPermissionVIP:
		if bool(account.WorkCert) {
			meta.CanJoin = true
		}
		break
	}

	switch topic.ViewPermission {
	case models.ViewPermissionPublic:
		meta.CanView = true
		break
	case models.ViewPermissionJoin:
		break
	}

	switch topic.EditPermission {
	case models.EditPermissionIDCert:
		break
	case models.EditPermissionWorkCert:
		break
	case models.EditPermissionIDCertJoined:
		break
	case models.EditPermissionWorkCertJoined:
		break
	case models.EditPermissionApprovedIDCertJoined:
		break
	case models.EditPermissionApprovedWorkCertJoined:
		break
	case models.EditPermissionAdmin:
		meta.CanEdit = true
		break
	}

	return
}

func (p *TopicUsecase) CreateNewVersion(c context.Context, ctx *biz.BizContext, arg *models.ArgNewVersion) (id int64, err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	t, exist, err := p.TopicRepository.GetByID(c, tx, arg.FromTopicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("话题不存在")
		return
	}

	_, exist, err = p.TopicRepository.GetByCondition(c, tx, map[string]string{
		"topic_set_id": strconv.FormatInt(t.TopicSetID, 10),
		"version_name": arg.VersionName,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if exist {
		err = berr.Errorf("话题版本已经存在")
		return
	}

	id, err = gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	t.ID = id
	t.VersionName = arg.VersionName
	t.CreatedBy = *ctx.AccountID
	t.CreatedAt = time.Now().Unix()
	t.UpdatedAt = time.Now().Unix()

	err = p.TopicRepository.Insert(c, tx, t)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	members, err := p.TopicMemberRepository.GetAllTopicMembers(c, tx, arg.FromTopicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	for _, v := range members {
		sid, errInner := gid.NextID()
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		v.ID = sid
		v.TopicID = t.ID
		v.CreatedAt = time.Now().Unix()
		v.UpdatedAt = time.Now().Unix()

		errInner = p.TopicMemberRepository.Insert(c, tx, v)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
	}

	relations, err := p.TopicRelationRepository.GetAllTopicRelations(c, tx, arg.FromTopicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	for _, v := range relations {
		sid, errInner := gid.NextID()
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		v.ID = sid
		v.FromTopicID = t.ID
		v.CreatedAt = time.Now().Unix()
		v.UpdatedAt = time.Now().Unix()

		errInner = p.TopicRelationRepository.Insert(c, tx, v)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
	}

	categories, err := p.TopicCategoryRepository.GetAllByCondition(c, tx, map[string]string{
		"topic_id": strconv.FormatInt(arg.FromTopicID, 10),
	})

	for _, v := range categories {
		sid, errInner := gid.NextID()
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		v.ID = sid
		v.TopicID = t.ID
		v.CreatedAt = time.Now().Unix()
		v.UpdatedAt = time.Now().Unix()

		errInner = p.TopicCategoryRepository.Insert(c, tx, v)
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
