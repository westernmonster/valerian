package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error) {
	item = new(model.Article)
	sqlSelect := "SELECT a.id,a.title,a.content,a.content_text,a.disable_revise,a.disable_comment,a.created_by,a.deleted,a.created_at,a.updated_at FROM articles a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleByID err(%+v), id(%+v)", err, id))
	}

	return
}
