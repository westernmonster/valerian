package service

import (
	"context"
	"fmt"
	"github.com/nats-io/stan.go"
	"valerian/app/service/feed/def"
	"valerian/app/service/search/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) onSearchStatAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)
	info := new(def.MsgSearchStatAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onSearchStatAdded Unmarshal failed %#v", err))
		return
	}

	item := &model.ESSearchStat{
		ID:         gid.NewID(),
		Keywords:   info.Keywords,
		CreatedBy:  info.CreatedBy,
		Hits:       info.Hits,
		Enterpoint: info.Enterpoint,
		CreatedAt:  info.CreatedAt,
	}

	if err = p.d.PutSearchStat2ES(c, item); err != nil {
		return
	}

	m.Ack()
}
