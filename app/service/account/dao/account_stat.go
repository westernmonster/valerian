package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/app/service/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by Account ID
func (p *Dao) GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountResStat, err error) {
	item = new(model.AccountResStat)
	sqlSelect := "SELECT a.* FROM account_res_stats a WHERE a.account_id=? "

	if err = node.GetContext(c, item, sqlSelect, aid); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountStatByID err(%+v), account_id(%+v)", err, aid))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccountStat(c context.Context, node sqalx.Node, item *model.AccountResStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO account_res_stats( account_id,topic_count,article_count,discussion_count,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.AccountID, item.TopicCount, item.ArticleCount, item.DiscussionCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountStat(c context.Context, node sqalx.Node, item *model.AccountResStat) (err error) {
	sqlUpdate := "UPDATE account_res_stats SET account_id=?,topic_count=?,article_count=?,discussion_count=?,updated_at=? WHERE account_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.TopicCount, item.ArticleCount, item.DiscussionCount, item.UpdatedAt, item.AccountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) IncrAccountStat(c context.Context, node sqalx.Node, item *model.AccountResStat) (err error) {
	sqlUpdate := "UPDATE account_res_stats SET topic_count=topic_count+?,article_count=article_count+?,discussion_count=discussion_count+?,updated_at=? WHERE account_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicCount, item.ArticleCount, item.DiscussionCount, time.Now().Unix(), item.AccountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrAccountStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
