package usecase

import (
	"strconv"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"

	"github.com/ztrue/tracerr"

	"valerian/infrastructure/berr"
	"valerian/infrastructure/biz"
	"valerian/infrastructure/gid"
	"valerian/models"
	"valerian/modules/repo"
)

type TopicCategoryUsecase struct {
	sqalx.Node
	*sqlx.DB

	TopicRepository interface {
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.Topic, exist bool, err error)
	}

	TopicCategoryRepository interface {
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.TopicCategory, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.TopicCategory, err error)
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
}

func (p *TopicCategoryUsecase) GetAll(ctx *biz.BizContext, topicID int64) (items []*models.TopicCategory, err error) {
	allItems, err := p.TopicCategoryRepository.GetAllByCondition(p.Node, map[string]string{
		"topic_id": strconv.FormatInt(topicID, 10),
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	// TODO: 是否有数据所有权的权限认证？

	items = make([]*models.TopicCategory, 0)
	for _, v := range allItems {
		item := &models.TopicCategory{
			ID:        v.ID,
			TopicID:   v.TopicID,
			Name:      v.Name,
			ParentID:  v.ParentID,
			CreatedBy: v.CreatedBy,
			Seq:       v.Seq,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		items = append(items, item)

	}

	return

}

func (p *TopicCategoryUsecase) GetHierarchyOfAll(ctx *biz.BizContext, topicID int64) (resp *models.TopicCategoriesResp, err error) {
	resp = new(models.TopicCategoriesResp)
	parents, err := p.TopicCategoryRepository.GetAllByCondition(p.Node, map[string]string{
		"topic_id":  strconv.FormatInt(topicID, 10),
		"parent_id": "0",
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	// TODO: 是否有数据所有权的权限认证？
	resp.TopicID = topicID
	resp.Items = make([]*models.TopicCategoryParentItem, 0)

	for _, v := range parents {
		parent := &models.TopicCategoryParentItem{
			ID:       &v.ID,
			Name:     v.Name,
			Seq:      v.Seq,
			Children: make([]*models.TopicCategoryChildItem, 0),
		}

		children, errInner := p.TopicCategoryRepository.GetAllByCondition(p.Node, map[string]string{
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

		resp.Items = append(resp.Items, parent)

	}

	return

}

func (p *TopicCategoryUsecase) Create(ctx *biz.BizContext, req *models.CreateTopicCategoryReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	// detect topic exist
	_, exist, err := p.TopicRepository.GetByID(tx, req.TopicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("未找到所属话题")
		return
	}

	// detect parentID correct
	if req.ParentID != 0 {
		parent, exist, errInner := p.TopicCategoryRepository.GetByID(tx, req.ParentID)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if !exist {
			err = berr.Errorf("未找到所属话题")
			return
		}

		if parent.ParentID != 0 {
			err = berr.Errorf("不能创建三级分类")
			return
		}
	}

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	item := &repo.TopicCategory{
		ID:        id,
		TopicID:   req.TopicID,
		Name:      req.Name,
		ParentID:  req.ParentID,
		Seq:       req.Seq,
		CreatedBy: *ctx.AccountID,
	}

	err = p.TopicCategoryRepository.Insert(tx, item)
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

func (p *TopicCategoryUsecase) Update(ctx *biz.BizContext, id int64, req *models.UpdateTopicCategoryReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	// detect parentID correct
	if req.ParentID != 0 {
		parent, exist, errInner := p.TopicCategoryRepository.GetByID(tx, req.ParentID)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if !exist {
			err = berr.Errorf("未找到所属话题")
			return
		}

		if parent.ParentID != 0 {
			err = berr.Errorf("不能创建三级分类")
			return
		}
	}

	// detect topic exist
	item, exist, err := p.TopicCategoryRepository.GetByID(tx, id)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("未找到该分类")
		return
	}

	item.Name = req.Name
	item.ParentID = req.ParentID
	item.Seq = req.Seq

	err = p.TopicCategoryRepository.Update(tx, item)
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

func (p *TopicCategoryUsecase) Delete(ctx *biz.BizContext, id int64) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	// detect topic exist
	item, exist, err := p.TopicCategoryRepository.GetByID(tx, id)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		return
	}

	children, err := p.TopicCategoryRepository.GetAllByCondition(tx, map[string]string{
		"parent_id": strconv.FormatInt(item.ID, 10),
	})

	for _, v := range children {
		errInner := p.TopicCategoryRepository.Delete(tx, v.ID)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
	}

	err = p.TopicCategoryRepository.Delete(tx, id)
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

func (p *TopicCategoryUsecase) BulkSave(ctx *biz.BizContext, req *models.SaveTopicCategoriesReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	// detect topic exist
	_, exist, err := p.TopicRepository.GetByID(tx, req.TopicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("未找到所属话题")
		return
	}

	dicMap := make(map[int64]bool)
	items, err := p.TopicCategoryRepository.GetAllByCondition(tx, map[string]string{
		"topic_id": strconv.FormatInt(req.TopicID, 10),
	})
	for _, v := range items {
		dicMap[v.ID] = true
	}

	for _, parent := range req.Items {
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
				TopicID:   req.TopicID,
				Name:      parent.Name,
				ParentID:  0,
				Seq:       parent.Seq,
				CreatedBy: *ctx.AccountID,
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
					TopicID:   req.TopicID,
					Name:      child.Name,
					ParentID:  parentID,
					Seq:       child.Seq,
					CreatedBy: *ctx.AccountID,
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
