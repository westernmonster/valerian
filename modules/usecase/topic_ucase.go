package usecase

import (
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"

	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
)

type TopicUsecase struct {
	sqalx.Node
	*sqlx.DB

	TopicRepository interface {
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.Topic, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.Topic, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.Topic, err error)
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
	}
}

func (p *TopicUsecase) Create(ctx *biz.BizContext, req *models.CreateTopicReq) (id int64, err error) {
	return
}

func (p *TopicUsecase) Delete(ctx *biz.BizContext, topicID int64) (err error) {
	return
}

func (p *TopicUsecase) Update(ctx *biz.BizContext, topicID int64, req *models.UpdateTopicReq) (err error) {
	return
}

func (p *TopicUsecase) BulkSaveMembers(ctx *biz.BizContext, topicID int64, req *models.BulkSaveTopicMembersReq) (err error) {
	return
}
