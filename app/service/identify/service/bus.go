package service

import (
	"context"
	"fmt"

	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) onAccountAdded(c context.Context, aid, actionTime int64) {
	msg := &def.MsgAccountAdded{AccountID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onAccountAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusAccountAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onAccountAdded.Publish(), err(%+v)", err))
		return
	}

	return
}
