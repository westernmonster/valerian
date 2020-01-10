package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/log"
)

// delRevise 删除补充
func (p *Service) delRevise(c context.Context, aid, id int64) (err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var item *model.Revise
	if item, err = p.d.GetReviseByID(c, tx, id); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}

	if err = p.checkEditPermission(c, tx, aid, item.ArticleID); err != nil {
		return
	}

	if err = p.d.DelRevise(c, tx, id); err != nil {
		return
	}

	if err = p.d.IncrArticleStat(c, tx, &model.ArticleStat{ArticleID: item.ArticleID, ReviseCount: -1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelReviseCache(context.TODO(), id)
		p.onReviseDeleted(context.Background(), id, aid, time.Now().Unix())
	})
	return
}
