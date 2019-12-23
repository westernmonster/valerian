package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked Unmarshal failed %#v", err))
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

	var article *model.Article
	if article, err = p.getArticle(c, tx, info.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  article.CreatedBy,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeArticle,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   article.ID,
		TargetType: model.TargetTypeArticle,
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
		if setting.NotifyLike {
			if _, err := p.pushSingleUser(context.Background(),
				msg.AccountID,
				msg.ID,
				def.PushMsgTitleArticleLiked,
				def.PushMsgTitleArticleLiked,
				fmt.Sprintf(def.LinkArticle, article.ID),
			); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onArticleLiked Push message failed %#v", err))
			}
		}
	})

}

func (p *Service) onReviseLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked Unmarshal failed %#v", err))
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

	var revise *model.Revise
	if revise, err = p.getRevise(c, tx, info.ReviseID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetRevise", "GetRevise(), id(%d),error(%+v)", info.ReviseID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  revise.CreatedBy,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeRevise,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   revise.ID,
		TargetType: model.TargetTypeRevise,
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
		if setting.NotifyLike {
			if _, err := p.pushSingleUser(context.Background(),
				msg.AccountID,
				msg.ID,
				def.PushMsgTitleReviseLiked,
				def.PushMsgTitleReviseLiked,
				fmt.Sprintf(def.LinkRevise, revise.ID),
			); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onReviseLiked Push message failed %#v", err))
			}
		}
	})

}

func (p *Service) onDiscussionLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked Unmarshal failed %#v", err))
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

	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, tx, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked GetDiscussion failed %#v", err))
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  discuss.CreatedBy,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeDiscussion,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   discuss.ID,
		TargetType: model.TargetTypeDiscussion,
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
		if setting.NotifyLike {
			if _, err := p.pushSingleUser(context.Background(),
				msg.AccountID,
				msg.ID,
				def.PushMsgTitleDiscussionLiked,
				def.PushMsgTitleDiscussionLiked,
				fmt.Sprintf(def.LinkDiscussion, discuss.ID),
			); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onDiscussionLiked Push message failed %#v", err))
			}
		}
	})
}

func (p *Service) onCommentLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgCommentLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentLiked Unmarshal failed %#v", err))
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

	var comment *model.Comment
	if comment, err = p.getComment(c, tx, info.CommentID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentLiked GetComment failed %#v", err))
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  comment.CreatedBy,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeComment,
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
		if setting.NotifyLike {
			if _, err := p.pushSingleUser(context.Background(),
				msg.AccountID,
				msg.ID,
				def.PushMsgTitleCommentLiked,
				def.PushMsgTitleCommentLiked,
				fmt.Sprintf(def.LinkComment, comment.OwnerType, comment.OwnerID, comment.ID),
			); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onCommentLiked Push message failed %#v", err))
			}
		}
	})

}
