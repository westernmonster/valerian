package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/kamilsk/retry/v4"
	"github.com/kamilsk/retry/v4/strategy"
	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleCommented Unmarshal failed %#v", err))
		return
	}

	var comment *comment.CommentInfo
	action := func(c context.Context, _ uint) error {
		ct, e := p.d.GetComment(c, info.CommentID, true)
		if e != nil {
			return e
		}

		comment = ct
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	var article *article.ArticleInfo
	action = func(c context.Context, _ uint) error {
		ct, e := p.d.GetArticle(c, comment.OwnerID, true)
		if e != nil {
			return e
		}

		article = ct
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  article.Creator.ID,
		ActionType: model.MsgArticleCommented,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextArticleCommented,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   comment.ID,
		TargetType: model.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentAdded AddMessage failed %#v", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(context.Background(), p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}

		if setting.NotifyComment {
			if _, err := p.pushSingleUser(context.Background(),
				msg.AccountID,
				msg.ID,
				def.PushMsgTitleArticleCommented,
				def.PushMsgTitleArticleCommented,
				fmt.Sprintf(def.LinkComment, model.TargetTypeArticle, comment.OwnerID, comment.ID),
			); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onArticleCommented Push message failed %#v", err))
			}
		}
	})

}

func (p *Service) onReviseCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseCommented Unmarshal failed %#v", err))
		return
	}

	var comment *comment.CommentInfo
	if comment, err = p.d.GetComment(c, info.CommentID, true); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseCommented GetComment failed %#v", err))
		return
	}

	var revise *article.ReviseInfo
	if revise, err = p.d.GetRevise(c, comment.OwnerID, true); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentAdded GetRevise failed %#v", err))
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  revise.Creator.ID,
		ActionType: model.MsgReviseCommented,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextReviseCommented,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   comment.ID,
		TargetType: model.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentAdded AddMessage failed %#v", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(context.Background(), p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}
		if setting.NotifyComment {
			if _, err := p.pushSingleUser(context.Background(),
				msg.AccountID,
				msg.ID,
				def.PushMsgTitleReviseCommented,
				def.PushMsgTitleReviseCommented,
				fmt.Sprintf(def.LinkComment, model.TargetTypeRevise, comment.OwnerID, comment.ID),
			); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onReviseCommented Push message failed %#v", err))
			}
		}
	})
}

func (p *Service) onDiscussionCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionCommented Unmarshal failed %#v", err))
		return
	}

	var comment *comment.CommentInfo
	if comment, err = p.d.GetComment(c, info.CommentID, true); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionCommented GetComment failed %#v", err))
		return
	}

	var discussion *discuss.DiscussionInfo
	if discussion, err = p.d.GetDiscussion(c, comment.OwnerID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentAdded GetDiscussion failed %#v", err))
		return
	}

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

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  discussion.Creator.ID,
		ActionType: model.MsgDiscussionCommented,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextDiscussionCommented,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   comment.ID,
		TargetType: model.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentAdded AddMessage failed %#v", err))
		return
	}

	if err = p.d.IncrMessageStat(c, tx, &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(context.Background(), p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}
		if setting.NotifyComment {
			if _, err := p.pushSingleUser(context.Background(),
				msg.AccountID,
				msg.ID,
				def.PushMsgTitleDiscussionCommented,
				def.PushMsgTitleDiscussionCommented,
				fmt.Sprintf(def.LinkComment, model.TargetTypeRevise, comment.OwnerID, comment.ID),
			); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onDiscussionCommented Push message failed %#v", err))
			}
		}
	})
}

func (p *Service) onCommentReplied(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgCommentReplied)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentReplied Unmarshal failed %#v", err))
		return
	}

	var comment *comment.CommentInfo
	action := func(c context.Context, _ uint) error {
		ct, e := p.d.GetComment(c, info.CommentID, true)
		if e != nil {
			return e
		}

		comment = ct
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentReplied GetComment failed %#v", err))
		m.Ack()
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  comment.Creator.ID,
		ActionType: model.MsgCommentReplied,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextCommentReplied,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   comment.ID,
		TargetType: model.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentReplied AddMessage failed %#v", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(context.Background(), p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}

		if setting.NotifyComment {
			if _, err := p.pushSingleUser(context.Background(),
				msg.AccountID,
				msg.ID,
				def.PushMsgTitleCommentReplied,
				def.PushMsgTitleCommentReplied,
				fmt.Sprintf(def.LinkSubComment, comment.OwnerType, comment.OwnerID, comment.ResourceID, comment.ID),
			); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onCommentReplied Push message failed %#v", err))
			}
		}
	})

}
