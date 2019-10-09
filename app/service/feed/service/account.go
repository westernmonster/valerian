package service

import (
	"context"
	"fmt"
	"time"

	account "valerian/app/service/account/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
	relation "valerian/app/service/relation/api"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onMemberFollowed(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgMemberFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated Unmarshal failed %#v", err))
		return
	}

	var member *account.BaseInfoReply
	if member, err = p.d.GetAccountBaseInfo(context.Background(), info.TargetAccountID); err != nil {
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.AccountID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated GetFansIDs failed %#v", err))
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

	for _, v := range fansResp.IDs {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeFollowMember,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextFollowMember,
			ActorID:    info.AccountID,
			ActorType:  def.ActorTypeUser,
			TargetID:   member.ID,
			TargetType: def.TargetTypeMember,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}
