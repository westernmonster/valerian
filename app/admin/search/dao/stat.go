package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/admin/search/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetArticleStatByID(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleStat, err error) {
	item = new(model.ArticleStat)
	sqlSelect := "SELECT a.* FROM article_stats a WHERE a.article_id=?"

	if err = node.GetContext(c, item, sqlSelect, articleID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleStatByID err(%+v), article_id(%+v)", err, articleID))
	}

	return
}

func (p *Dao) GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountStat, err error) {
	item = new(model.AccountStat)
	sqlSelect := "SELECT a.* FROM account_stats a WHERE a.account_id=? "

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

func (p *Dao) GetDiscussionStatByID(c context.Context, node sqalx.Node, discussionID int64) (item *model.DiscussionStat, err error) {
	item = new(model.DiscussionStat)
	sqlSelect := "SELECT a.* FROM discussion_stats a WHERE a.discussion_id=?"

	if err = node.GetContext(c, item, sqlSelect, discussionID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionStatByID err(%+v), discussion_id(%+v)", err, discussionID))
	}

	return
}

func (p *Dao) GetTopicStatByID(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicStat, err error) {
	item = new(model.TopicStat)
	sqlSelect := "SELECT a.* FROM topic_stats a WHERE a.topic_id=? "

	if err = node.GetContext(c, item, sqlSelect, topicID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicStatByID err(%+v), topic_id(%+v)", err, topicID))
	}

	return
}
