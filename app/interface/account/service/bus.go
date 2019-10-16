package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) onAccountUpdated(c context.Context, aid, actionTime int64) {
	msg := &def.MsgAccountUpdated{AccountID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onAccountUpdated.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusAccountUpdated, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onAccountUpdated.Publish(), err(%+v)", err))
		return
	}

	return
}
