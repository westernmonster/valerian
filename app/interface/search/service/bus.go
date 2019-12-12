package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) triggerSearchStatAdded(c context.Context, keywords, enterpoint string, createBy, hits int64) {
	msg := &def.MsgSearchStatAdded{
		Keywords:   keywords,
		CreatedBy:  createBy,
		Hits:       hits,
		Enterpoint: enterpoint,
		CreatedAt:  time.Now().Unix(),
	}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("triggerSearchStatAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusSearchStatAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("triggerSearchStatAdded.Publish(), err(%+v)", err))
		return
	}

	return
}
