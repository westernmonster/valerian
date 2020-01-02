package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetArticleHistoryByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleHistory, err error) {
	item = new(model.ArticleHistory)
	sqlSelect := "SELECT a.id,a.article_id,a.content,a.content_text,a.seq,a.diff,a.updated_by,a.change_desc,a.deleted,a.created_at,a.updated_at FROM article_histories a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoryByID err(%+v), id(%+v)", err, id))
	}

	return
}
