package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by Account ID
func (p *Dao) GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountStat, err error) {
	item = new(model.AccountStat)
	sqlSelect := "SELECT a.account_id,a.following,a.fans,a.article_count,a.discussion_count,a.topic_count,a.black,a.created_at,a.updated_at FROM account_stats a WHERE a.account_id=? "

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

func (p *Dao) IncrAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error) {
	sqlUpdate := "UPDATE account_stats SET following=following + ?,fans=fans+?,article_count=article_count+?,discussion_count=discussion_count+?,topic_count=topic_count+?,black=black+?,updated_at=? WHERE account_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Following, item.Fans, item.ArticleCount, item.DiscussionCount, item.TopicCount, item.Black, item.UpdatedAt, item.AccountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrAccountStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) AddAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO account_stats( account_id,following,fans,article_count,discussion_count,topic_count,black,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.AccountID, item.Following, item.Fans, item.ArticleCount, item.DiscussionCount, item.TopicCount, item.Black, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error) {
	sqlUpdate := "UPDATE account_stats SET following=?,fans=?,article_count=?,discussion_count=?,topic_count=?,black=?,updated_at=? WHERE account_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Following, item.Fans, item.ArticleCount, item.DiscussionCount, item.TopicCount, item.Black, item.UpdatedAt, item.AccountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
