package dao

import (
	"context"
	"fmt"

	message "valerian/app/service/message/api"
	"valerian/library/log"
)

func (p *Dao) GetMessageStat(c context.Context, aid int64) (info *message.MessageStat, err error) {
	if info, err = p.messageRPC.GetMessageStat(c, &message.AidReq{AccountID: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetMessageStat err(%+v) aid(%d)", err, aid))
	}
	return
}
