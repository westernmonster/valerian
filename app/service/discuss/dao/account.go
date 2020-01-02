package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error) {
	item = &model.Account{}
	sqlSelect := "SELECT a.id,a.mobile,a.user_name,a.email,a.password,a.role,a.salt,a.gender,a.birth_year,a.birth_month,a.birth_day,a.location,a.introduction,a.avatar,a.source,a.ip,a.id_cert,a.work_cert,a.is_org,a.is_vip,a.deleted,a.created_at,a.updated_at,a.prefix,a.is_lock,a.deactive FROM accounts a WHERE a.id=? AND a.deleted=0"
	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByID error(%+v), id(%d)", err, id))
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
